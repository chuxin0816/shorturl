package svc

import (
	"fmt"
	"shorturl/dal/query"
	"shorturl/internal/config"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn       = "root:123456@tcp(127.0.0.1:3306)/shorturl?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	redisAddr = "127.0.0.1:6379"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	G   *singleflight.Group
)

type ServiceContext struct {
	Config config.Config
	G      *singleflight.Group
	DB     *gorm.DB
	RDB    *redis.Client
}

func init() {
	DB = connectDB(dsn)
	RDB = connectRedis(redisAddr)
	query.SetDefault(DB)
	G = &singleflight.Group{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		G:      G,
		DB:     DB,
		RDB:    RDB,
	}
}
func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func connectRedis(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	return rdb
}
