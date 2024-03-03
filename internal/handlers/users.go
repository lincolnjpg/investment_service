package handlers

import (
	"context"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

func CreateUserHandle(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userService.Create(context.Background(), domain.CreateUserInput{})
	}
}
