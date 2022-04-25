package handlers

import (
	"github/Services/workers/storage/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser godoc
// @Summary Create new user
// @Description This API for creating a new user
// @Tags User
// @Accept json
// @Param body body models.User true "body"
// @Produce json
// @Success 201 {object} models.User
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        models.User
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.inMemoryStorage.Create(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	var (
		body        models.GetUser
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, _ := h.inMemoryStorage.CheckField(models.PasswordReq{Password: body.Password})

	if res.Position != "admin" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: "You have not permission to get user information",
			},
		})
		return
	}

	response, err := h.inMemoryStorage.Get(models.ById{Userid: body.Id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user [post]
func (h *handlerV1) Login(c *gin.Context) {
	code := c.Param("code")

	response, err := h.inMemoryStorage.Login(models.PasswordReq{Password: code})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        models.UpReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bul, err := h.inMemoryStorage.Update(body)

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
// @Success 200 {object} models.EmptyResp
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("password")

	response, err := h.inMemoryStorage.Delete(models.PasswordReq{Password: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
