package repositories

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
)

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) userRepository {
	return userRepository{db: db}
}

func (repository userRepository) Create(ctx context.Context, input dtos.CreateUserInput) (entities.User, error) {
	var user entities.User

	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return user, err
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
			INSERT INTO users(name, investor_profile)
			VALUES($1, $2)
			RETURNING id, name, investor_profile;
		`,
		input.Name,
		input.InvestorProfile,
	)
	if err := row.Scan(&user.Id, &user.Name, &user.InvestorProfile); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not create a new user: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := customerror.NewAPIError("could not commit transaction", http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (repository userRepository) UpdateById(ctx context.Context, input dtos.UpdateUserInput) (entities.User, error) {
	var user entities.User

	row := repository.db.QueryRow(
		ctx,
		`
			UPDATE users
			SET name = $2, investor_profile = $3
			WHERE id = $1
			RETURNING id, name, investor_profile;
		`,
		input.Id,
		input.Name,
		input.InvestorProfile,
	)

	if err := row.Scan(&user.Id, &user.Name, &user.InvestorProfile); err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not update user: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (repository userRepository) GetById(ctx context.Context, input dtos.GetUserByIdInput) (entities.User, error) {
	var user entities.User

	row := repository.db.QueryRow(
		ctx,
		`
			SELECT * FROM users
			WHERE id = $1;
		`,
		input.Id,
	)
	if err := row.Scan(&user.Id, &user.Name, &user.InvestorProfile); err != nil {
		if err == pgx.ErrNoRows {
			return user, customerror.NewAPIError(fmt.Sprintf("user not found: %s", err.Error()), http.StatusNotFound)
		}

		err := customerror.NewAPIError(fmt.Sprintf("could not get user from database: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (repository userRepository) DeleteById(ctx context.Context, input dtos.DeleteUserByIDInput) error {
	_, err := repository.db.Exec(
		ctx,
		`
			DELETE FROM users
			WHERE id = $1;
		`,
		input.Id,
	)

	if err != nil {
		err := customerror.NewAPIError(fmt.Sprintf("could not delete user: %s", err.Error()), http.StatusInternalServerError)

		return err
	}

	return nil
}
