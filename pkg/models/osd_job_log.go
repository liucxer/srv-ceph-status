package models

import (
	"github.com/liucxer/courier/sqlx/datatypes"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
)

//go:generate codegen gen model2 OSDJobLog --database DBCephStatus --with-comments
// @def primary ID
// @def unique_index uk_osd_job_log_id OsdJobLogID
// @def index idx_created_at CreatedAt
type OSDJobLog struct {
	PrimaryID
	RefNodeID
	RefOSDJobLogID

	OSDJobLogBody
	OperationTimesWithDeletedAt
}

type RefOSDJobLogID struct {
	OsdJobLogID tools.SFID `db:"f_osd_job_log_id" json:"osdJobLogID"`
}

type OSDJobLogBody struct {
	OSDID              int64               `db:"f_osd_id" json:"osdID"`
	Content            string              `db:"f_content" json:"content"`
	LogTime            datatypes.Timestamp `db:"f_log_time" json:"logTime"`
	JobID              string              `db:"f_job_id,default=''" json:"jobID"`
	GlobalID           string              `db:"f_global_id,default=''" json:"globalID"`
	JobType            string              `db:"f_job_type,default=''" json:"jobType"`
	Size               int64               `db:"f_size,default='0'" json:"size"`
	ExpectCost         float64             `db:"f_expect_cost,default='0'" json:"expectCost"`
	OutQueueActualCost float64             `db:"f_out_queue_actual_cost,default='0'" json:"outQueueActualCost"`
	ActualCost         float64             `db:"f_actual_cost,default='0'" json:"actualCost"`
}
