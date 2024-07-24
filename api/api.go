package api

import (
	ewaggerfiles "github.com/swaggo/files"
	ginSwagger	"github.com/swaggo/gin-swagger"
	_ "github.com/jasurxaydarov/my-bloog-site-backend/api/docs"
	v1 "jasurxaydarov/my-bloog-site-backend/api/handlers/v1"
	"jasurxaydarov/my-bloog-site-backend/api/middlewares"
	"jasurxaydarov/my-bloog-site-backend/storage"
	"jasurxaydarov/my-bloog-site-backend/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Options struct {
	Storage storage.StorageI
	Log     logger.LoggerI
	Cache   redis.RedisRepoI
}

func Api(O Options) *gin.Engine {

	h := v1.NewHandlers(v1.Handlers(O))

	engine := gin.Default()

	api := engine.Group("/api")

	own := api.Group("/own")

	{

		own.POST("/sign-out", h.OwnSignOut)

		own.POST("/category", h.CreateCategory) // completed
		own.PUT("/cattegory/:id", h.UpdateCategory)
		own.DELETE("/cattegory/:id", h.DeleteCategory)

		own.POST("/sub-category", h.CreateSubCategory) //
		own.PUT("/sub-cattegory/:id", h.UpdateSubCategory)
		own.DELETE("/sub-cattegory/:id", h.DeleteSubCategory)

		own.POST("/article", h.CreateArticle) //
		own.PUT("/article/:id", h.UpdateSubCategory)
		own.DELETE("/article/:id", h.DeleteSubCategory)

	}

	vw := api.Group("/vw")
	vw.Use(middlewares.AuthMiddleware())
	{
		vw.POST("log-out", h.SignOut)

		vw.POST("/comment", h.AddComment) //completed
		vw.PUT("/comment", h.UpdateComment)
		vw.DELETE("/comment", h.DeleteComment) //

	}

	pb := api.Group("/pb")

	{
		pb.POST("/own/sign-in", h.OwnSignIn) // in-prosses
		pb.POST("/check-user", h.CheckUser)  //completed
		pb.POST("/check-otp", h.CheckOtp)    //completed
		pb.POST("/sign-up", h.SignUp)        //completed
		pb.POST("/sign-in", h.SignIn)        //

		pb.GET("/category/:id", h.GetCategory) //completed
		pb.GET("/category", h.GetCategories)   //completed

		pb.GET("/sub-category/:id", h.GetSubCategory) //completed
		pb.GET("/sub-category", h.GetSubCategories)   //completed

		pb.GET("/article/:id", h.GetArticle) //completed
		pb.GET("/article", h.GetArticles)    //completed

		vw.GET("/comment/:id", h.GetComment) //completed
		vw.GET("/comment", h.GetComments)    //completed

	}

	return engine

}
