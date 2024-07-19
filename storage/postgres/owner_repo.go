package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ownerRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewOwneRrepo(db  *pgx.Conn,log logger.LoggerI)OwnerRepoI{

	return &ownerRepo{db: db,log: log}
}

func (o *ownerRepo)Login(ctx context.Context,login *modles.LoginOwn)(*modles.Owner,error){

	return nil,nil
}