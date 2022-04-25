package model

type ResponseTable struct {
	ResponseId int32 `json: "response_id"`
	UserId int32 `json: "user_id"`
	ThreadId int32 `json: "thread_id"`
	Reply string `json: "reply"`
}