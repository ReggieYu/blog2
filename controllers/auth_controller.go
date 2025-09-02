package controllers

import (
	"blog/config"
	"blog/database"
	"blog/dto"
	"blog/middleware"
	"blog/model"
	"blog/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Cfg *config.Config
}

func NewAuthController(cfg *config.Config) *AuthController {
	return &AuthController{Cfg: cfg}
}

func (a *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", "failed to hash password")
		return
	}

	user := model.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	if err := database.DB.Create(&user).Error; err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "create user failed", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registered",
	})
}

func (a *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JOSNError(c, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		responses.JOSNError(c, http.StatusUnauthorized, "unauthorized", err.Error())
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		responses.JOSNError(c, http.StatusUnauthorized, "invaild credentials", "username or password is wrong")
		return
	}

	token, err := middleware.NewToken(a.Cfg, user.ID, user.Username)
	if err != nil {
		responses.JOSNError(c, http.StatusInternalServerError, "server error", "failed to create token")
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}

func (a *AuthController) Me(c *gin.Context) {
	uid := middleware.MustGetUserID(c)
	var user model.User
	if err := database.DB.First(&user, uid).Error; err != nil {
		responses.JOSNError(c, http.StatusNotFound, "not found", "user not found")
		return
	}

	user.Password = "" // no psssword return
	responses.JSONOk(c, user)
}
