package models

import (
	fmt "fmt"
	time "time"

	github_com_go_courier_sqlx_v2 "github.com/liucxer/courier/sqlx"
	github_com_go_courier_sqlx_v2_builder "github.com/liucxer/courier/sqlx/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/liucxer/courier/sqlx/datatypes"
	github_com_liucxer_srv_ceph_status_pkg_tools "github.com/liucxer/srv-ceph-status/pkg/tools"
)

func (OSDJob) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (OSDJob) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"idx_created_at": []string{
			"CreatedAt",
		},
	}
}

func (OSDJob) UniqueIndexUkOsdJobID() string {
	return "uk_osd_job_id"
}

func (OSDJob) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"uk_osd_job_id": []string{
			"OsdJobID",
			"DeletedAt",
		},
	}
}

func (OSDJob) Comments() map[string]string {
	return map[string]string{
		"CreatedAt": "创建时间",
		"DeletedAt": "删除时间",
		"ID":        "自增ID",
		"UpdatedAt": "更新时间",
	}
}

var OSDJobTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	OSDJobTable = DBCephStatus.Register(&OSDJob{})
}

type OSDJobIterator struct {
}

func (OSDJobIterator) New() interface{} {
	return &OSDJob{}
}

func (OSDJobIterator) Resolve(v interface{}) *OSDJob {
	return v.(*OSDJob)
}

func (OSDJob) TableName() string {
	return "t_osd_job"
}

func (OSDJob) ColDescriptions() map[string][]string {
	return map[string][]string{
		"CreatedAt": []string{
			"创建时间",
		},
		"DeletedAt": []string{
			"删除时间",
		},
		"ID": []string{
			"自增ID",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
	}
}

func (OSDJob) FieldKeyID() string {
	return "ID"
}

func (m *OSDJob) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyID())
}

func (OSDJob) FieldKeyOsdJobID() string {
	return "OsdJobID"
}

func (m *OSDJob) FieldOsdJobID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyOsdJobID())
}

func (OSDJob) FieldKeyDiskType() string {
	return "DiskType"
}

func (m *OSDJob) FieldDiskType() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyDiskType())
}

func (OSDJob) FieldKeyOpType() string {
	return "OpType"
}

func (m *OSDJob) FieldOpType() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyOpType())
}

func (OSDJob) FieldKeyRuntime() string {
	return "Runtime"
}

func (m *OSDJob) FieldRuntime() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyRuntime())
}

func (OSDJob) FieldKeyBlockSize() string {
	return "BlockSize"
}

func (m *OSDJob) FieldBlockSize() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyBlockSize())
}

func (OSDJob) FieldKeyIoDepth() string {
	return "IoDepth"
}

func (m *OSDJob) FieldIoDepth() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyIoDepth())
}

func (OSDJob) FieldKeyRecoveryLimit() string {
	return "RecoveryLimit"
}

func (m *OSDJob) FieldRecoveryLimit() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyRecoveryLimit())
}

func (OSDJob) FieldKeyOSDJobStatus() string {
	return "OSDJobStatus"
}

func (m *OSDJob) FieldOSDJobStatus() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyOSDJobStatus())
}

func (OSDJob) FieldKeyResultTime() string {
	return "ResultTime"
}

func (m *OSDJob) FieldResultTime() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyResultTime())
}

func (OSDJob) FieldKeyDataPool() string {
	return "DataPool"
}

func (m *OSDJob) FieldDataPool() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyDataPool())
}

func (OSDJob) FieldKeyRecoveryPool() string {
	return "RecoveryPool"
}

func (m *OSDJob) FieldRecoveryPool() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyRecoveryPool())
}

func (OSDJob) FieldKeyDataVolume() string {
	return "DataVolume"
}

func (m *OSDJob) FieldDataVolume() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyDataVolume())
}

func (OSDJob) FieldKeyRecoveryVolume() string {
	return "RecoveryVolume"
}

func (m *OSDJob) FieldRecoveryVolume() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyRecoveryVolume())
}

func (OSDJob) FieldKeyOsdNum() string {
	return "OsdNum"
}

func (m *OSDJob) FieldOsdNum() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyOsdNum())
}

func (OSDJob) FieldKeyExpectCost() string {
	return "ExpectCost"
}

func (m *OSDJob) FieldExpectCost() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyExpectCost())
}

func (OSDJob) FieldKeyActualCost() string {
	return "ActualCost"
}

func (m *OSDJob) FieldActualCost() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyActualCost())
}

func (OSDJob) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *OSDJob) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyCreatedAt())
}

func (OSDJob) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *OSDJob) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyUpdatedAt())
}

func (OSDJob) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *OSDJob) FieldDeletedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobTable.F(m.FieldKeyDeletedAt())
}

func (OSDJob) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *OSDJob) IndexFieldNames() []string {
	return []string{
		"CreatedAt",
		"ID",
		"OsdJobID",
	}
}

func (m *OSDJob) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_go_courier_sqlx_v2_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_go_courier_sqlx_v2_builder.And(conditions...)

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))
	return condition
}

func (m *OSDJob) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *OSDJob) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_go_courier_sqlx_v2_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_go_courier_sqlx_v2_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_go_courier_sqlx_v2_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_go_courier_sqlx_v2_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_go_courier_sqlx_v2_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *OSDJob) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.DeleteByStruct"),
			),
	)

	return err
}

func (m *OSDJob) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.FetchByID"),
			),
		m,
	)

	return err
}

func (m *OSDJob) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *OSDJob) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *OSDJob) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OSDJob) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.DeleteByID"),
			))

	return err
}

func (m *OSDJob) SoftDeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OSDJob) FetchByOsdJobID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobID").Eq(m.OsdJobID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.FetchByOsdJobID"),
			),
		m,
	)

	return err
}

func (m *OSDJob) UpdateByOsdJobIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobID").Eq(m.OsdJobID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.UpdateByOsdJobIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByOsdJobID(db)
	}

	return nil

}

func (m *OSDJob) UpdateByOsdJobIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByOsdJobIDWithMap(db, fieldValues)

}

func (m *OSDJob) FetchByOsdJobIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobID").Eq(m.OsdJobID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.FetchByOsdJobIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OSDJob) DeleteByOsdJobID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobID").Eq(m.OsdJobID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.DeleteByOsdJobID"),
			))

	return err
}

func (m *OSDJob) SoftDeleteByOsdJobID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobID").Eq(m.OsdJobID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.SoftDeleteByOsdJobID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OSDJob) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]OSDJob, error) {

	list := make([]OSDJob, 0)

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *OSDJob) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("OSDJob.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(
			github_com_go_courier_sqlx_v2_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *OSDJob) BatchFetchByCreatedAtList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_go_courier_sqlx_v2_datatypes.Timestamp) ([]OSDJob, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *OSDJob) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]OSDJob, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *OSDJob) BatchFetchByOsdJobIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_liucxer_srv_ceph_status_pkg_tools.SFID) ([]OSDJob, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OsdJobID").In(values)

	return m.List(db, condition)

}
