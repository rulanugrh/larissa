package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/util"
)

type jwtclaim struct {
	ID     uint   `json:"id"`
	RoleID uint   `json:"role_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, roleid uint, email string) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	claims := &jwtclaim{
		ID:     id,
		RoleID: roleid,
		Email:  email,
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

func AdminVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig()
		token := r.Header.Get("Authorization")

		if token == "" {
			response := util.WriteJSON(util.Unauthorized("token missing"))
			w.WriteHeader(401)
			w.Write(response)
			return
		}

		claim, err := CheckToken(token)
		if err != nil {
			response := util.WriteJSON(util.Unauthorized(err.Error()))
			w.WriteHeader(401)
			w.Write(response)
			return
		}

		if claim.Email != conf.Admin.Email {
			response := util.WriteJSON(util.Forbidden("you are not admin"))
			w.WriteHeader(403)
			w.Write(response)
			return
		}

		if claim.ExpiresAt.Unix() < time.Now().Unix() {
			response := util.WriteJSON(util.Forbidden("token expired"))
			w.WriteHeader(403)
			w.Write(response)
			return
		}
	})
}

func GeneralVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			response := util.WriteJSON(util.Unauthorized("token missing"))
			w.WriteHeader(401)
			w.Write(response)
			return
		}

		claim, err := CheckToken(token)
		if err != nil {
			response := util.WriteJSON(util.Unauthorized(err.Error()))
			w.WriteHeader(401)
			w.Write(response)
			return
		}

		if claim.ExpiresAt.Unix() < time.Now().Unix() {
			response := util.WriteJSON(util.Forbidden("token expired"))
			w.WriteHeader(403)
			w.Write(response)
			return
		}
	})
}

func GetUserID(r *http.Request) (uint, error) {
	token := r.Header.Get("Authorization")
	claim, err := CheckToken(token)
	if err != nil {
		return 0, err
	}

	return claim.ID, nil
}
