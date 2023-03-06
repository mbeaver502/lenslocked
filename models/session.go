package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/mbeaver502/lenslocked/rand"
)

const (
	MinBytesPerSessionToken = 32
)

type Session struct {
	ID     uint
	UserID uint
	// Token is only set when creating a new session.
	// When looking up a session, this will be left empty,
	// as we only store the hash of a session token in the DB
	// and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB                   *sql.DB
	BytesPerSessionToken int
}

func (ss *SessionService) Create(userID uint) (*Session, error) {
	token, err := ss.sessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}

	// Do an upsert to add the new session token to the DB.
	row := ss.DB.QueryRow(`update sessions set token_hash = $2 where user_id = $1 returning id`,
		session.UserID,
		session.TokenHash,
	)
	err = row.Scan(&session.ID)

	// Update returned no rows affected, so we know we can insert a new record.
	if err == sql.ErrNoRows {
		row = ss.DB.QueryRow(`insert into sessions (user_id, token_hash) values ($1, $2) returning id`,
			session.UserID,
			session.TokenHash,
		)

		err = row.Scan(&session.ID)
	}

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	hash := ss.hash(token)

	var user User

	row := ss.DB.QueryRow(`select user_id from sessions where token_hash = $1`, hash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	row = ss.DB.QueryRow(`select email, password_hash from users where id = $1`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	hash := ss.hash(token)

	_, err := ss.DB.Exec(`delete from sessions where token_hash = $1`, hash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (ss *SessionService) sessionToken() (string, error) {
	n := ss.BytesPerSessionToken
	if n < MinBytesPerSessionToken {
		n = MinBytesPerSessionToken
	}

	return rand.String(n)
}

func (ss *SessionService) hash(token string) string {
	hash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(hash[:])
}
