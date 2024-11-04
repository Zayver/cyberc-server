package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zayver/cyberc-server/dto/request"
	"github.com/zayver/cyberc-server/service"
)

type LoginController struct{
	loginService service.LoginService
}

func NewLoginController(loginS service.LoginService) LoginController{
	return LoginController{
		loginService: loginS,
	}
}

func (c *LoginController) Login(ctx *gin.Context){
	var loginRequest request.LoginRequest
	if err:= ctx.ShouldBindJSON(&loginRequest); err!=nil{
		ctx.Status(http.StatusBadRequest)
		return
	}
	response, err := c.loginService.Login(loginRequest) 
	if err != nil{
		ctx.Status(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *LoginController) Signup(ctx *gin.Context){
	c.loginService.Signup()
}