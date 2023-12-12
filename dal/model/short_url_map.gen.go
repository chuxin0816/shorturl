// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameShortURLMap = "short_url_map"

// ShortURLMap ⻓短链映射表
type ShortURLMap struct {
	ID       int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                      // 主键
	CreateAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"` // 创建时间
	CreateBy string    `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                            // 创建者
	IsDel    int32     `gorm:"column:is_del;not null;comment:是否删除：0正常1删除" json:"is_del"`                          // 是否删除：0正常1删除
	Lurl     string    `gorm:"column:lurl;comment:⻓链接" json:"lurl"`                                               // ⻓链接
	Md5      string    `gorm:"column:md5;comment:⻓链接MD5" json:"md5"`                                              // ⻓链接MD5
	Surl     string    `gorm:"column:surl;comment:短链接" json:"surl"`                                               // 短链接
}

// TableName ShortURLMap's table name
func (*ShortURLMap) TableName() string {
	return TableNameShortURLMap
}
