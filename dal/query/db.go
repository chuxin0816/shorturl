package query

import (
	"fmt"

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
	db          *gorm.DB
	RDB         *redis.Client
	SFGroup *singleflight.Group
)

func init() {
	connectDB(dsn)
	connectRedis(redisAddr)
	SetDefault(db)
	SFGroup = &singleflight.Group{}
}

func connectDB(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
}

func connectRedis(addr string) {
	RDB = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
}
