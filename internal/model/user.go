package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Username string `json:"login" gorm:"unique"`
	Password string `json:"password"`
	Session  string `json:"-"`
}
