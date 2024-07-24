package v1

import (
	"jasurxaydarov/my-bloog-site-backend/modles"
	"jasurxaydarov/my-bloog-site-backend/pgx/helpers"
	"jasurxaydarov/my-bloog-site-backend/storage"
	"jasurxaydarov/my-bloog-site-backend/storage/redis"
	"jasurxaydarov/my-bloog-site-backend/token"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

type handlers struct {
	storage storage.StorageI
	log     logger.LoggerI
	cache   redis.RedisRepoI
}

type Handlers struct {
	Storage storage.StorageI
	Log     logger.LoggerI
	Cache   redis.RedisRepoI
}

func NewHandlers(h Handlers) handlers {

	return handlers{h.Storage, h.Log, h.Cache}
}

func Auth(ctx *gin.Context) *modles.Clamis {

	
		tokenString := ctx.GetHeader("authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
			ctx.Abort()
		}

		Tokenclaim, err := token.ParseJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
		}

		claim:=&modles.Clamis{}
	
		helpers.DataParser(Tokenclaim,claim)
	return claim
}
