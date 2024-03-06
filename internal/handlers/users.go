package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/unrolled/render"
)

func CreateUserHandle(logger *httplog.Logger, userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)
		httplog.LogEntrySetField(ctx, "requestInput", slog.AnyValue(body))

		render := render.New()

		user, err := userService.Create(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, map[string]interface{}{"error": apiError.Err.Error()})

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}
