package model

import (
	"time"
)

type UserJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Data []UserTable `json: "data"`
}

type FavoritesJsonResponse struct {
	Type string `json: "type'`
	Message string `json: "message"`
	Data []FavoritesTable `json: "data"`
}

type ThreadJsonResponse struct {
	Type string `json: "type'`
	Message string `json: "message"`
	Data []ThreadTable `json: "data"`
}

type UserTable struct {
	UserId int32 `json: "userid"`
	Email string `json: "email"`
	Username string `json: "username"`
	Password string `json: "password"`
}

type FavoritesTable struct {
	FavoritesId int32 `json: "favoritesid"`
	UserId int32 `json: "userid"`
	StockId string `json: "stockid"`
	ApiUrl string `json: "apiurl"`
}

type ThreadTable struct {
	ThreadId int32 `json: "threadid"`
	UserId int32 `json: "userid"`
	Name string `json: "name"`
	Description string `json: "description"`
	DateCreated time.Time `json: "datecreated"`
}
