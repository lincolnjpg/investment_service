package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/domain"
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
	return domain.User{}, nil
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
