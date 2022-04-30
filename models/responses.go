package model

type Response struct {
	UserId int32 `json: "user_id"`
	ThreadId int32 `json: "thread_id"`
	Description string `json: "description"`
	//DateCreated time.Time `json: "date_created"`
	DateCreated string `json: "date_created"`
}

type ResponseTable struct {
	ResponseId int32 `json: "response_id"`
	UserId int32 `json: "user_id"`
	ThreadId int32 `json: "thread_id"`
	Description string `json: "description"`
	//DateCreated time.Time `json: "date_created"`
	DateCreated string `json: "date_created"`
}