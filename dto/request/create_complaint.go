package request
type CreateComplaintRequest struct{
	Name string `json:"name" binding:"required"`
	Surname string `json:"surName" binding:"required"`
}