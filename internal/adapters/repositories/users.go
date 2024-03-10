package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(ctx context.Context, input domain.CreateUserInput) (domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(
		ctx,
		`
			INSERT INTO users(name, investor_profile)
			VALUES($1, $2)
			RETURNING id, name, investor_profile;
		`,
		input.Name,
		input.InvestorProfile,
	)
	if err := row.Scan(&user.ID, &user.Name, &user.InvestorProfile); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not create a new user: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (r UserRepository) Update(ctx context.Context, input domain.UpdateUserInput) (domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(
		ctx,
		`
			UPDATE users
			SET name = $2, investor_profile = $3
			WHERE id = $1
			RETURNING id, name, investor_profile;
		`,
		input.ID,
		input.Name,
		input.InvestorProfile,
	)

	if err := row.Scan(&user.ID, &user.Name, &user.InvestorProfile); err != nil {
		err := infra.NewAPIError(fmt.Sprintf("could not update user: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (r UserRepository) GetById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(
		ctx,
		`
			SELECT * FROM users
			WHERE id = $1;
		`,
		id.String(),
	)
	if err := row.Scan(&user.ID, &user.Name, &user.InvestorProfile); err != nil {
		if err == pgx.ErrNoRows {
			return user, infra.NewAPIError(fmt.Sprintf("user not found: %s", err.Error()), http.StatusNotFound)
		}

		err := infra.NewAPIError(fmt.Sprintf("could not get user from database: %s", err.Error()), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (r UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
