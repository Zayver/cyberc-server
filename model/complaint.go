package model

import "github.com/google/uuid"

type Complaint struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string `gorm:"column:name; not null" json:"name"`
	SecondName string `gorm:"column:secondName" json:"secondName"`
	SurName string `gorm:"column:surName; not null" json:"surName"`
	SecondSurName string `gorm:"column:secondSurName; not null" json:"secondSurName"`
	Cellphone string `gorm:"column:cellphone; not null" json:"cellphone"`
	Email string `gorm:"column:email; not null" json:"email"`
	Cc string `gorm:"column:cc; not null" json:"cc"`
	Description string `gorm:"column:description; not null" json:"description"`
	Type ComplaintType `gorm:"column:type; not null" json:"type"`
}