package controller

import (
	"api/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type orderRoutes struct {
	orderService service.Order
}

func newOrderRoutes(g *gin.RouterGroup, orderService service.Order) *orderRoutes {
	r := &orderRoutes{
		orderService: orderService,
	}

	g.POST("/create", r.createOrder)
	// g.GET("/", r.getOrderByID)

	return r
}


func (o *orderRoutes) createOrder(c *gin.Context) {
	o.orderService.Create()
	fmt.Println("Order create")
}

// type getAllListsResponse struct {
// 	Data []restAPI.TodoList `json:"data"`
// }


// func (o *orderRoutes) getAllOrders(c *gin.Context) {
// 	fmt.Println("Order list show")
// }


// func (o *orderRoutes) getOrderByID(c *gin.Context) {
// 	fmt.Println("Order show")
// }

// func (o *orderRoutes) updateOrder(c *gin.Context) {
// 	fmt.Println("Order update")
// }


// func (o *orderRoutes) deleteOrder(c *gin.Context) {
// 	fmt.Println("Order delete")
// }
