package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	email     string
	id        int
	password  string
	firstName string
}

// Email - Get member's email
func (member *Member) Email() string {
	return member.email
}

// ID - Get member's Id
func (member *Member) ID() int {
	return member.id
}

// Password - Get member's Password
func (member *Member) Password() string {
	return member.password
}

// FirstName - Get member's FirstName
func (member *Member) FirstName() string {
	return member.firstName
}

// SetEmail - Sets email for a member
func (member *Member) SetEmail(value string) {
	member.email = value
}

// SetID - Sets id for a member
func (member *Member) SetID(value int) {
	member.id = value
}

// SetPassword - Sets password for a member
func (member *Member) SetPassword(value string) {
	member.password = value
}

// SetFirstName - Sets the first name of a member
func (member *Member) SetFirstName(value string) {
	member.firstName = value
}

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		row := db.QueryRow(`SELECT id, email, first_name
			FROM Member
			WHERE email = $1 AND password = $2`, email, hex.EncodeToString(password))
		result := Member{}
		err := row.Scan(&result.id, &result.email, &result.firstName)

		if err != nil {
			return result, errors.New("Unable to find member with email: " + email)
		}
		return result, nil
	} else {
		return Member{}, errors.New("Unable to get database connection")
	}
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])

	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO Session
			(member_id, session_id)
			VALUES ($1, $2)
			RETURNING id`, member.ID(), result.sessionId).Scan(&result.id)

		if err == nil {
			return result, nil
		} else {
			return Session{}, erros.New("Unable to save session to database")
		}
	} else {
		return result, Errors.New("Unable to get database connection")
	}
}
