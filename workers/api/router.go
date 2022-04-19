package api

import (
	v1 "github/Services/workers/api/handler"
	"github/Services/workers/config"
	"github/Services/workers/pkg/logger"
	"github/Services/workers/services"
	"github/Services/workers/storage/repo"

	"github/Services/workers/api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 			Task api
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
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	InMemoryStorage repo.UserStorageI
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:          option.Logger,
		ServiceManager:  option.ServiceManager,
		Cfg:             option.Conf,
		InMemoryStorage: option.InMemoryStorage,
	})

	api := router.Group("/v1")


	api.POST("/users", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.Get)
	api.PUT("/user/:id", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.POST("/user", handlerV1.Login)

	return router
}
