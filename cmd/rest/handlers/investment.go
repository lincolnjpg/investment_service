package handlers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/unrolled/render"
)

func CreateInvestmentHandler(investmentService ports.InvestmentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body dtos.CreateInvestmentInput
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

		user, err := investmentService.CreateInvestment(ctx, body)
		if err != nil {
			customerror := err.(customerror.APIError)
			render.JSON(w, customerror.StatusCode, customerror.ToMap())

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": user})
	}
}

func GetInvestmentByIdHandler(investmentService ports.InvestmentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		id, _ := uuid.Parse(chi.URLParam(r, "id"))
		body := dtos.GetInvestmentByIdInput{
			Id: id,
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

		user, err := investmentService.GetInvestmentById(ctx, body)
		if err != nil {
			customerror := err.(customerror.APIError)
			render.JSON(w, customerror.StatusCode, customerror.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": user})
	}
}
