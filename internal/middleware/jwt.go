package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/larissa/internal/config"
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
