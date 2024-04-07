package controller

import (
	"api/internal/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserBannerRoutes struct {
	userService service.User
}

func newUserBannerRoutes(g *gin.RouterGroup, userService service.User) *UserBannerRoutes {
	h := &UserBannerRoutes{
		userService: userService,
	}

	g.GET("/", h.getUserBanner)

	return h
}

func (h *UserBannerRoutes) getUserBanner(c *gin.Context) {
	var lastRevision bool

	tagId, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		// newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	featureId, err := strconv.Atoi(c.Param("feature_id"))
	if err != nil {
		// newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	h.userService.getBanner(tagId, featureId, lastRevision)

	fmt.Println("User see banner")
}