package main

import (
	"context"
	"fmt"
	"jasurxaydarov/my-bloog-site-backend/api"
	"jasurxaydarov/my-bloog-site-backend/config"
	"jasurxaydarov/my-bloog-site-backend/pgx/db"
	"jasurxaydarov/my-bloog-site-backend/storage"
	"jasurxaydarov/my-bloog-site-backend/storage/redis"

	logg "github.com/saidamir98/udevs_pkg/logger"
)

func main() {

	cfg := config.Load()

	log := logg.NewLogger(cfg.GeneralConfig.AppName, config.DebugMode)

	pgx, err := db.ConnToDb(cfg.PgConfig, log)

	if err != nil {

		return
	}

	fmt.Println(pgx)

	redisCli, err := db.ConnRedis(log, context.Background(), cfg.RedisConfig)

	if err != nil {

		return
	}

	fmt.Println(redisCli)
	cache := redis.NewRedisRepo(redisCli, log)

	//storage:=storage.NewStorage(&pgxpool.Pool{},log)

	storage := storage.NewStorage(pgx, log)

	engine := api.Api(api.Options{Storage: storage, Log: log, Cache: cache})

	engine.Run()

	log.Debug("gin run ")
}
