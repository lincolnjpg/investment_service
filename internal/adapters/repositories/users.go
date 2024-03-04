package repositories

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/rotisserie/eris"
)

type UserRepository struct {
	ctx    context.Context
	logger *slog.Logger
	db     *pgx.Conn
}

func NewUserRepository(ctx context.Context, logger *slog.Logger, db *pgx.Conn) UserRepository {
	return UserRepository{
		ctx:    ctx,
		logger: logger,
		db:     db,
	}
}

func (r UserRepository) Create(input domain.CreateUserInput) (domain.User, error) {
	var user domain.User

	row := r.db.QueryRow(
		r.ctx,
		`
		INSERT INTO users(name, investor_profile)
		VALUES($1, $2)
		RETURNING id, name, investor_profile;
		`,
		[]interface{}{input.Name, input.InvestorProfile}...,
	)
	if err := row.Scan(&user.Id, &user.Name, &user.InvestorProfile); err != nil {
		err := infra.NewAPIError(err.Error(), http.StatusInternalServerError)
		r.logger.Error("Error while persisting new user", "error", eris.ToJSON(err.Err, true))

		return user, err
	}

	return user, nil
}

func (r UserRepository) Update(input domain.UpdateUserInput) (domain.User, error) {
	return domain.User{}, nil
}

func (r UserRepository) GetById(id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (r UserRepository) DeleteById(id uuid.UUID) error {
	return nil
}
