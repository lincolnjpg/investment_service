package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type assetRepository struct {
	db *pgx.Conn
}

func NewAssetRepository(db *pgx.Conn) assetRepository {
	return assetRepository{db: db}
}

func (repository assetRepository) Create(ctx context.Context, input dtos.CreateAssetInput) (entities.Asset, error) {
	var asset entities.Asset

	row := repository.db.QueryRow(
		ctx,
		`
			INSERT INTO assets(name, unit_price, rentability, due_date, ticker, type, asset_index_id)
			VALUES($1, $2, $3, $4, $5, $6, $7)
			RETURNING id, name, unit_price, rentability, due_date, ticker, type, asset_index_id;
		`,
		input.Name,
		input.UnitPrice,
		input.Rentability,
		input.DueDate,
		input.Ticker,
		input.Type,
		input.AssetIndexId,
	)
	if err := row.Scan(&asset.Id, &asset.Name, &asset.UnitPrice, &asset.Rentability, &asset.DueDate, &asset.Ticker, &asset.Type, &asset.AssetIndexId); err != nil {
		return asset, infra.NewAPIError(fmt.Sprintf("could not create a new asset: %s", err.Error()), http.StatusInternalServerError)
	}

	return asset, nil
}

func (repository assetRepository) GetById(ctx context.Context, input dtos.GetAssetByIdInput) (entities.Asset, error) {
	var asset entities.Asset

	row := repository.db.QueryRow(
		ctx,
		`
			SELECT
				id, name, unit_price, rentability, due_date, ticker, type, asset_index_id
			FROM
				assets
			WHERE
				id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&asset.Id, &asset.Name, &asset.UnitPrice, &asset.Rentability, &asset.DueDate, &asset.Ticker, &asset.Type, &asset.AssetIndexId); err != nil {
		if err == pgx.ErrNoRows {
			return asset, infra.NewAPIError(fmt.Sprintf("asset not found: %s", err.Error()), http.StatusNotFound)
		}

		err := infra.NewAPIError(fmt.Sprintf("could not get asset from database: %s", err.Error()), http.StatusInternalServerError)

		return asset, err
	}

	return asset, nil
}

func (repository assetRepository) UpdateById(ctx context.Context, input dtos.UpdateAssetByIdInput) (entities.Asset, error) {
	var asset entities.Asset

	row := repository.db.QueryRow(
		ctx,
		`
			UPDATE assets
			SET name = $2, unit_price = $3, rentability = $4, due_date = $5, ticker = $6, type = $7, asset_index_id = $8
			WHERE id = $1
			RETURNING id, name, unit_price, rentability, due_date, ticker, type, asset_index_id;
		`,
		input.Id,
		input.Name,
		input.UnitPrice,
		input.Rentability,
		input.DueDate,
		input.Ticker,
		input.Type,
		input.AssetIndexId,
	)

	if err := row.Scan(&asset.Id, &asset.Name, &asset.UnitPrice, &asset.Rentability, &input.DueDate, &asset.Ticker, &asset.Type, &asset.AssetIndexId); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not update asset: %s", err.Error()), http.StatusInternalServerError)

		return asset, err
	}

	return asset, nil
}

func (repository assetRepository) DeleteById(ctx context.Context, input dtos.DeleteAssetByIdInput) error {
	_, err := repository.db.Exec(
		ctx,
		`
			DELETE FROM assets
			WHERE id = $1;
		`,
		input.Id,
	)

	if err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not delete asset: %s", err.Error()), http.StatusInternalServerError)

		return err
	}

	return nil
}
