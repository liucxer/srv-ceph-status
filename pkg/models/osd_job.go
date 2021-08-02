package models

import (
	"github.com/liucxer/courier/sqlx/datatypes"
	"github.com/liucxer/srv-ceph-status/pkg/constants/types"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
)

//go:generate codegen gen model2 OSDJob --database DBCephStatus --with-comments
// @def primary ID
// @def unique_index uk_osd_job_id OsdJobID
// @def index idx_created_at CreatedAt
type OSDJob struct {
	PrimaryID
	RefOSDJobID

	OSDJobBody
	OSDJobResult
	OperationTimesWithDeletedAt
}

type RefOSDJobID struct {
	OsdJobID tools.SFID `db:"f_osd_job_id" json:"osdJobID"`
}

type OSDJobBody struct {
	DiskType      types.DiskType     `db:"f_disk_type" json:"diskType"`
	OpType        types.RWType       `db:"f_op_type" json:"opType"`
	Runtime       int64              `db:"f_runtime" json:"runtime"`
	BlockSize     string             `db:"f_block_size" json:"blockSize"`
	IoDepth       int64              `db:"f_io_depth" json:"ioDepth"`
	RecoveryLimit float64            `db:"f_recovery_limit" json:"recoveryLimit"`
	OSDJobStatus  types.OSDJobStatus `db:"f_osd_job_status,default='1'" json:"osdJobStatus"`
}

type OSDJobResult struct {
	ResultTime     datatypes.Timestamp `db:"f_result_time,default='0'" json:"resultTime"`
	DataPool       string              `db:"f_data_pool,default=''" json:"dataPool"`
	RecoveryPool   string              `db:"f_recovery_pool,default=''" json:"recoveryPool"`
	DataVolume     string              `db:"f_data_volume,default=''" json:"dataVolume"`
	RecoveryVolume string              `db:"f_recovery_volume,default=''" json:"recoveryVolume"`
	OsdNum         string              `db:"f_osd_num,default=''" json:"osdNum"`
	ExpectCost     float64             `db:"f_expect_cost,default='0'" json:"expectCost"`
	ActualCost     float64             `db:"f_actual_cost,default='0'" json:"actualCost"`
}
