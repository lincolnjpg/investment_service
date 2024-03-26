package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type AssetTypeRepository struct {
	db *pgx.Conn
}

func NewAssetTypeRepository(db *pgx.Conn) AssetTypeRepository {
	return AssetTypeRepository{db: db}
}

func (repository AssetTypeRepository) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.AssetType, error) {
	var assetType domain.AssetType

	row := repository.db.QueryRow(
		ctx,
		`
			INSERT INTO asset_types(name, description, index_id, class)
			VALUES($1, $2, $3, $4)
			RETURNING id, name, description, index_id, class;
		`,
		input.Name,
		input.Description,
		input.IndexId,
		input.Class,
	)
	if err := row.Scan(&assetType.Id, &assetType.Name, &assetType.Description, &assetType.IndexId, &assetType.Class); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not create a new asset type: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}

func (repository AssetTypeRepository) GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.AssetType, error) {
	var assetType domain.AssetType

	row := repository.db.QueryRow(
		ctx,
		`
			SELECT
				id, name, description, index_id, class
			FROM
				asset_types
			WHERE
				id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&assetType.Id, &assetType.Name, &assetType.Description, &assetType.IndexId, &assetType.Class); err != nil {
		if err == pgx.ErrNoRows {
			return assetType, infra.NewAPIError(fmt.Sprintf("asset type not found: %s", err.Error()), http.StatusNotFound)
		}

		err := infra.NewAPIError(fmt.Sprintf("could not get asset type from database: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}

func (repository AssetTypeRepository) UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.AssetType, error) {
	var assetType domain.AssetType

	row := repository.db.QueryRow(
		ctx,
		`
			UPDATE asset_types
			SET name = $2, description = $3, index_id = $4, class = $5
			WHERE id = $1
			RETURNING id, name, description, index_id, class;
		`,
		input.Id,
		input.Name,
		input.Description,
		input.IndexId,
		input.Class,
	)

	if err := row.Scan(&assetType.Id, &assetType.Name, &assetType.Description, &input.IndexId, &input.Class); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not update asset type: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}
