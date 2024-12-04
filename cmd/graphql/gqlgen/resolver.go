package gqlgen

import "github.com/lincolnjpg/investment_service/internal/ports"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ports.UserService
	ports.AssetIndexService
	ports.AssetService
	ports.InvestmentService
}
