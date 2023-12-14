package svc

import (
	"context"
	"fmt"
	"shorturl/dal/query"
	"shorturl/internal/config"

	"github.com/bits-and-blooms/bloom/v3"
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
	g           *singleflight.Group
	db          *gorm.DB
	rdb         *redis.Client
	bloomFilter *bloom.BloomFilter
)

type ServiceContext struct {
	Config config.Config
	G      *singleflight.Group
	DB     *gorm.DB
	RDB    *redis.Client
	Bloom  *bloom.BloomFilter
}

func init() {
	g = &singleflight.Group{}
	db = connectDB(dsn)
	rdb = connectRedis(redisAddr)
	query.SetDefault(db)
	bloomFilter = bloom.NewWithEstimates(100000, 0.01)
	// 初始化布隆过滤器
	su, err := query.ShortURLMap.WithContext(context.Background()).Find()
	if err != nil {
		panic(fmt.Errorf("query.ShortURLMap.WithContext(context.Background()).Find() error: %v", err))
	}
	for _, v := range su {
		bloomFilter.Add([]byte(v.Surl))
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		G:      g,
		DB:     db,
		RDB:    rdb,
		Bloom:  bloomFilter,
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
