package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/saidamir98/udevs_pkg/logger"
)

type RedisRepoI interface {
	Exist(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key, value string, exp int) error
	Get(ctx context.Context, key string) error
	Del(ctx context.Context, key string) (any, error)
	GetDel(ctx context.Context, key string) (any, error)
}
type redisRepo struct {
	cli *redis.Client
	log logger.LoggerI
}

func NewRedisRepo(cli *redis.Client, log logger.LoggerI) RedisRepoI {

	return &redisRepo{cli, log}
}

func (r *redisRepo) Exist(ctx context.Context, key string) (bool, error) {

	defer r.cli.Close()

	isExists, err := r.cli.Do(ctx, "EXISTS", key).Result()

	if err != nil {
		r.log.Error("error on check exists", logger.Error(err))
		return false, err
	}

	value, _ := isExists.(int)

	return value == 1, nil
}


func (r *redisRepo) Set(ctx context.Context, key, value string, exp int) error {
	r.log.Debug("req in Set")
	_, err := r.cli.SetEX(ctx, key, value, time.Second*time.Duration(exp)).Result()

	if err != nil {
		r.log.Error("erro on setting to cache ", logger.Error(err))
		return err
	}

	return nil
}


func (r *redisRepo) Get(ctx context.Context, key string) error {

	return nil
}
func (r *redisRepo) Del(ctx context.Context, key string) (any, error) {

	return nil, nil
}
func (r *redisRepo) GetDel(ctx context.Context, key string) (any, error) {

	return nil, nil
}
