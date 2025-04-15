package models

// Các trường được viết hoa để có thể truy cập từ package khác
type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
