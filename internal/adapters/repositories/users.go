package repositories

import (
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

func (r UserRepository) Create(input domain.CreateUserInput) (domain.User, error) {
	return domain.User{}, nil
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
