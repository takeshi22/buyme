package auth

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/buyme/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

func GenerateToken(userid uint) (*TokenDetails, error) {
	tokenDetail := &TokenDetails{}

	tokenDetail.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetail.AccessUuid = uuid.NewV4().String()

	tokenDetail.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetail.RefreshUuid = uuid.NewV4().String()

	var err error

	// generate access token
	secret := os.Getenv("AUTH_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokenDetail.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = tokenDetail.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetail.AccessToken, err = at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshSecret := os.Getenv("REFRESH_SECRET")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetail.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = tokenDetail.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenDetail.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func CreateAuth(userid uint, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := models.Client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := models.Client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{"msg": "could not bind user"})
		c.Abort()
		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{"msg": "error encrypt password"})
		c.Abort()
		return
	}

	err = user.CreateUser()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"msg": "could not create user"})
		c.Abort()
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid login",
		})
		c.Abort()
		return
	}

	result := models.Db.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(user.ID, token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
