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

type investmentRepository struct {
	db *pgx.Conn
}

func (r investmentRepository) CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (entities.Investment, error) {
	var investment entities.Investment

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return investment, err
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
			INSERT INTO investments(user_id, asset_id, quantity, purchase_date, status, type)
			VALUES($1, $2, $3, $4, $5, $6)
			RETURNING id, user_id, asset_id, quantity, purchase_date, status, type;
		`,
		input.UserId,
		input.AssetId,
		input.Quantity,
		time.Now(),
		enum.Pending,
		input.Type,
	)
	if err := row.Scan(&investment.Id, &investment.UserId, &investment.AssetId, &investment.Quantity, &investment.PurchaseDate, &investment.Status, &investment.Type); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not create a new investment: %s", err.Error()), http.StatusInternalServerError)

		return investment, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := customerror.NewAPIError("could not commit transaction", http.StatusInternalServerError)

		return investment, err
	}

	return investment, nil
}

func (r investmentRepository) GetInvestmentById(ctx context.Context, input dtos.GetInvestmentByIdInput) (entities.Investment, error) {
	var investment entities.Investment

	row := r.db.QueryRow(
		ctx,
		`
			SELECT 
				id,
				user_id,
				asset_id,
				quantity,
				purchase_date,
				status,
				message
			FROM
				investments
			WHERE
				id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&investment.Id, &investment.UserId, &investment.AssetId, &investment.Quantity, &investment.PurchaseDate, &investment.Status, &investment.Message); err != nil {
		if err == pgx.ErrNoRows {
			return investment, customerror.NewAPIError(fmt.Sprintf("investment not found: %s", err.Error()), http.StatusNotFound)
		}

		err := customerror.NewAPIError(fmt.Sprintf("could not get investment from database: %s", err.Error()), http.StatusInternalServerError)

		return investment, err
	}

	return investment, nil
}

func (r investmentRepository) UpdateInvestmentById(ctx context.Context, input dtos.UpdateInvestmentByIdInput) (entities.Investment, error) {
	var investment entities.Investment

	row := r.db.QueryRow(
		ctx,
		`
			UPDATE
				investments
			SET
				status = $2,
				message = $3,
				updated_at = NOW()
			WHERE
				id = $1
			RETURNING
				id,
				user_id,
				asset_id,
				quantity,
				status,
				type,
				purchase_date,
				message;
		`,
		input.Id,
		input.Status,
		input.Message,
	)

	if err := row.Scan(&investment.Id, &investment.UserId, &investment.AssetId, &investment.Quantity, &investment.Status, &investment.Type, &investment.PurchaseDate, &investment.Message); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not update investment: %s", err.Error()), http.StatusInternalServerError)

		return investment, err
	}

	return investment, nil
}

func NewInvestmentRepository(db *pgx.Conn) investmentRepository {
	return investmentRepository{db: db}
}
