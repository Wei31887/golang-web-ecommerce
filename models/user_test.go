package models

import (
	"fmt"
	"testing"
)

// Set up for testing
var testUsers = []User{
	{
		Name: "testPerson1", 
		Password: "123",
		Email: "test1@gmail.com",
	},
	{
		Name: "testPerson2", 
		Password: "abc",
		Email: "test2@gmail.com",
	},
	{
		Name: "admin", 
		Password: "1234",
		Email: "admin@test.com",
	},
}

var err error

func TestCreateSession(t *testing.T) {
	setup()
	fmt.Println("Testing create session ...")
	if err = testUsers[0].Create(); err != nil {
		t.Error(err, "can not create user")
	}
	if session, err := testUsers[0].CreateSession(); err != nil {
		t.Error(err, "can not create session")
	} else {
		fmt.Println(session)
	}
}

func TestUpdateSession(t *testing.T){
	setup()
	if err = testUsers[0].Create(); err != nil {
		t.Error(err, "Can not create user")
	}

	sessionInterval = 0
	fmt.Println("Testing renew session ... ")
	if err := UpdateSession(); err != nil {
		t.Error(err, "Can not renew session")
	}
}


func TestCreateUser(t *testing.T) {
	setup()
	fmt.Println("Testing create user ...")
	for _, user := range testUsers {
		if err = user.Create(); err != nil {
			t.Error(err, "Can not create user")
		}
	}
}

func TestDeleteUser(t *testing.T) {
	setup()
	fmt.Println("Testing delete user ...")
	for _, user := range testUsers {
		if err = user.Create(); err != nil {
			t.Error(err, "Can not create user")
		}
	}
	for _, user := range testUsers {
		if err = user.Delete(); err != nil {
			t.Error(err, "Can not delete user")
		}
	}
}

func TestGetUsers(t *testing.T) {
	setup()
	fmt.Println("Testing get all users ...")
	for _, user := range testUsers {
		if err = user.Create(); err != nil {
			t.Error(err, "Can not create user")
		}
	}	

	var users = []User{}
	users, err = GetUsers()
	if err != nil {
		t.Error(err, "Can not get all users")
	}
	fmt.Println("Print all user information")
	for _, user := range users {
		fmt.Println(user)
	} 
}

func TestGetUserByEmail(t *testing.T) {
	setup()
	if err = testUsers[0].Create(); err != nil {
		t.Error(err, "Can not create user")
	}
	
	fmt.Println("Testing get user by email ...")
	email := testUsers[0].Email

	if user, err := GetUserByEmail(email); err != nil {
		t.Error(err, "Can not get user by email")
	} else {
		fmt.Println("Get the user by email: ", user)
	}
}

func TestGetUserByUuid(t *testing.T) {
	setup()
	if err = testUsers[0].Create(); err != nil {
		t.Error(err, "Can not create user")
	}
	
	fmt.Println("Testing get user by uuid ...")
	uuid := testUsers[0].Uuid

	if user, err := GetUserByUuid(uuid); err != nil {
		t.Error(err, "Can not get user by uuid")
	} else {
		fmt.Println("Get the user by uuid: ", user)
	}
}

