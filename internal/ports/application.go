package ports

type Application interface {
	UserService
	AssetIndexService
	AssetService
	InvestmentService
}
