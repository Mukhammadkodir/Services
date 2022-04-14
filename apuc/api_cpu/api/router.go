package api

import (
	v1 "github/Services/apuc/api_cpu/api/handler"
	"github/Services/apuc/api_cpu/config"
	"github/Services/apuc/api_cpu/pkg/logger"
	"github/Services/apuc/api_cpu/services"

	"github/Services/apuc/api_cpu/api/docs"

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
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/posts", handlerV1.CreatePost)
	api.GET("/post/:id", handlerV1.GetPost)
	api.PUT("/post/:id", handlerV1.UpdatePost)
	api.DELETE("/post/:id", handlerV1.DeletePost)

	api.POST("/users", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.Get)
	api.PUT("/user/:id", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
