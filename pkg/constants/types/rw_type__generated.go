package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/liucxer/courier/enumeration"
)

var InvalidRWType = errors.New("invalid RWType type")

func ParseRWTypeFromLabelString(s string) (RWType, error) {
	switch s {
	case "":
		return RW_TYPE_UNKNOWN, nil
	case "顺序写":
		return RW_TYPE__READ, nil
	case "顺序读":
		return RW_TYPE__WRITE, nil
	case "随机写":
		return RW_TYPE__RANDWRITE, nil
	case "随机读":
		return RW_TYPE__RANDREAD, nil
	case "读写混合":
		return RW_TYPE__RW, nil
	case "随机读写混合":
		return RW_TYPE__RANDRW, nil
	}
	return RW_TYPE_UNKNOWN, InvalidRWType
}

func (v RWType) String() string {
	switch v {
	case RW_TYPE_UNKNOWN:
		return ""
	case RW_TYPE__READ:
		return "READ"
	case RW_TYPE__WRITE:
		return "WRITE"
	case RW_TYPE__RANDWRITE:
		return "RANDWRITE"
	case RW_TYPE__RANDREAD:
		return "RANDREAD"
	case RW_TYPE__RW:
		return "RW"
	case RW_TYPE__RANDRW:
		return "RANDRW"
	}
	return "UNKNOWN"
}

func ParseRWTypeFromString(s string) (RWType, error) {
	switch s {
	case "":
		return RW_TYPE_UNKNOWN, nil
	case "READ":
		return RW_TYPE__READ, nil
	case "WRITE":
		return RW_TYPE__WRITE, nil
	case "RANDWRITE":
		return RW_TYPE__RANDWRITE, nil
	case "RANDREAD":
		return RW_TYPE__RANDREAD, nil
	case "RW":
		return RW_TYPE__RW, nil
	case "RANDRW":
		return RW_TYPE__RANDRW, nil
	}
	return RW_TYPE_UNKNOWN, InvalidRWType
}

func (v RWType) Label() string {
	switch v {
	case RW_TYPE_UNKNOWN:
		return ""
	case RW_TYPE__READ:
		return "顺序写"
	case RW_TYPE__WRITE:
		return "顺序读"
	case RW_TYPE__RANDWRITE:
		return "随机写"
	case RW_TYPE__RANDREAD:
		return "随机读"
	case RW_TYPE__RW:
		return "读写混合"
	case RW_TYPE__RANDRW:
		return "随机读写混合"
	}
	return "UNKNOWN"
}

func (v RWType) Int() int {
	return int(v)
}

func (RWType) TypeName() string {
	return "github.com/liucxer/srv-ceph-status/pkg/constants/types.RWType"
}

func (RWType) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{RW_TYPE__READ, RW_TYPE__WRITE, RW_TYPE__RANDWRITE, RW_TYPE__RANDREAD, RW_TYPE__RW, RW_TYPE__RANDRW}
}

func (v RWType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRWType
	}
	return []byte(str), nil
}

func (v *RWType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRWTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v RWType) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *RWType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = RWType(i)
	return nil
}
