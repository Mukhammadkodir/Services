package handlers

import (
	"github/Services/workers/api/models"
	_ "github/Services/workers/api/models"
	pb "github/Services/workers/genproto/user_service"
	l "github/Services/workers/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser godoc
// @Summary Create new user
// @Description This API for creating a new user
// @Tags User
// @Accept json
// @Param body body models.CreateUser true "body"
// @Produce json
// @Success 201 {object} models.User
// @Router /users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	response, err := h.inMemoryStorage.Create(&body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser godoc
// @Summary GetUser
// @Schemes
// @Description  Get User
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param body body models.GetUser true "body"
// @Produce json
// @Success 200 {object} models.Get
// @Router /user/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	var (
		body        pb.LogReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	res, _ := h.inMemoryStorage.CheckField(&pb.PasswordReq{Password: body.Password})

	if res.Position != "Admin" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: "You have not permission to get user information",
			},
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	response, err := h.inMemoryStorage.Get(&pb.ById{Userid: body.Id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// Login godoc
// @Summary Login
// @Schemes
// @Description  Get My Profile
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param code path string true "Password"
// @Produce json
// @Success 200 {object} models.User
// @Router /user [post]
func (h *handlerV1) Login(c *gin.Context) {
	code := c.Param("code")

	response, err := h.inMemoryStorage.Login(&pb.PasswordReq{Password: code})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User Profile", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update User
// @Schemes
// @Description  Update User
// @Security 	BearerAuth
// @Tags User
// @Accept json
// @Param body body models.UpReq true "body"
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.UpReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	bul, err := h.inMemoryStorage.Update(&body)

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, bul)
}

// @Summary DeleteUser
// @Schemes
// @Description  Delete User
// @Security 	BearerAuth
// @Tags User
// @Param password path string true "Password"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("password")

	response, err := h.inMemoryStorage.Delete(&pb.PasswordReq{Password: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
