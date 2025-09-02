package router

import (
	"blog/config"
	"blog/controllers"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// static file server
	r.Static("/static", "./static")

	// index router -- direct return html
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	v1 := r.Group("/api/v1")
	{
		authCtrl := controllers.NewAuthController(cfg)
		postCtrl := controllers.NewPostController()
		cmtCtrl := controllers.NewCommentController()

		// auth
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authCtrl.Register)
			auth.POST("/login", authCtrl.Login)
		}

		v1.GET("/posts", postCtrl.List)
		v1.GET("/posts/:id", postCtrl.Get)
		v1.GET("/posts/:id/comments", cmtCtrl.ListByPost)

		// protected routers
		pr := v1.Group("")
		pr.Use(middleware.AuthRequred(cfg))
		{
			pr.GET("/me", authCtrl.Me)
			pr.POST("/posts", postCtrl.Create)
			pr.PUT("/posts/:id", postCtrl.Update)
			pr.DELETE("/posts/:id", postCtrl.Delete)

			pr.POST("/posts/:id/comments", cmtCtrl.Create)
		}
	}

	return r
}
