package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "test",
	}
	err := user.HashPassword(user.Password)
	emptyPassword := user.HashPassword("")

	assert.NoError(t, err)
	assert.Error(t, emptyPassword)

	os.Setenv("hashedPassword", user.Password)
}

func TestCreateUser(t *testing.T) {
	// var userResult User

	err := InitDatabase()
	if err != nil {
		t.Error(err)
	}
}
