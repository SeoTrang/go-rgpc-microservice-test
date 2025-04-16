package models

type Order struct {
	Id     int32  `json:"id"`
	Code   string `json:"code"`
	UserId int32  `json:"user_id"`
	User   *User  `json:"user,omitempty"`
}
