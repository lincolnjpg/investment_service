package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

func respond(w http.ResponseWriter, statusCode int, serviceOutput interface{}) {
	responseBody := map[string]interface{}{}

	switch v := serviceOutput.(type) {
	case error:
		responseBody["error"] = v
	default:
		responseBody["data"] = v
	}

	w.Header().Add("Content-Type", "application/json")

	serializedResponseBody, err := json.Marshal(&responseBody)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while marshaling: %s", err.Error())))
	}

	w.WriteHeader(statusCode)
	w.Write(serializedResponseBody)
}

func CreateUserHandle(ctx context.Context, userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)

		user, err := userService.Create(context.Background(), body)
		if err != nil {
			respond(w, http.StatusInternalServerError, err)
		}

		respond(w, http.StatusCreated, &user)
	}
}
