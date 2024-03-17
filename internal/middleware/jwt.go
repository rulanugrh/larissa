package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/util"
)

type jwtclaim struct {
	ID     uint `json:"id"`
	RoleID uint `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, roleid uint) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	claims := &jwtclaim{
		ID:     id,
		RoleID: roleid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte(conf.App.Secret))
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}


func CheckToken(token string) (*jwtclaim, error) {
	conf := config.GetConfig()
	tokens, err := jwt.ParseWithClaims(token, &jwtclaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.App.Secret), util.Unauthorized("token missing or invalid")
	})

	if err != nil {
		return nil, util.Errors(err)
	}

	claim, valid := tokens.Claims.(*jwtclaim)
	if !valid {
		return nil, util.Unauthorized("sorry this token invalid")
	}

	return claim, nil
}