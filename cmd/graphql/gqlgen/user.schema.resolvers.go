package gqlgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/enum"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (string, error) {
	user, err := r.Resolver.UserService.Create(ctx, dtos.CreateUserInput{Name: input.Name, InvestorProfile: enum.InvestorProfileEnum(input.InvestorProfile)})
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*User, error) {
	u, err := r.Resolver.UserService.GetById(context.Background(), dtos.GetUserByIDInput{Id: id})
	if err != nil {
		return nil, err
	}

	return &User{ID: u.Id, Name: u.Name, InvestorProfile: int(u.InvestorProfile)}, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }