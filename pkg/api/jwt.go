package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (api *Api) authJWT(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var auth authInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if auth.User == "admin" && auth.Pass == "admin" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"usr": auth.User,
			"nbf": time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte("a-string-secret-at-least-256-bits-long"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
	}
}
