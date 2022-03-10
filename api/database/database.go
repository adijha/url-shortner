package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}

// package database

// import (
// 	"context"
// 	"errors"
// 	"os"

// 	"github.com/go-redis/redis/v8"
// )

// type Database struct {
// 	Client *redis.Client
// }

// var (
// 	ErrNil = errors.New("no matching record found in redis database")
// 	Ctx    = context.TODO()
// )

// func CreateClient(dbNo int) (*Database, error) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("DB_ADDR"),
// 		Password: os.Getenv("DB_PASS"),
// 		DB:       dbNo,
// 	})
// 	if err := client.Ping(Ctx).Err(); err != nil {
// 		return nil, err
// 	}
// 	return &Database{
// 		Client: client,
// 	}, nil
// 	// return &Database{
// 	// 	Client: client,
// 	// }, nil
// }
