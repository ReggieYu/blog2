package controllers

import (
	"blog/database"
	"blog/dto"
	"blog/middleware"
	"blog/model"
	"blog/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (cc *CommentController) ListByPost(c *gin.Context) {
	postId := c.Param("id")
	// check post exists or not
	var post model.Post
	if err := database.DB.First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JOSNError(c, http.StatusNotFound, "not found", "post not found")
			return
		}

		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	var comments []model.Comment
	if err := database.DB.Where("post_id = ?", postId).Preload("User").Order("id ASC").Find(&comments).Error; err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	responses.JSONOk(c, comments)
}

func (cc *CommentController) Create(c *gin.Context) {
	postId := c.Param("id")
	var post model.Post
	if err := database.DB.First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JOSNError(c, http.StatusNotFound, "not found", "post not found")
			return
		}

		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	uid := middleware.MustGetUserID(c)
	comment := model.Comment{
		Content: req.Content,
		UserId:  uid,
		PostId:  post.ID,
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "create failed", err.Error())
		return
	}

	database.DB.Preload("User").First(&comment, comment.ID)
	c.JSON(http.StatusCreated, comment)
}
