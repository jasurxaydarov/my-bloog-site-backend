package storage

import (
	"jasurxaydarov/my-bloog-site-backend/storage/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type storage struct {
	contentRepo postgres.ContentRepoI 
	ownerRepo   postgres.OwnerRepoI
	commonRepo  postgres.CommonRepoI
}

type StorageI interface {
	GetContentRepo() postgres.ContentRepoI
	GetOwnerRepo() postgres.OwnerRepoI
	GetCommonRepo() postgres.CommonRepoI
}

func NewStorage(db *pgx.Conn, log logger.LoggerI) StorageI {

	return &storage{
		contentRepo: postgres.NewContent(db, log),
	}
}

func (s *storage) GetContentRepo() postgres.ContentRepoI {

	return s.contentRepo
}
func (s *storage) GetOwnerRepo() postgres.OwnerRepoI {

	return s.ownerRepo
}
func (s *storage) GetCommonRepo() postgres.CommonRepoI {

	return s.commonRepo  
}
