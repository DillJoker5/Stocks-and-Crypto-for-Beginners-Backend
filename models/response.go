package model

type GenericJsonResponse struct {
	Type string `json: "type"`
	Message string `json: "message"`
}

type LoginJsonResponse struct {
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