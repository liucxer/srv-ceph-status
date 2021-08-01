package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidDiskType = errors.New("invalid DiskType type")

func ParseDiskTypeFromLabelString(s string) (DiskType, error) {
	switch s {
	case "":
		return DISK_TYPE_UNKNOWN, nil
	case "普通磁盘":
		return DISK_TYPE__HDD, nil
	case "固态硬盘":
		return DISK_TYPE__SSD, nil
	}
	return DISK_TYPE_UNKNOWN, InvalidDiskType
}

func (v DiskType) String() string {
	switch v {
	case DISK_TYPE_UNKNOWN:
		return ""
	case DISK_TYPE__HDD:
		return "HDD"
	case DISK_TYPE__SSD:
		return "SSD"
	}
	return "UNKNOWN"
}

func ParseDiskTypeFromString(s string) (DiskType, error) {
	switch s {
	case "":
		return DISK_TYPE_UNKNOWN, nil
	case "HDD":
		return DISK_TYPE__HDD, nil
	case "SSD":
		return DISK_TYPE__SSD, nil
	}
	return DISK_TYPE_UNKNOWN, InvalidDiskType
}

func (v DiskType) Label() string {
	switch v {
	case DISK_TYPE_UNKNOWN:
		return ""
	case DISK_TYPE__HDD:
		return "普通磁盘"
	case DISK_TYPE__SSD:
		return "固态硬盘"
	}
	return "UNKNOWN"
}

func (v DiskType) Int() int {
	return int(v)
}

func (DiskType) TypeName() string {
	return "github.com/liucxer/srv-ceph-status/pkg/constants/types.DiskType"
}

func (DiskType) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{DISK_TYPE__HDD, DISK_TYPE__SSD}
}

func (v DiskType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDiskType
	}
	return []byte(str), nil
}

func (v *DiskType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDiskTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v DiskType) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *DiskType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = DiskType(i)
	return nil
}
