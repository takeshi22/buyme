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
	var userResult User

	err := InitDatabase()
	if err != nil {
		t.Error(err)
	}

	user := User{
		Name:     "test user",
		Email:    "test@email.com",
		Password: "test",
		Nickname: "tu",
	}
	err = user.CreateUser()
	assert.NoError(t, err)

	Db.Where("email = ?", user.Email).Find(&userResult)
	Db.Unscoped().Where("email = ?", user.Email).Delete(&user)

	assert.Equal(t, "test@email.com", userResult.Email)
}
