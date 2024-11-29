package repositories

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
	"github.com/lincolnjpg/investment_service/internal/enum"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
)

type userAssetRepository struct {
	db *pgx.Conn
}

func (r userAssetRepository) CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (entities.UserAsset, error) {
	var userAsset entities.UserAsset

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return userAsset, err
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			log.Println("could not rollback transaction:", err)
		}
	}()

	row := tx.QueryRow(
		ctx,
		`
			INSERT INTO users_assets(user_id, asset_id, quantity, purchase_date, status)
			VALUES($1, $2, $3, $4, $5)
			RETURNING id, user_id, asset_id, quantity, purchase_date, status;
		`,
		input.UserId,
		input.AssetId,
		input.Quantity,
		time.Now(),
		enum.Pending,
	)
	if err := row.Scan(&userAsset.Id, &userAsset.UserId, &userAsset.AssetId, &userAsset.Quantity, &userAsset.PuchaseDate, &userAsset.Status); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not create a new user asset: %s", err.Error()), http.StatusInternalServerError)

		return userAsset, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := customerror.NewAPIError("could not commit transaction", http.StatusInternalServerError)

		return userAsset, err
	}

	return userAsset, nil
}

func NewUserAssetRepository(db *pgx.Conn) userAssetRepository {
	return userAssetRepository{db: db}
}
