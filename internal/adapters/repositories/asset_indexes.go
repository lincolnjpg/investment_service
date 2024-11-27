package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
)

type assetIndexRepository struct {
	db *pgx.Conn
}

func NewAssetIndexRepository(db *pgx.Conn) assetIndexRepository {
	return assetIndexRepository{db: db}
}

func (repository assetIndexRepository) CreateAssetIndex(ctx context.Context, input dtos.CreateAssetIndexInput) (entities.AssetIndex, error) {
	var assetIndex entities.AssetIndex

	row := repository.db.QueryRow(
		ctx,
		`
			INSERT INTO asset_indexes(name, acronym)
			VALUES($1, $2)
			RETURNING id, name, acronym;
		`,
		input.Name,
		input.Acronym,
	)
	if err := row.Scan(&assetIndex.Id, &assetIndex.Name, &assetIndex.Acronym); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not create a new asset index: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}

func (repository assetIndexRepository) GetAssetIndexById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (entities.AssetIndex, error) {
	var assetIndex entities.AssetIndex

	row := repository.db.QueryRow(
		ctx,
		`
			SELECT
				id,
				name,
				acronym	
			FROM
				asset_indexes
			WHERE
				id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&assetIndex.Id, &assetIndex.Name, &assetIndex.Acronym); err != nil {
		if err == pgx.ErrNoRows {
			return assetIndex, customerror.NewAPIError(fmt.Sprintf("asset index not found: %s", err.Error()), http.StatusNotFound)
		}

		err := customerror.NewAPIError(fmt.Sprintf("could not get asset index from database: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}

func (repository assetIndexRepository) UpdateAssetIndexById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (entities.AssetIndex, error) {
	var assetIndex entities.AssetIndex

	row := repository.db.QueryRow(
		ctx,
		`
			UPDATE
				asset_indexes
			SET
				name = $2,
				acronym = $3,
				updated_at = NOW()
			WHERE
				id = $1
			RETURNING
				id,
				name,
				acronym;
		`,
		input.Id,
		input.Name,
		input.Acronym,
	)

	if err := row.Scan(&assetIndex.Id, &assetIndex.Name, &input.Acronym); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not update asset index: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}

func (repository assetIndexRepository) DeleteAssetIndexById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error {
	_, err := repository.db.Exec(
		ctx,
		`
			DELETE FROM
				asset_indexes
			WHERE
				id = $1;
		`,
		input.Id,
	)

	if err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not delete asset index: %s", err.Error()), http.StatusInternalServerError)

		return err
	}

	return nil
}
