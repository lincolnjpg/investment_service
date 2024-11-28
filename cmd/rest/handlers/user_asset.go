package handlers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/unrolled/render"
)

func CreateUserAssetHandler(userAssetService ports.UserAssetService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body dtos.CreateUserAssetInput
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Println(err)
			return
		}

		err = body.Validate()
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

		user, err := userAssetService.CreateUserAsset(ctx, body)
		if err != nil {
			customerror := err.(customerror.APIError)
			render.JSON(w, customerror.StatusCode, customerror.ToMap())

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}
