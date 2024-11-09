package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zayver/cyberc-server/controller"
	"github.com/zayver/cyberc-server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Routes struct {
	loginController     controller.LoginController
	complaintController controller.ComplaintController
	jwtMiddleware       middleware.JwtMiddleware
	corsMiddleware      middleware.CorsMiddleware
}

func NewRoutes(loginC controller.LoginController, complaintC controller.ComplaintController, jwt middleware.JwtMiddleware, cors middleware.CorsMiddleware) Routes {
	return Routes{
		loginController: loginC,
		jwtMiddleware:   jwt,
		corsMiddleware:  cors,
		complaintController: complaintC,
	}
}

func (r *Routes) Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	api.Use(r.corsMiddleware.Setup())
	{
		complaint := api.Group("/complaint")
		complaint.GET("", r.jwtMiddleware.Handler(), r.complaintController.GetAllComplaints)
		complaint.GET("/filter", r.jwtMiddleware.Handler(), r.complaintController.GetComplaintsByCC)
		complaint.GET("/:id", r.complaintController.GetComplaintById)
		complaint.PUT("/:id/progress", r.complaintController.ProgressStatus)
		complaint.POST("", r.complaintController.CreateComplaint)

	}
	{
		login := api.Group("/login")
		login.POST("", r.loginController.Login)
	}

	return router
}
