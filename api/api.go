package api

import (
	v1 "jasurxaydarov/my-bloog-site-backend/api/handlers/v1"
	"jasurxaydarov/my-bloog-site-backend/storage"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Options struct {
	Storage storage.StorageI
	Log     logger.LoggerI
}

func Api(O Options) *gin.Engine {

	h := v1.NewHandlers(v1.Handlers(O))

	engine := gin.Default()

	api := engine.Group("/api")

	own:=api.Group("/own")

	{
		own.POST("/category", h.CreateCategory)
		own.GET("/category/:id",h.GetCategory)
		own.GET("/category",h.GetCategories)
		//api.PUT("/cattegory/:id")
		//api.DELETE("/cattegory/:id")
	}

	return engine

}
