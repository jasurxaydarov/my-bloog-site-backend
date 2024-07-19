package v1

import (
	"jasurxaydarov/my-bloog-site-backend/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type handlers struct{
	storage 	storage.StorageI
	log 		logger.LoggerI

}

type Handlers struct{
	Storage 	storage.StorageI
	Log 		logger.LoggerI

}

func NewHandlers(h Handlers)handlers{

	return handlers{h.Storage,h.Log}
}

