package model

type User struct {
	ID int `json:"id"`
	//ID    int    `gorm:"primaryKey;autoIncrement"
	Name  string `json:"name"`
	Email string `json:"email"`
}
