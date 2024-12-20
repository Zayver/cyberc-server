package service

import (
	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	"github.com/zayver/cyberc-server/dto/request"
	"github.com/zayver/cyberc-server/dto/response"
	"github.com/zayver/cyberc-server/model"
	"github.com/zayver/cyberc-server/repository"
)

type ComplaintService struct{
	complaintRepository repository.ComplaintRepository
}

func NewComplaintService(complaintRepo repository.ComplaintRepository) ComplaintService{
	return ComplaintService{
		complaintRepository: complaintRepo,
	}
}

func (c *ComplaintService) CreateComplaint(request request.CreateComplaintRequest) (response.ComplaintResponse, error){
	entity := model.Complaint{}
	mapper.AutoMapper(&request, &entity)
	entity.Status = model.CREATED
	entity, err := c.complaintRepository.CreateComplaint(entity)
	if err != nil{
		return response.ComplaintResponse{},err
	}
	response := response.ComplaintResponse{}
	mapper.AutoMapper(&entity, &response)
	return response, nil
}

func (c * ComplaintService) GetAllComplaints(page, pageSize int) ([]response.ComplaintResponse, int64 , error){
	entities, total, err := c.complaintRepository.GetAllComplaints(page, pageSize)
	if err != nil{
		return []response.ComplaintResponse{}, 0, err
	}
	var res = []response.ComplaintResponse{}
	mapper.MapperSlice(&entities, &res)
	return res, total, nil
}

func (c * ComplaintService) GetComplaintById(id uuid.UUID) (model.Complaint, error){
	entity, err := c.complaintRepository.GetComplaintById(id)
	if err != nil{
		return model.Complaint{}, err
	}
	return entity, nil
}

func (c * ComplaintService) DeleteComplaint(id uuid.UUID) error{
	if err := c.complaintRepository.DeleteComplaint(id); err != nil{
		return err
	}
	return nil
}

func (c * ComplaintService) GetComplaintsByCC(cc string) ([]response.ComplaintResponse, error){
	complaints, err := c.complaintRepository.GetComplaintsByCC(cc)
	if err != nil{
		return []response.ComplaintResponse{}, nil
	}
	var res = []response.ComplaintResponse{}
	mapper.MapperSlice(&complaints, &res)
	return res, nil
}

func (c * ComplaintService) ProgressStatus(id uuid.UUID) error{
	if err := c.complaintRepository.ProgressStatus(id); err != nil{
		return err
	}
	return nil
}