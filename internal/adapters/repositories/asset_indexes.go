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
