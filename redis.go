package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type URedis struct {
	Enable bool
	Client *redis.Client
}

func (uredis *URedis) Open(gconfig Config) {
	uredis.Enable = gconfig.Redis.Enable

	if !uredis.Enable {
		return
	}

	uredis.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", gconfig.Redis.DBHost, gconfig.Redis.DBPort),
		Password: gconfig.Redis.DBPass,
		DB:       gconfig.Redis.DBName,
	})

	if err := uredis.Client.Ping(ctx).Err(); err != nil {
		log.Fatalf("%s.\n", err)
	}
}

func (uredis *URedis) Set(short, long string) error {
	if err := uredis.Client.Set(ctx, short, long, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (uredis *URedis) Get(short string) (string, error) {
	var val string
	val, err := uredis.Client.Get(ctx, short).Result()

	if err == redis.Nil || err != nil {
		// 空 || 异常
		return val, err
	} else {
		return val, nil
	}
}

func (uredis *URedis) Del(short string) error {
	return uredis.Client.Del(ctx, short).Err()
}
