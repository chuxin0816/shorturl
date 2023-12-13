package logic

import (
	"context"
	"errors"

	"shorturl/dal/query"
	"shorturl/internal/svc"
	"shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	ErrShortUrlNotExist = errors.New("short url not exist")
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
	// 查询短链接是否存在
	suMap, err := su.WithContext(l.ctx).Where(su.Surl.Eq(req.ShortURL)).First()
	if err != nil {
		logx.Errorf("su.WithContext(l.ctx).Where(su.Surl.Eq(req.ShortURL)).First() error: %v", err)
		return nil, err
	}
	if suMap.ID == 0 {
		return nil, ErrShortUrlNotExist
	}

	return &types.ShowResponse{LongURL: suMap.Lurl}, nil
}
