package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(data any) []byte {
	response, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	
	return response
}

func SetCookie(name string, value *string, w http.ResponseWriter) error {
	cookie := http.Cookie{
		Name: name,
		Value: *value,
		MaxAge: 900,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
		Path: "/",
	}

	valid := cookie.Valid()
	if valid != nil {
		return Errors(valid)
	}

	http.SetCookie(w, &cookie)

	return nil
}