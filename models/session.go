package models

import (
	"database/sql"
	"fmt"

	"github.com/mbeaver502/lenslocked/rand"
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
	DB *sql.DB
}

func (ss *SessionService) Create(userID uint) (*Session, error) {
	token, err := rand.SessionToken()
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
