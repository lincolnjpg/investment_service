package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type AssetIndexRepository struct {
	db *pgx.Conn
}

func NewAssetIndexRepository(db *pgx.Conn) AssetIndexRepository {
	return AssetIndexRepository{db: db}
}

func (repository AssetIndexRepository) Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.AssetIndex, error) {
	var assetIndex domain.AssetIndex

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
		err := infra.NewAPIError(fmt.Sprintf("could not create a new asset index: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}

func (repository AssetIndexRepository) GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.AssetIndex, error) {
	var assetIndex domain.AssetIndex

	row := repository.db.QueryRow(
		ctx,
		`
			SELECT * FROM asset_indexes
			WHERE id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&assetIndex.Id, &assetIndex.Name, &assetIndex.Acronym); err != nil {
		if err == pgx.ErrNoRows {
			return assetIndex, infra.NewAPIError(fmt.Sprintf("asset index not found: %s", err.Error()), http.StatusNotFound)
		}

		err := infra.NewAPIError(fmt.Sprintf("could not get user from database: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}

func (repository AssetIndexRepository) UpdateById(ctx context.Context, input domain.UpdateAssetIndexByIdInput) (domain.AssetIndex, error) {
	var assetIndex domain.AssetIndex

	row := repository.db.QueryRow(
		ctx,
		`
			UPDATE asset_indexes
			SET name = $2, acronym = $3
			WHERE id = $1
			RETURNING id, name, acronym;
		`,
		input.Id,
		input.Name,
		input.Acronym,
	)

	if err := row.Scan(&assetIndex.Id, &assetIndex.Name, &input.Acronym); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not update asset index: %s", err.Error()), http.StatusInternalServerError)

		return assetIndex, err
	}

	return assetIndex, nil
}
