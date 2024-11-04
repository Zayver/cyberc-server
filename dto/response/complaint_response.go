package response

import "github.com/google/uuid"

type ComplaintResponse struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	SecondName string `json:"secondName"`
	SurName string `json:"surName"`
	SecondSurName string `json:"secondSurName"`
	Cellphone string `json:"cellphone"`
	Email string `json:"email"`
	Cc string `json:"cc"`
	Description string `json:"description"`
}