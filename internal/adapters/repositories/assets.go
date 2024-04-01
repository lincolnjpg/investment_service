package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type AssetRepository struct {
	db *pgx.Conn
}

func NewAssetRepository(db *pgx.Conn) AssetRepository {
	return AssetRepository{db: db}
}

func (repository AssetRepository) Create(ctx context.Context, input domain.CreateAssetInput) (domain.Asset, error) {
	var asset domain.Asset

	row := repository.db.QueryRow(
		ctx,
		`
			INSERT INTO assets(name, unit_price, rentability, due_date, ticker, asset_type_id, asset_index_id)
			VALUES($1, $2, $3, $4, $5, $6, $7)
			RETURNING id, name, unit_price, rentability, due_date, ticker, asset_type_id, asset_index_id;
		`,
		input.Name,
		input.UnitPrice,
		input.Rentability,
		input.DueDate,
		input.Ticker,
		input.AssetTypeId,
		input.AssetIndexId,
	)
	if err := row.Scan(&asset.Id, &asset.Name, &asset.UnitPrice, &asset.Rentability, &asset.DueDate, &asset.Ticker, &asset.AssetTypeId, &asset.AssetIndexId); err != nil {
		return asset, infra.NewAPIError(fmt.Sprintf("could not create a new asset: %s", err.Error()), http.StatusInternalServerError)
	}

	return asset, nil
}
