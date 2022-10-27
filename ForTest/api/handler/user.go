package handlers

import (
	"fmt"
	_ "github/Services/ForTest/api/models"
	pb "github/Services/ForTest/genproto/user"
	l "github/Services/ForTest/pkg/logger"
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

	response, err := h.storageManager.User().Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to create user", l.Error(err))
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
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	// var jspbMarshal protojson.MarshalOptions
	// jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	fmt.Println("handler: ",guid)
	
	response, err := h.storageManager.User().Get(&pb.ById{Userid: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get User", l.Error(err))
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
// @Param body body models.UpdateUser true "body"
// @Produce json
// @Success 200 {object} models.User
// @Router /user [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json of body", l.Error(err))
		return
	}

	fmt.Println(body.Id)
	response, err := h.storageManager.User().Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteUser
// @Schemes
// @Description  Delete User
// @Security 	BearerAuth
// @Tags User
// @Param id path string true "ID"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	
	response, err := h.storageManager.User().Delete(&pb.ById{Userid: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
