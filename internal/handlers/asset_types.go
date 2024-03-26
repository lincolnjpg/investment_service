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

func CreateAssetTypeHandler(assetTypesService ports.AssetTypeService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.CreateAssetTypeInput
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

		assetType, err := assetTypesService.Create(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusCreated, map[string]interface{}{"data": assetType})
	}
}

func GetAssetTypeByIDHandler(assetTypeService ports.AssetTypeService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		body := domain.GetAssetTypeByIDInput{
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

		assetType, err := assetTypeService.GetById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": assetType})
	}
}

func UpdateAssetTypeByIdHandler(assetTypeService ports.AssetTypeService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body domain.UpdateAssetTypeByIdInput
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

		assetType, err := assetTypeService.UpdateById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": assetType})
	}
}

func DeleteAssetTypeByIDHandler(assetTypeService ports.AssetTypeService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		body := domain.DeleteAssetTypeByIdInput{
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

		err = assetTypeService.DeleteById(ctx, body)
		if err != nil {
			apiError := err.(infra.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusNoContent, map[string]interface{}{})
	}
}
