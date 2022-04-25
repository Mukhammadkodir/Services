package api

import (
	v1 "github/Services/workers/api/handler"
	"github/Services/workers/config"
	"github/Services/workers/storage/repo"

	"github/Services/workers/api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Register api
// @version			1.0
// @description		This is User and Task service Api
// @termsOfService	http://swagger.io/terms/

// @securityDefinitions.apikey BearerAuth
// @in  header
// @name Authorization

// @contact.name	Api Support
// @contact.url		http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/license/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath /v1

type Option struct {
	Conf            config.Config
	InMemoryStorage repo.UserStorageI
}

func New(option Option) *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/v1"
	router.Use(gin.Logger())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Cfg:             option.Conf,
		InMemoryStorage: option.InMemoryStorage,
	})


	router.POST("/users", handlerV1.CreateUser)
	router.GET("/user/:id", handlerV1.Get)
	router.PUT("/user/:id", handlerV1.UpdateUser)
	router.DELETE("/user/:id", handlerV1.DeleteUser)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/user", handlerV1.Login)

	return router
}
