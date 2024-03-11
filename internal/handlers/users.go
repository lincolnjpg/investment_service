package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/unrolled/render"
)

func CreateUserHandler(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
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

func GetUserByIDHandler(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		body := domain.GetUserByIDInput{
			ID: chi.URLParam(r, "id"),
		}

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

		user, err := userService.GetById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": user})
	}
}

func UpateUserByIdHandler(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.UpdateUserInput
		json.NewDecoder(r.Body).Decode(&body)
		body.ID = chi.URLParam(r, "id")

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

		user, err := userService.UpdateById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": user})
	}
}

func DeleteUserByIDHandler(userService ports.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		body := domain.DeleteUserByIDInput{
			ID: chi.URLParam(r, "id"),
		}

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

		err = userService.DeleteById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusNoContent, map[string]interface{}{})
	}
}
