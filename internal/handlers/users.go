package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/rotisserie/eris"
	"github.com/unrolled/render"
)

func CreateUserHandle(ctx context.Context, logger *slog.Logger, userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)

		render := render.New()

		user, err := userService.Create(body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, eris.ToJSON(apiError.Err, true))
			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}
