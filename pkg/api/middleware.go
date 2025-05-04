package api

import (
	"net/http"
)

// JsonHeaderMiddleware - устанавливает заголовок Content-Type: application/json
func JsonHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ApiSessionMiddleware - проверка валидности сеанса
func (api *Api) ApiSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/authSession" {
			next.ServeHTTP(w, r)
			return
		}

		session, _ := api.Store.Get(r, "session-cookie")

		if session.Values["Authenticated"] != true {
			http.Error(w, "доступ запрещен", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
