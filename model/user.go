package model

type User struct{
	ID uint `gorm:"column:id; primary_key; not null" json:"id"`
	Username string `gorm:"column:username; unique; not null" json:"username"`
	Password string `gorm:"column:password; not null" json:"-"`
}