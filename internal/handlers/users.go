package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

func CreateUserHandle(ctx context.Context, userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)
		userService.Create(context.Background(), body)
	}
}
