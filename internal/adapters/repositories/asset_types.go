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

func (r AssetTypeRepository) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.AssetType, error) {
	var assetType domain.AssetType

	row := r.db.QueryRow(
		ctx,
		`
			INSERT INTO asset_types(name, description, class)
			VALUES($1, $2, $3)
			RETURNING id, name, description, class;
		`,
		input.Name,
		input.Description,
		input.Class,
	)
	if err := row.Scan(&assetType.Id, &assetType.Name, &assetType.Description, &assetType.Class); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not create a new asset type: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}

func (r AssetTypeRepository) GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.AssetType, error) {
	var assetType domain.AssetType

	row := r.db.QueryRow(
		ctx,
		`
			SELECT
				id, name, description, class
			FROM
				asset_types
			WHERE
				id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&assetType.Id, &assetType.Name, &assetType.Description, &assetType.Class); err != nil {
		if err == pgx.ErrNoRows {
			return assetType, infra.NewAPIError(fmt.Sprintf("asset type not found: %s", err.Error()), http.StatusNotFound)
		}

		err := infra.NewAPIError(fmt.Sprintf("could not get asset type from database: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}
