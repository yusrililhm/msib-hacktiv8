package router

import (
	"go-jwt/controller"
	"go-jwt/database"
	_ "go-jwt/docs"
	"go-jwt/entity"
	"go-jwt/helper"
	"go-jwt/middleware"
	"go-jwt/repository"
	"go-jwt/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT = helper.GetEnv("PORT")

//	@title	Authentication With JWT
//	@version	1.0
//	@description	Auth with JWT
//
//	@contact.name	Yusril Ilham
//	@contact.email	yusrililham62@gmail.com
//	@contact.url	https://github.com/yusrililhm
//
//	@host	0.0.0.0:3000
//	@BasePath	/

func router(userController controller.UserController, productController controller.ProductController) *gin.Engine {
	r := gin.Default()

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	user := r.Group("user")

	{
		user.POST("/register", userController.UserRegister)
		user.POST("/login", userController.UserLogin)
	}

	product := r.Group("product")

	{
		product.Use(middleware.Auth())
		product.POST("/add", productController.CreateProduct)
		product.PATCH("/:productId", middleware.ProductAuth(), productController.UpdateProduct)
	}

	return r
}

func StartServer() {
	db, err := database.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	db.AutoMigrate(entity.User{}, entity.Product{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := router(userController, productController)
	router.Run("0.0.0.0:" + PORT)
}
