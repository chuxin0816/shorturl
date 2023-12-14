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
	loadDataToBloom()
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

func loadDataToBloom() {
	cnt, err := query.ShortURLMap.WithContext(context.Background()).Count()
	if err != nil {
		panic(fmt.Errorf("query.ShortURLMap.WithContext(context.Background()).Count() error: %v", err))
	}
	
	count := int(cnt)
	pageTotal := 0
	pageSize := 50
	if count%pageSize == 0 {
		pageTotal = count / pageSize
	} else {
		pageTotal = count/pageSize + 1
	}

	for page := 1; page <= pageTotal; page++ {
		su, err := query.ShortURLMap.WithContext(context.Background()).Where(query.ShortURLMap.IsDel.Eq(0)).Offset((page - 1) * pageSize).Limit(int(pageSize)).Select(query.ShortURLMap.Surl).Find()
		if err != nil {
			panic(fmt.Errorf("query.ShortURLMap.WithContext(context.Background()).Find() error: %v", err))
		}
		
		for _, v := range su {
			bloomFilter.Add([]byte(v.Surl))
		}
	}
}
