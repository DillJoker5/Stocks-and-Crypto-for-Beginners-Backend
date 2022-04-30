package model

type Thread struct {
	UserId int32 `json: "user_id"`
	ResponseId int32 `json: "response_id"`
	Name string `json: "name"`
	Description string `json: "description"`
	//DateCreated time.Time `json: "date_created"`
	DateCreated string `json: "date_created"`
}

type ThreadTable struct {
	ThreadId int32 `json: "thread_id"`
	UserId int32 `json: "user_id"`
	ResponseId int32 `json: "response_id"`
	Name string `json: "name"`
	Description string `json: "description"`
	//DateCreated time.Time `json: "date_created"`
	DateCreated string `json: "date_created"`
}