package model

type User struct {
	Id       int    `json:"id" gorm:"primaryKey; autoIncrement"`
	Nama     string `json:"user"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	BaseModel
}
