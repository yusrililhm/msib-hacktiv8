package handler

import (
	"h8-assignment-2/infra/database"
	"h8-assignment-2/repository/item_repository/item_pg"
	"h8-assignment-2/repository/order_repository/order_pg"
	"h8-assignment-2/service"
	
	_ "h8-assignment-2/docs"

	"github.com/gin-gonic/gin"

	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order API Specs
// @version 1.0
// @description Tugas ke 2 Kampus Merdeka
//
// @contact.name Yusril Ilham Kholid
// @contact.email yusrililham@gmail.com
// @contact.url http://github.com/yusrililhm
//
// @host localhost:8080
// @BasePath /
func StartApplication() {
	app := gin.Default()

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	// init database
	database.InitiliazeDatabase()

	// dependencies injection
	db := database.GetDatabaseInstance()

	itemRepository := item_pg.NewItemRepositoryImpl(db)
	orderRepository := order_pg.NewOrderRepositoryImpl(db)
	orderService := service.NewOrderServiceImpl(orderRepository, itemRepository)
	orderHandler := NewOrderHandler(orderService)

	// grouping route
	orders := app.Group("/orders")

	{
		orders.POST("/", orderHandler.CreateOrder)
		orders.GET("", orderHandler.FetchOrder)
		orders.PUT("/:orderId", orderHandler.UpdateOrder)
		orders.DELETE("/:orderId", orderHandler.DeleteOrder)
	}

	// run server
	app.Run(":8080")
}
