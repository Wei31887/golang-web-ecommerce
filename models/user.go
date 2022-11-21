package models

import (
	"SideProject/store/data"
	"time"
)

type User struct {
	Id int
	Uuid string
	Name string
	Password string
	Email string
	Role int
	CreatedAt time.Time 
} 


// Create : create the new user
func (user *User) Create() (err error) {
	statement := "INSERT INTO users (uuid, name, password, email, role, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, uuid, created_at"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// create the admin user if the email and the user name are matched
	if user.Name == "admin" && user.Email == "admin@test.com" {
		user.Role = 1
		err = stmt.QueryRow(
			createUuid(), user.Name, user.Password, user.Email, user.Role, time.Now(),
			).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
		return
	}
	
	// create the standard user
	user.Role = 0
	err = stmt.QueryRow(
		createUuid(), user.Name, user.Password, user.Email, user.Role, time.Now(),
		).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// // GetUserRoleByEmail : get the user role  by email
// func GetUserRoleByEmail(email string) (user User, err error) {
// 	user = User{}
// 	err = data.Db.QueryRow(
// 		"SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
// 		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt) 
// 	return
// }

// GetUserByEmail : get the user info by email
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = data.Db.QueryRow(
		"SELECT id, uuid, name, email, password, role, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt) 
	return
}

func GetUserByUuid(uuid string) (user User, err error) {
	user = User{}
	err = data.Db.QueryRow(
		"SELECT id, uuid, name, email, password, role, created_at FROM users WHERE uuid = $1", 
		uuid,
		).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt) 
	return
}

//Delete : delete user from database
func (user *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE id=$1"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

//GetUsers : get all users in database
func GetUsers() (users []User, err error) {
	statement := "SELECT id, uuid, name, email, password, role, created_at FROM users"
	rows, err := data.Db.Query(statement)
	if err != nil {
		return
	}	
	
	for rows.Next() {
		user := User{}
		if err = rows.Scan(
			&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt,
			); err != nil {
				return 
		}
		users = append(users, user)
	} 
	return
}

func UserDeleteAll() (err error) {
	statement := "DELETE FROM users"
	_, err = data.Db.Exec(statement)
	return
}
