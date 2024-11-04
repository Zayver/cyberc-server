package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zayver/cyberc-server/dto/request"
	"github.com/zayver/cyberc-server/scopes"
	"github.com/zayver/cyberc-server/service"
	"gorm.io/gorm"
)

type ComplaintController struct{
	complaintService service.ComplaintService
}

func NewComplaintController(complaintS service.ComplaintService) ComplaintController{
	return ComplaintController{
		complaintService: complaintS,
	}
}

func (c * ComplaintController) GetAllComplaints(ctx * gin.Context){
	pageQuery := ctx.DefaultQuery("page", fmt.Sprint(scopes.DEFAULT_PAGE))
	pageSizeQuery := ctx.DefaultQuery("pageSize", string(scopes.DEFAULT_PAGE_SIZE))
	page, err := strconv.Atoi(pageQuery)
	if err != nil{
		page = scopes.DEFAULT_PAGE
	}
	pageSize, err := strconv.Atoi(pageSizeQuery)
	if err != nil{
		pageSize = scopes.DEFAULT_PAGE_SIZE
	}	
	complaints, err := c.complaintService.GetAllComplaints(page, pageSize)
	if err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"complaints": complaints,
		"hasNext": len(complaints) == pageSize,
	})
}

func (c * ComplaintController) GetComplaintById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}
	entity, err := c.complaintService.GetComplaintById(id)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, entity)
}

func (c *ComplaintController) CreateComplaint(ctx * gin.Context){
	var request request.CreateComplaintRequest
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}
	_, err := c.complaintService.CreateComplaint(request)

	if err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
