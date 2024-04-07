package controller

import (
	"api/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authService service.Auth
}

func newAuthRoutes(g *gin.RouterGroup, authService service.Auth) *AuthRoutes {
	r := &AuthRoutes{
		authService: authService,
	}

	g.POST("/sign-up", r.signUp)
	g.POST("/sign-in", r.signIn)

	return r
}

func (h *AuthRoutes) signUp(c *gin.Context) {
	fmt.Println("signIn")
}

func (h *AuthRoutes) signIn(c *gin.Context) {
	fmt.Println("signIn")
}