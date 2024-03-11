package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.CreateUserOutput, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (domain.UpdateUserOutput, error)
	GetById(ctx context.Context, input domain.GetUserByIDInput) (domain.GetUserByIdOutput, error)
	DeleteById(ctx context.Context, input domain.DeleteUserByIDInput) error
}

type UserRepository interface {
	Create(ctx context.Context, input domain.CreateUserInput) (domain.User, error)
	Update(ctx context.Context, input domain.UpdateUserInput) (domain.User, error)
	GetById(ctx context.Context, input domain.GetUserByIDInput) (domain.User, error)
	DeleteById(ctx context.Context, input domain.DeleteUserByIDInput) error
}
