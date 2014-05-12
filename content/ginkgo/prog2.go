package user_test

import (
	"testing"
	"user"
)

var fullNameCases = []struct {
	FirstName string
	LastName  string
	Result    string
}{
	{"Peyton", "Manning", "Peyton Manning"},
	{"Peyton", "", "Peyton"},
	{"", "Manning", "Manning"},
	{"", "", ""},
}

func TestUserFullName(t *testing.T) {
	for _, fullNameCase := range fullNameCases {
		u, err := user.New()
		if err != nil {
			t.Errorf("Got an unexpected error: %v", err)
		}
		u.FirstName = fullNameCase.FirstName
		u.LastName = fullNameCase.LastName
		fullName := u.FullName()
		if fullName != fullNameCase.Result {
			t.Errorf("Expected '%s' to equal '%s'", fullName, fullNameCase.Result)
		}
	}
}
