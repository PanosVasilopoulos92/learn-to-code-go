package main

import (
	"fmt"
	"log"
)

/*
Interfaces in Go are a powerful tool for abstraction and can be especially useful when you
want to write unit tests for functions that interact with a database. By creating an interface
that describes the behavior of the database interactions your code needs, you can swap out
the real database for a mock one in your tests.
*/

// User represents a user with an ID and a username.
type User struct {
	Id       int
	Username string
}

// MockDatastore is a temporary service that stores retrievable data.
type MockDatastore struct {
	Users map[int]User
}

// GetUser function was created in order to retrieve a user from the db or the data structure in this case.
func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf(" User with ID: %v was not found. ", id)
	}
	return user, nil
}

// InsertUser function was created in order to insert a new user to db or the data structure in this case.
func (md MockDatastore) InsertUser(user User) error {
	_, ok := md.Users[user.Id]
	if ok {
		return fmt.Errorf(" Username '%v' already exist in DB. ", user.Username)
	}

	md.Users[user.Id] = user
	return nil
}

// Datastore defines an interface for storing retrievable data.
type Datastore interface {
	GetUser(id int) (User, error)
	InsertUser(user User) error
}

// Service represents a service that stores retrievable data.
type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) InsertUser(user User) error {
	return s.ds.InsertUser(user)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	service := Service{
		ds: db,
	}

	user1 := User{
		Id:       1,
		Username: "Viator",
	}

	user2 := User{
		Id:       2,
		Username: "Vectronium",
	}

	err := service.InsertUser(user1)
	if err != nil {
		log.Fatalf("error %s", err)
	} else {
		fmt.Println("User succefully inserted.")
	}

	userRetrived1, err := service.GetUser(user1.Id)
	if err != nil {
		log.Fatalf("error %s", err)
	} else {
		fmt.Println(userRetrived1)
	}

	err = service.InsertUser(user2)
	if err != nil {
		log.Fatalf("error: %s", err)
	} else {
		fmt.Println("User succefully inserted.")
	}

	userRetrived2, err := service.GetUser(user2.Id)
	if err != nil {
		log.Fatalf("error %s", err)
	} else {
		fmt.Println(userRetrived2)
	}
}
