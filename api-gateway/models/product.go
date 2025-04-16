package models

type Product struct {
	Id     int32   `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	UserId int32   `json:"user_id"`
	User   *User   `json:"user,omitempty"` // Nếu User là nil, nó sẽ không được trả về trong JSON
}
