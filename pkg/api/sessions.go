package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// authInfo - структура для хранения данных аутентификации
type authInfo struct {
	User string
	Pass string
}

/*
authSession - метод для аутентификации пользователя
и сохранения данных сессии в cookie:

Читаем тело запроса -> Декодируем из JSON в authInfo -> Если все совпадает, сохраняем в cookie
*/
func (api *Api) authSession(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var auth authInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if auth.User == "admin" && auth.Pass == "admin" {
		session, _ := api.Store.Get(r, "session-cookie")
		session.Values["User"] = auth.User
		session.Values["Pass"] = auth.Pass
		session.Values["Authenticated"] = true
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
