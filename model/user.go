package model

type User struct {
	//ID int `json:"id"`
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	//LastName string `json:"lastname"`
	Email string `json:"email"`
}
