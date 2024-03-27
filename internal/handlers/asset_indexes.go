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

func CreateAssetIndexHandler(assetIndexesService ports.AssetIndexService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.CreateAssetIndexInput
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

		assetType, err := assetIndexesService.Create(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": assetType})
	}
}

func GetAssetIndexByIdHandler(assetIndexesService ports.AssetIndexService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		body := domain.GetAssetIndexByIdInput{
			Id: chi.URLParam(r, "id"),
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

		assetType, err := assetIndexesService.GetById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": assetType})
	}
}

func UpdateAssetIndexByIdHandler(assetIndexesService ports.AssetIndexService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.UpdateAssetIndexByIdInput
		json.NewDecoder(r.Body).Decode(&body)
		body.Id = chi.URLParam(r, "id")

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

		assetType, err := assetIndexesService.UpdateById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": assetType})
	}
}
