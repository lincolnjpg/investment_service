package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/unrolled/render"
)

func CreateUserHandle(ctx context.Context, userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body domain.CreateUserInput
		json.NewDecoder(r.Body).Decode(&body)

		render := render.New()

		user, err := userService.Create(ctx, body)
		if err != nil {
			fmt.Println(err)
			render.JSON(w, http.StatusInternalServerError, infra.NewAPIError(http.StatusInternalServerError, err.Error()))
			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}
