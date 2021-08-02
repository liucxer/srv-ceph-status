package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/liucxer/courier/enumeration"
)

var InvalidOSDJobStatus = errors.New("invalid OSDJobStatus type")

func ParseOSDJobStatusFromLabelString(s string) (OSDJobStatus, error) {
	switch s {
	case "":
		return OSD_JOB_STATUS_UNKNOWN, nil
	case "等待执行":
		return OSD_JOB_STATUS__WAIT, nil
	case "执行中":
		return OSD_JOB_STATUS__RUNNING, nil
	case "执行成功":
		return OSD_JOB_STATUS__SUCCESS, nil
	case "执行失败":
		return OSD_JOB_STATUS__FAILURE, nil
	}
	return OSD_JOB_STATUS_UNKNOWN, InvalidOSDJobStatus
}

func (v OSDJobStatus) String() string {
	switch v {
	case OSD_JOB_STATUS_UNKNOWN:
		return ""
	case OSD_JOB_STATUS__WAIT:
		return "WAIT"
	case OSD_JOB_STATUS__RUNNING:
		return "RUNNING"
	case OSD_JOB_STATUS__SUCCESS:
		return "SUCCESS"
	case OSD_JOB_STATUS__FAILURE:
		return "FAILURE"
	}
	return "UNKNOWN"
}

func ParseOSDJobStatusFromString(s string) (OSDJobStatus, error) {
	switch s {
	case "":
		return OSD_JOB_STATUS_UNKNOWN, nil
	case "WAIT":
		return OSD_JOB_STATUS__WAIT, nil
	case "RUNNING":
		return OSD_JOB_STATUS__RUNNING, nil
	case "SUCCESS":
		return OSD_JOB_STATUS__SUCCESS, nil
	case "FAILURE":
		return OSD_JOB_STATUS__FAILURE, nil
	}
	return OSD_JOB_STATUS_UNKNOWN, InvalidOSDJobStatus
}

func (v OSDJobStatus) Label() string {
	switch v {
	case OSD_JOB_STATUS_UNKNOWN:
		return ""
	case OSD_JOB_STATUS__WAIT:
		return "等待执行"
	case OSD_JOB_STATUS__RUNNING:
		return "执行中"
	case OSD_JOB_STATUS__SUCCESS:
		return "执行成功"
	case OSD_JOB_STATUS__FAILURE:
		return "执行失败"
	}
	return "UNKNOWN"
}

func (v OSDJobStatus) Int() int {
	return int(v)
}

func (OSDJobStatus) TypeName() string {
	return "github.com/liucxer/srv-ceph-status/pkg/constants/types.OSDJobStatus"
}

func (OSDJobStatus) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{OSD_JOB_STATUS__WAIT, OSD_JOB_STATUS__RUNNING, OSD_JOB_STATUS__SUCCESS, OSD_JOB_STATUS__FAILURE}
}

func (v OSDJobStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidOSDJobStatus
	}
	return []byte(str), nil
}

func (v *OSDJobStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseOSDJobStatusFromString(string(bytes.ToUpper(data)))
	return
}

func (v OSDJobStatus) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *OSDJobStatus) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = OSDJobStatus(i)
	return nil
}
