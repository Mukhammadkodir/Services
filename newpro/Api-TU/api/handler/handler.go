package handlers

import (
	"errors"
	"net/http"
	"newpro/Api-TU/api/models"
	"newpro/Api-TU/api/token"
	"newpro/Api-TU/config"
	"newpro/Api-TU/pkg/logger"
	"newpro/Api-TU/services"
	"newpro/Api-TU/storage/repo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log             logger.Logger
	serviceManager  services.IServiceManager
	cfg             config.Config
	inMemoryStorage repo.InMemoryStorageI
	jwtHandler      token.JWTHandler
}

type HandlerV1Config struct {
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	Cfg             config.Config
	InMemoryStorage repo.InMemoryStorageI
	JwtHandler      token.JWTHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:             c.Logger,
		serviceManager:  c.ServiceManager,
		cfg:             c.Cfg,
		inMemoryStorage: c.InMemoryStorage,
	}
}

func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		Authorization   models.GetProfileByJwtRequestModel
		claims          jwt.MapClaims
		err             error
	)

	Authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Unauthorized request",
			},
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return nil
	}

	h.jwtHandler.Token = Authorization.Token
	claims, err = h.jwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Unauthorized request",
			},
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return nil
	}
	return claims
}
