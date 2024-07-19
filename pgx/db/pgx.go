package db

import (
	"context"
	"fmt"
	"jasurxaydarov/my-bloog-site-backend/config"
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)


func ConnToDb(pgCfg config.PgConfig , log logger.LoggerI)(*pgx.Conn,error){

	dbUrl:=fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)
	conn,err:=pgx.Connect(context.Background(),dbUrl)

	if err!= nil{
		log.Error("error on Conn database",logger.Error(err))
		return nil,err
	}

	log.Debug("sucesfully conected with postgres")

	return conn,nil
}
