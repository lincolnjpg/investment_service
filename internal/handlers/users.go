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

// validator como dep
func CreateUserHandle(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)

		err := body.Validate()
		if err != nil {
			render.JSON(w, http.StatusUnprocessableEntity,
				map[string]map[string]interface{}{
					"error": {
						"message": "input validation error",
						"fields":  err,
					},
				},
			)

			return
		}

		httplog.LogEntrySetField(ctx, "requestInput", slog.AnyValue(body))

		user, err := userService.Create(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}
