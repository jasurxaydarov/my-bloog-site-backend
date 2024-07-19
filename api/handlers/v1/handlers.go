package v1

import (
	"jasurxaydarov/my-bloog-site-backend/storage"
	"jasurxaydarov/my-bloog-site-backend/storage/redis"

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
