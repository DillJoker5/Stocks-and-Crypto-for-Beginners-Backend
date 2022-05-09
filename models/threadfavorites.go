package model

type ThreadFavoritesTable struct {
	ThreadFavoritesId int32 `json: "thread_favorites_id`
	UserId            int32 `json: "user_id"`
}
type ThreadFavorites struct {
	UserId int32 `json: "user_id"`
}
