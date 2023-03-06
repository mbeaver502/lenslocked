package models

import (
	"database/sql"
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

	// TODO: Hash the session token

	session := Session{
		UserID: userID,
		Token:  token,
		//TokenHash: ???,
	}

	// TODO: Store session in DB

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}

func (ss *SessionService) sessionToken() (string, error) {
	n := ss.BytesPerSessionToken
	if n < MinBytesPerSessionToken {
		n = MinBytesPerSessionToken
	}

	return rand.String(n)
}
