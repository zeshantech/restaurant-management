package table

import (
	"github.co/golang-programming/restaurant/api/internal/middleware"
	"github.co/golang-programming/restaurant/api/internal/modules/table/controller"
	"github.co/golang-programming/restaurant/api/internal/modules/table/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/table")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateTableInput{}), controller.CreateTable)
	groupRouter.GET("/:id", controller.GetTableByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateTableInput{}), controller.UpdateTable)
	groupRouter.DELETE("/:id", controller.DeleteTable)
	groupRouter.GET("/", controller.ListTables)
}
