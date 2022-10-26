package handlers

import (
	"errors"
	"github/Services/ForTest/api/models"
	"github/Services/ForTest/config"
	"github/Services/ForTest/pkg/logger"
	"github/Services/ForTest/storage"
	"net/http"
	"github.com/jmoiron/sqlx"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	storageManager storage.IStorage
	cfg            config.Config
}

type HandlerV1Config struct {
	Logger         logger.Logger
	StorageManager storage.IStorage
	Cfg            config.Config
}

func New(c *HandlerV1Config, db *sqlx.DB) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		storageManager: storage.NewStoragePg(db),
		cfg:            c.Cfg,
	}
}

func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		claims          jwt.MapClaims
		err             error
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.Error{
				Message: "Unauthorized request",
			},
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return nil
	}
	return claims
}
