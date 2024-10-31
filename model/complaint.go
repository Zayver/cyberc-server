package model
type Complaint struct{
	ID uint `gorm:"column:id; primary_key; not null" json:"id"`
	Name string `gorm:"column:name; not null" json:"name"`
	Surname string `gorm:"column:surname; not null" json:"surName"`
}