package ports

type Application interface {
	UserService
	AssetService
	AssetIndexService
}
