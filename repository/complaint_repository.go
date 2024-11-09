package repository

import (
	"errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/zayver/cyberc-server/config"
	"github.com/zayver/cyberc-server/model"
	"github.com/zayver/cyberc-server/scopes"
)

var ErrUnprocessableEntity = errors.New("unprocessable entity")

type ComplaintRepository struct{
	db config.DB
}
func NewComplaintRepository(db config.DB) ComplaintRepository{
	return ComplaintRepository{
		db: db,
	}
}

func (c *ComplaintRepository) CreateComplaint(entity model.Complaint) (model.Complaint, error){
	if err := c.db.DB.Save(&entity).Error; err != nil{
		log.Error("Error saving entity Complaint: ", err)
		return model.Complaint{}, err
	}
	return entity, nil
}

func (c *ComplaintRepository) GetAllComplaints(page, pageSize int) ([]model.Complaint, int64, error){
	var complaints []model.Complaint
	if err := c.db.DB.Scopes(scopes.Pagination(page, pageSize)).Find(&complaints).Error; err!=nil{
		log.Error("Error retriving all complaints: ", err)
		return []model.Complaint{},0, err
	}
	var total int64
	if err := c.db.DB.Model(&model.Complaint{}).Count(&total).Error; err != nil{
		log.Error("Error counting all complaints: ", err)
		return []model.Complaint{}, 0, err
	} 
	return complaints, total, nil
}

func (c *ComplaintRepository) GetComplaintById(id uuid.UUID) (model.Complaint, error){
	var complaint model.Complaint
	if err := c.db.DB.First(&complaint, "id = ? ", id).Error; err!=nil{
		log.Error("Error retriving all complaints: ", err)
		return model.Complaint{}, err
	}
	return complaint, nil
}

func (c *ComplaintRepository) DeleteComplaint(id uuid.UUID) error {
	complaint, err := c.GetComplaintById(id)
	if err != nil{
		return err
	}
	if err := c.db.DB.Delete(&complaint).Error; err != nil{
		log.Error("Error deleting entity complaint: ", err)
		return err
	}
	return nil
}

func (c *ComplaintRepository) GetComplaintsByCC(cc string) ([]model.Complaint, error){
	var complaints []model.Complaint
	if err := c.db.DB.Limit(10).Where("cc LIKE ?", cc +"%").Find(&complaints).Error; err != nil{
		log.Error("Error getting complaints by cc: ", err)
	}
	return complaints, nil
}

func(c *ComplaintRepository) ProgressStatus(id uuid.UUID) error{
	entity, err := c.GetComplaintById(id)
	if err != nil{
		return err
	}
	if entity.Status == model.FINALIZED {
		return ErrUnprocessableEntity
	}
	entity.Status += 1
	if err := c.db.DB.Save(&entity).Error; err != nil{
		return err
	}
	return nil
} 