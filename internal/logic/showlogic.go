package logic

import (
	"context"
	"errors"
	"time"

	"shorturl/dal/query"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"shorturl/pkg/randtime"
	"shorturl/pkg/singleflight"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	ErrShortUrlNotExist = errors.New("short url not exist")
	delayTime           = time.Millisecond * 100
)

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	su := query.ShortURLMap
	// 使用singleflight减轻缓存压力，并解决缓存击穿问题
	key := query.GetPrefix(req.ShortURL)
	lURL, err, _ := singleflight.G.Do(key, func() (interface{}, error) {
		go func() {
			time.Sleep(delayTime)
			singleflight.G.Forget(key)
		}()
		// 查询缓存
		lURL, err := query.RDB.Get(l.ctx, key).Result()
		if err == nil {
			return lURL, nil
		}
		if err != nil && err != redis.Nil {
			return nil, err
		}

		// 查询短链接是否存在
		suMap, err := su.WithContext(l.ctx).Where(su.Surl.Eq(req.ShortURL)).First()
		if err != nil {
			logx.Errorf("su.WithContext(l.ctx).Where(su.Surl.Eq(req.ShortURL)).First() error: %v", err)
			return nil, err
		}
		if suMap.ID == 0 {
			return nil, ErrShortUrlNotExist
		}
		// 写入缓存
		go func() {
			if err := query.RDB.Set(l.ctx, key, suMap.Lurl, randtime.GetRandTime()).Err(); err != nil {
				logx.Errorf("query.RDB.Set(l.ctx, key, suMap.Lurl, randtime.GetRandTime()).Err() error: %v", err)
			}
		}()
		return suMap.Lurl, nil
	})
	if err != nil {
		return nil, err
	}

	return &types.ShowResponse{LongURL: lURL.(string)}, nil
}
