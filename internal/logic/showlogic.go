package logic

import (
	"context"
	"errors"
	"time"

	"shorturl/dal/query"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"shorturl/pkg/randtime"

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
	// 使用布隆过滤器避免缓存穿透
	if !l.svcCtx.Bloom.Test([]byte(req.ShortURL)) {
		return nil, ErrShortUrlNotExist
	}
	su := query.ShortURLMap
	// 使用singleflight减轻缓存压力，并解决缓存击穿问题
	key := query.GetPrefix(req.ShortURL)
	lURL, err, _ := l.svcCtx.G.Do(key, func() (interface{}, error) {
		go func() {
			time.Sleep(delayTime)
			l.svcCtx.G.Forget(key)
		}()
		// 查询缓存
		lURL, err := l.svcCtx.RDB.Get(l.ctx, key).Result()
		if err == nil {
			return lURL, nil
		}
		if err != nil && err != redis.Nil {
			logx.Errorf("l.svcCtx.RDB.Get(context.Background(), key).Result() error: %v", err)
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
			if err := l.svcCtx.RDB.Set(l.ctx, key, suMap.Lurl, randtime.GetRandTime()).Err(); err != nil {
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
