package types

//go:generate codegen gen enum RWType
type RWType uint8

const (
	RW_TYPE_UNKNOWN    RWType = iota
	RW_TYPE__READ             // 顺序写
	RW_TYPE__WRITE            // 顺序读
	RW_TYPE__RANDWRITE        // 随机写
	RW_TYPE__RANDREAD         // 随机读
	RW_TYPE__RW               // 读写混合
	RW_TYPE__RANDRW           // 随机读写混合
)
