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
	if err := row.Scan(&assetType.ID, &assetType.Name, &assetType.Description, &assetType.Class); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not create a new asset type: %s", err.Error()), http.StatusInternalServerError)

		return assetType, err
	}

	return assetType, nil
}
