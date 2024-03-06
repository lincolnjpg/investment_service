package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.CreateUserOutput, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (domain.UpdateUserOutput, error)
	GetById(ctx context.Context, id uuid.UUID) (domain.GetUserByIdOutput, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type UserRepository interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.User, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (domain.User, error)
	GetById(ctx context.Context, id uuid.UUID) (domain.User, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
}
