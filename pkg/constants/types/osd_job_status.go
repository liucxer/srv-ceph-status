package types

//go:generate codegen gen enum OSDJobStatus
type OSDJobStatus uint8

const (
	OSD_JOB_STATUS_UNKNOWN  OSDJobStatus = iota
	OSD_JOB_STATUS__WAIT                 // 等待执行
	OSD_JOB_STATUS__RUNNING              // 执行中
	OSD_JOB_STATUS__SUCCESS              // 执行成功
	OSD_JOB_STATUS__FAILURE              // 执行失败
)
