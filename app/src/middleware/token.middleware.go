package middleware

import (
	"fmt"
	"net/http"

	database "main/app/database/repository"

	"gorm.io/gorm"
)

type IUserRepository interface {
	SaveUser(uid string) *gorm.DB
}

var userRepository IUserRepository = &database.UserRepository{}

func TokenGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		uid := req.Header.Get("uid")
		if uid == "" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userRepository.SaveUser(uid)

		fmt.Printf("UID: %s\n", uid)
		next.ServeHTTP(res, req)
	})
}
