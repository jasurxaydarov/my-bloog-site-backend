package postgres

import (
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ContentRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewContent(db *pgx.Conn, log logger.LoggerI) ContentRepoI {

	return &ContentRepo{db, log}
}
