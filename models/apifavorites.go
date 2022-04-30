package model

type ApiFavorite struct {
	UserId int32 `json: "user_id"`
	StockId string `json: "stock_id"`
	ApiUrl string `json: "api_url"`
}

type ApiFavoritesTable struct {
	ApiFavoritesId int32 `json: "api_favorites_id"`
	UserId int32 `json: "user_id"`
	StockId string `json: "stock_id"`
	ApiUrl string `json: "api_url"`
}