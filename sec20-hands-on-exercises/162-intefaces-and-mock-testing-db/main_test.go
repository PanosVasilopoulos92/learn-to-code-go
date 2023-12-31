package main

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	fmt.Println("Test of func 'TestGetUser' has started.")
	md := &MockDatastore{
		Users: map[int]User{
			1: {Id: 3, Username: "Viator"},
		},
	}

	s := &Service{
		ds: md,
	}

	user, err := s.GetUser(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if user.Username != "Viator" {
		t.Errorf("Error: Got '%s' but was expecting '%s'.", user.Username, "Viator")
	}

	fmt.Println("------")
}

func TestInsertUser(t *testing.T) {
	fmt.Println("Test of func 'TestInsertUser' has started.")
	md := &MockDatastore{
		Users: map[int]User{},
	}

	s := &Service{
		ds: md,
	}

	user1 := User{
		Id:       1,
		Username: "Joker",
	}

	err := s.InsertUser(user1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Println("------")
}
