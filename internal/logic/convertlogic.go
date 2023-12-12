package logic

import (
	"context"
	"errors"

	"shorturl/dal"
	"shorturl/dal/query"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"shorturl/pkg/connect"
	"shorturl/pkg/md5"
	"shorturl/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

const (
	dsn = "root:123456@tcp(127.0.0.1:3306)/shorturl?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func init() {
	dal.DB = dal.ConnectDB(dsn)
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	query.SetDefault(dal.DB)
	su := query.ShortURLMap

	// 判断链接是否已经转过
	md5Value := md5.Sum([]byte(req.LongURL))
	url, err := su.WithContext(l.ctx).Where(su.Md5.Eq(md5Value)).First()
	if err != gorm.ErrRecordNotFound {
		if err == nil {
			return &types.ConvertResponse{ShortURL: url.Surl}, nil
		}
		logx.Errorf("su.WithContext(l.ctx).Where(su.Md5.Eq(md5Value)).First() error: %v", err)
		return nil, err
	}

	// 判断链接是否为短链接
	baseURL, err := urltool.GetBaseURL(req.LongURL)
	if err != nil {
		logx.Errorf("urltool.GetBaseURL error: %v", err)
		return nil, err
	}
	if _, err = su.WithContext(l.ctx).Where(su.Surl.Eq(baseURL)).First(); err != gorm.ErrRecordNotFound {
		if err == nil {
			return &types.ConvertResponse{ShortURL: baseURL}, nil
		}
		logx.Errorf("su.WithContext(l.ctx).Where(su.Surl.Eq(baseURL)).First() error: %v", err)
		return nil, err
	}

	// 判断链接是否有效
	if ok := connect.Get(req.LongURL); !ok {
		return nil, errors.New("链接无效")
	}

	// 取号
	seqID, err := GetSeqID(l.ctx)
	if err != nil {
		logx.Errorf("GetSeqID error: %v", err)
		return nil, err
	}
	
	return nil, nil
}

func GetSeqID(ctx context.Context) (uint64, error) {
	tx := query.Sequence.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true})
	if err := tx.ReplaceStub(); err != nil {
		return 0, err
	}

	return tx.LastInsertID()
}
