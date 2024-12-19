package entities

type SearchlistUsecase interface {
	Search() (*BeefRes, error)
}

type SearchlistRepository interface {
	SearchListDb() (*string, error)
}

type BeefRes struct {
	Beef map[string]int32 `json:"beef"`
}