package types

//go:generate codegen gen enum DiskType
type DiskType uint8

const (
	DISK_TYPE_UNKNOWN DiskType = iota
	DISK_TYPE__HDD             // 普通磁盘
	DISK_TYPE__SSD             // 固态硬盘
)
