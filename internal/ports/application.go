package ports

type Application struct {
	UserService       UserService
	AssetService      AssetService
	AssetIndexService AssetIndexService
}
