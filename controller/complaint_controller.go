package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zayver/cyberc-server/dto/request"
	"github.com/zayver/cyberc-server/repository"
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
	complaints, total, err := c.complaintService.GetAllComplaints(page, pageSize)
	if err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"complaints": complaints,
		"total": total,
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
	res, err := c.complaintService.CreateComplaint(request)

	if err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": res.ID,
	})
}

func (c * ComplaintController) GetComplaintsByCC(ctx *gin.Context){
	cc := ctx.Query("cc")
	if cc == ""{
		ctx.Status(http.StatusBadRequest)
		return
	}
	complaints, err := c.complaintService.GetComplaintsByCC(cc)
	if err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"complaints": complaints,
		"hasNext": len(complaints) == 10,
	})
}

func (c *ComplaintController) ProgressStatus(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err:= c.complaintService.ProgressStatus(id); err != nil{
		if errors.Is(err, repository.ErrUnprocessableEntity){
			ctx.Status(http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound){
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return 
	}
}
