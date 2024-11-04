package request

import "github.com/zayver/cyberc-server/model"
type CreateComplaintRequest struct{
	Name string `json:"name" validate:"required"`
	SecondName string `json:"secondName" validate:"required"`
	SurName string `json:"surName" validate:"required"`
	SecondSurName string `json:"secondSurName" validate:"required"`
	Cellphone string `json:"cellphone" validate:"required,numeric"`
	Email string `json:"email" validate:"required,email"`
	Cc string `json:"cc" validate:"required,numeric"`
	Description string `json:"description" validate:"required"`
	Type model.ComplaintType `json:"type" validate:"required"` 
}