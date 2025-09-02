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

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (p *PostController) List(c *gin.Context) {
	var posts []model.Post
	if err := database.DB.Preload("User").Find(&posts).Error; err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "internal error", err.Error())
		return
	}

	responses.JSONOk(c, posts)
}

func (p *PostController) Get(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := database.DB.Preload("User").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JOSNError(c, http.StatusNotFound, "not found", "post not found")
			return
		}
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	responses.JSONOk(c, post)
}

func (p *PostController) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	uid := middleware.MustGetUserID(c)
	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserId:  uid,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	database.DB.Preload("User").First(&post, post.ID)
	responses.JSONOk(c, post)
}

func (p *PostController) Update(c *gin.Context) {
	id := c.Param("id")
	uid := middleware.MustGetUserID(c)
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JOSNError(c, http.StatusNotFound, "not found", "post not found")
			return
		}
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	if post.UserId != uid {
		responses.JOSNError(c, http.StatusForbidden, "forbidden", "only author can modify")
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	if err := database.DB.Save(&post).Error; err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	database.DB.Preload("User").First(&post, post.ID)
	responses.JSONOk(c, post)
}

func (p *PostController) Delete(c *gin.Context) {
	id := c.Param("id")
	uid := middleware.MustGetUserID(c)
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JOSNError(c, http.StatusNotFound, "not found", "post not found")
			return
		}

		responses.JOSNError(c, http.StatusInternalServerError, "internal error", err.Error())
		return
	}

	if post.UserId != uid {
		responses.JOSNError(c, http.StatusForbidden, "forbidden", "only author can delete")
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
