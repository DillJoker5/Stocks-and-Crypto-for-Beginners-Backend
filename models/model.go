package model

import (
	"time"
)

type UserJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Data []UserTable `json: "data"`
}

type ApiFavoritesJsonResponse struct {
	Type string `json: "type'`
	Message string `json: "message"`
	Data []ApiFavoritesTable `json: "data"`
}

type ThreadJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Data []ThreadTable `json: "data"`
}

type ResponsesJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Data []ResponseTable `json: "data"`
}

type ThreadFavoritesJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Data []ThreadFavoritesTable `json: "data"`
}

type UserTable struct {
	UserId int32 `json: "user_id"`
	Email string `json: "email"`
	Username string `json: "username"`
	Password string `json: "password"`
	ThreadId int32 `json: "thread_id"`
	ResponseId int32 `json: "response_id"`
}

type ApiFavoritesTable struct {
	ApiFavoritesId int32 `json: "api_favorites_id"`
	UserId int32 `json: "user_id"`
	StockId string `json: "stock_id"`
	ApiUrl string `json: "api_url"`
}

type ThreadTable struct {
	ThreadId int32 `json: "thread_id"`
	UserId int32 `json: "user_id"`
	ResponseId int32 `json: "response_id"`
	Name string `json: "name"`
	Description string `json: "description"`
	DateCreated time.Time `json: "date_created"`
}

type ResponseTable struct {
	ResponseId int32 `json: "response_id"`
	UserId int32 `json: "user_id"`
	ThreadId int32 `json: "thread_id"`
	Reply string `json: "reply"`
}

type ThreadFavoritesTable struct {
	ThreadFavoritesId int32 `json: "thread_favorites_id`
	UserId int32 `json: "user_id"`
}
