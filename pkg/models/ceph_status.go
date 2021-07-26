package models

import "github.com/liucxer/srv-ceph-status/pkg/tools"

//go:generate codegen gen model2 CephStatus --database DBCephStatus --with-comments
// @def primary ID
// @def unique_index uk_ceph_status_id CephStatusID
// @def index idx_created_at CreatedAt
type CephStatus struct {
	PrimaryID
	RefNodeID
	RefCephStatusID

	CephStatusBody
	OperationTimesWithDeletedAt
}

type RefCephStatusID struct {
	CephStatusID tools.SFID `db:"f_ceph_status_id" json:"cephStatusID"`
}

type CephStatusBody struct {
	ReadBytesSec          float64 `db:"f_read_bytes_sec,default='0'" json:"readBytesSec"`
	ReadOpPerSec          float64 `db:"f_read_op_per_sec,default='0'" json:"readOpPerSec"`
	RecoveringBytesPerSec float64 `db:"f_recovering_bytes_per_sec,default='0'" json:"recoveringBytesPerSec"`
	WriteBytesSec         float64 `db:"f_write_bytes_sec,default='0'" json:"writeBytesSec"`
	WriteOpPerSec         float64 `db:"f_write_op_per_sec,default='0'" json:"writeOpPerSec"`
}
