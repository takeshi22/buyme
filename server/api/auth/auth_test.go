package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/buyme/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	var actual models.User

	user := models.User{
		Name:     "test user",
		Email:    "test@email.com",
		Password: "test",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/signUp", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = models.InitDatabase()
	assert.NoError(t, err)

	models.Db.AutoMigrate(&models.User{})

	SignUp(c)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)

	assert.Equal(t, user.Name, actual.Name)
}

func TestLogin(t *testing.T) {
	user := LoginPayload{
		Email:    "test@email.com",
		Password: "test",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Login(c)

	assert.Equal(t, 400, w.Code)
}
