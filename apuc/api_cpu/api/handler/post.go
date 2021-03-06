package handlers

import (
	"context"
	_ "github/Services/apuc/api_cpu/api/models"
	pb "github/Services/apuc/api_cpu/genproto/post_service"
	l "github/Services/apuc/api_cpu/pkg/logger"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost godoc
// @Summary Create new post
// @Description This API for creating a new post
// @Tags Post
// @Accept json
// @Param body body models.CreatePost true "body"
// @Produce json
// @Success 201 {object} models.Post
// @Router /posts [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        pb.Post
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetPost godoc
// @Summary Get post
// @Description  Get post
// @Tags Post
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} models.Post
// @Router /post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Get(ctx, &pb.ById{Userid: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update post
// @Schemes
// @Description  Update post
// @Tags Post
// @Accept json
// @Param body body models.UpdatePost true "body"
// @Produce json
// @Success 200 {object} models.Post
// @Router /post/{id} [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pb.Post
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePost godoc
// @Summary Delete post
// @Schemes
// @Description  Delete post
// @Tags Post
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200
// @Router /post/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Delete(ctx, &pb.ById{Userid: guid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
