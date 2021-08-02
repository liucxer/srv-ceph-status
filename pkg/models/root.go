package models

import (
	"github.com/liucxer/courier/sqlx"
	"github.com/liucxer/courier/sqlx/datatypes"
	"time"
)

var DBCephStatus = sqlx.NewDatabase("ceph_status")

type PrimaryID struct {
	// 自增ID
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

type OperationTimes struct {
	// 创建时间
	CreatedAt datatypes.Timestamp `db:"f_created_at,default='0'" json:"createdAt" `
	// 更新时间
	UpdatedAt datatypes.Timestamp `db:"f_updated_at,default='0'" json:"updatedAt"`
}

func (times *OperationTimes) MarkUpdatedAt() {
	times.UpdatedAt = datatypes.Timestamp(time.Now())
}

func (times *OperationTimes) MarkCreatedAt() {
	times.MarkUpdatedAt()
	times.CreatedAt = times.UpdatedAt
}

type OperationTimesWithDeletedAt struct {
	OperationTimes
	// 删除时间
	DeletedAt datatypes.Timestamp `db:"f_deleted_at,default='0'" json:"-"`
}

func (times *OperationTimesWithDeletedAt) MarkDeletedAt() {
	times.MarkUpdatedAt()
	times.DeletedAt = times.UpdatedAt
}
