package model

type UserTable struct {
	UserId int32 `json: "user_id"`
	Email string `json: "email"`
	Username string `json: "username"`
	Password string `json: "password"`
	ThreadId int32 `json: "thread_id"`
	ResponseId int32 `json: "response_id"`
}

type Session struct {
	SessionId int32 `json: "sessionid"`
	UserId int32 `json: "userid"`
	UserGuid string `json: "userguid"`
	IsActive bool `json: "isactive"`
}
