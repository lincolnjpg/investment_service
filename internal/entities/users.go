package entities

import "github.com/lincolnjpg/investment_service/internal/enum"

type User struct {
	Id              string
	Name            string
	InvestorProfile enum.InvestorProfileEnum
}
