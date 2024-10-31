package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zayver/cybercomplaint-server/controller"
	"github.com/zayver/cybercomplaint-server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Routes struct{
	loginController controller.LoginController
	jwtMiddleware middleware.JwtMiddleware
	corsMiddleware middleware.CorsMiddleware
}

func NewRoutes(loginC controller.LoginController, jwt middleware.JwtMiddleware, cors middleware.CorsMiddleware) Routes{
	return Routes{
		loginController: loginC,
		jwtMiddleware: jwt,
		corsMiddleware: cors,
	}
}

func (r *Routes) Init() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	api.Use(r.corsMiddleware.Setup())
	{
		//complaintController := controller.ComplaintController{}
		//complaint := api.Group("/complaint")
		//complaint.GET("", complaintController.GetAllComplaints)
		//complaint.GET("/:id", complaintController.GetById)
		//complaint.POST("", complaintController.CreateComplaint)
	}
	{
		login := api.Group("/login")
		login.POST("", r.loginController.Login)
		login.POST("/signup", r.loginController.Signup)
	}
	{
		test := api.Group("/test")
		test.GET("/protected", r.jwtMiddleware.Handler() ,func(ctx *gin.Context) {
			ctx.JSON(200, "ACCEPTED PROTECTED")
		})

		test.GET("/unprotected", func(ctx *gin.Context) {
			ctx.JSON(200, "ACCEPTED PROTECTED")
		})
	}

	return router
}