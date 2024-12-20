package handlers

import (
	"encoding/json"
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

func CreateAssetIndexHandler(assetIndexesService ports.AssetIndexService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		var body dtos.CreateAssetIndexInput
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

		assetType, err := assetIndexesService.CreateAssetIndex(ctx, body)
		if err != nil {
			apiError := err.(customerror.APIError)
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

		id, _ := uuid.Parse(chi.URLParam(r, "id"))
		body := dtos.GetAssetIndexByIdInput{
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

		assetType, err := assetIndexesService.GetAssetIndexById(ctx, body)
		if err != nil {
			apiError := err.(customerror.APIError)
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

		var body dtos.UpdateAssetIndexByIdInput
		json.NewDecoder(r.Body).Decode(&body)
		id, _ := uuid.Parse(chi.URLParam(r, "id"))
		body.Id = id

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

		assetType, err := assetIndexesService.UpdateAssetIndexById(ctx, body)
		if err != nil {
			apiError := err.(customerror.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusOK, map[string]interface{}{"data": assetType})
	}
}

func DeleteAssetIndexByIDHandler(assetIndexesService ports.AssetIndexService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		render := render.New()

		id, _ := uuid.Parse(chi.URLParam(r, "id"))
		body := dtos.DeleteAssetIndexByIdInput{
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

		err = assetIndexesService.DeleteAssetIndexById(ctx, body)
		if err != nil {
			apiError := err.(customerror.APIError)
			render.JSON(w, apiError.StatusCode, apiError.ToMap())

			return
		}

		render.JSON(w, http.StatusNoContent, map[string]interface{}{})
	}
}
