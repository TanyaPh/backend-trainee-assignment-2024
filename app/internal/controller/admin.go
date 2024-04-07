package controller

import (
	"api/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)


type AdminBannerRoutes struct {
	adminService service.Admin
}

func newBannerRoutes(g *gin.RouterGroup, adminService service.Admin) *AdminBannerRoutes {
	h := &AdminBannerRoutes{
		adminService: adminService,
	}

	g.POST("/", h.createBanner)
	g.GET("/", h.getBanner)
	g.PATCH("/{id}", h.updateBanner)
	g.DELETE("/{id}", h.deleteBanner)

	return h
}

func (h *AdminBannerRoutes) createBanner(c *gin.Context) {
	fmt.Println("Admin create banner")
}

func (h *AdminBannerRoutes) getBanner(c *gin.Context) {
	fmt.Println("Admin see banner")
}

func (h *AdminBannerRoutes) updateBanner(c *gin.Context) {
	fmt.Println("Admin update banner")
}

func (h *AdminBannerRoutes) deleteBanner(c *gin.Context) {
	fmt.Println("Admin delete banner")
}