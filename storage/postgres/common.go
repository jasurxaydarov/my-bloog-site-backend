package postgres

import (
	"context"
	"fmt"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type CommonRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}
func NewCoomnRepo(db  *pgx.Conn,log logger.LoggerI)CommonRepoI{

	return &CommonRepo{db,log}
}
func (c *CommonRepo)CheckExists(ctx context.Context,req *modles.Common)(bool,error){
	var isExists bool
	query:=fmt.Sprintf("SELECT EXISTS (SELECT 1 FROME %s WHERE %s ='%s')",req.TableName,req.ColumnName,req.ExpValue)

	err:=c.db.QueryRow(ctx,query).Scan(&isExists)

	if err!= nil{
		c.log.Error("error on CheckExists",logger.Error(err))
		return false,nil
	}
	return isExists,nil
}