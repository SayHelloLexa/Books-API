package api

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
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

// JwtMiddleware - проверка валидности JWT-токена
func (api *Api) JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/authJWT" {
			next.ServeHTTP(w, r)
			return
		}

		// 1. Извлекаем токен из заголовка
		// {"Authorization": Bearer <token>}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Токена авторизации не найдено", http.StatusUnauthorized)
			return
		}

		// 2. Проверяем формат заголовка
		// ["Bearer", "<token>"]
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Формат заголовка неверен, ожидалось Bearer <token>", http.StatusUnauthorized)
			return
		}

		// 3. Кладем токен в переменную
		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("невалидный метод: %v", token.Header["alg"])
			}

			return []byte("a-string-secret-at-least-256-bits-long"), nil
		})
		if err != nil {
			http.Error(w, "Ошибка при парсинге токена", http.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//ctx := context.WithValue(r.Context(), "jwtClaims", claims)
			//r.WithContext(ctx)
			fmt.Printf("Данные токена JWT: %+v\n", claims)
		}

		next.ServeHTTP(w, r)
	})
}
