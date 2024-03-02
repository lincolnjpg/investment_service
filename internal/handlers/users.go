package handlers

import (
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

func CreateUserHandle(userService domain.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userService.Create(domain.CreateUserInput{})
	}
}
