package repositories

import (
	"context"
	"net/http"

	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

type UserRepository struct {
	logger *httplog.Logger
	db     *pgx.Conn
}

func NewUserRepository(logger *httplog.Logger, db *pgx.Conn) UserRepository {
	return UserRepository{
		logger: logger,
		db:     db,
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
		[]interface{}{input.Name, input.InvestorProfile}...,
	)
	if err := row.Scan(&user.Id, &user.Name, &user.InvestorProfile); err != nil {
		err := infra.NewAPIError(err.Error(), http.StatusInternalServerError)

		return user, err
	}

	return user, nil
}

func (r UserRepository) Update(ctx context.Context, input domain.UpdateUserInput) (domain.User, error) {
	return domain.User{}, nil
}

func (r UserRepository) GetById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (r UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
