package user_test

import (
	"testing"
	"user"
)

func TestUserFullName(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	u.FirstName = "Peyton"
	u.LastName = "Manning"
	fullName := u.FullName()
	if fullName != "Peyton Manning" {
		t.Errorf("Expected '%s' to be Peyton Manning", fullName)
	}
}

func TestUserFullNameWithoutLastName(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	u.FirstName = "Peyton"
	fullName := u.FullName()
	if fullName != "Peyton" {
		t.Errorf("Expected '%s' to be Peyton", fullName)
	}
}

func TestUserFullNameWithoutFirstName(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	u.LastName = "Manning"
	fullName := u.FullName()
	if fullName != "Manning" {
		t.Errorf("Expected '%s' to be Manning", fullName)
	}
}

func TestUserFullNameWithNeither(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	fullName := u.FullName()
	if fullName != "" {
		t.Errorf("Expected '%s' to be empty", fullName)
	}
}
