package controller

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(services *service.Service) *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")
	{
		newAuthRoutes(auth, services.Authorization)
	}

	// authMiddleware := &AuthMiddleware{services.Authorization}

	api := router.Group("/api")
	{
		newUserBannerRoutes(api.Group("/user_banner"), services.User)
		newBannerRoutes(api.Group("/banner"), services.Admin)
	}
	
	return router
}