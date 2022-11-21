package models

import (
	"SideProject/store/data"
	"fmt"
	"time"
)

type Session struct {
	Id int 
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

// variable to store the last time update the session
var sessionInterval int

func init() {
	sessionInterval = 10 
}

//CreateSession :create the session for login user
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, user_id, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		user.Uuid, user.Id, user.Email, user.CreatedAt,
		).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// DeleteByUuid
func (session Session) DeleteByUuid() (err error) {
	statement := "DELETE FROM sessions WHERE uuid=$1"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// UpdateSession :
func UpdateSession() (err error) {
	statement := fmt.Sprintf(
		"DELETE FROM sessions WHERE (current_timestamp - created_at) > INTERVAL '%v MINUTES'", 
		sessionInterval,
	)
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	_, err = stmt.Exec()
	return
}

// SessionDeleteAll : delete all sessiom
func SessionDeleteAll()(err error) {
	statement := "DELETE FROM sessions"
	_, err = data.Db.Exec(statement)
	return
} 