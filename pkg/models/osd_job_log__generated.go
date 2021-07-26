package models

import (
	fmt "fmt"
	time "time"

	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
	github_com_liucxer_srv_ceph_status_pkg_tools "github.com/liucxer/srv-ceph-status/pkg/tools"
)

func (OSDJobLog) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (OSDJobLog) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"idx_created_at": []string{
			"CreatedAt",
		},
	}
}

func (OSDJobLog) UniqueIndexUkOsdJobLogID() string {
	return "uk_osd_job_log_id"
}

func (OSDJobLog) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"uk_osd_job_log_id": []string{
			"OsdJobLogID",
			"DeletedAt",
		},
	}
}

func (OSDJobLog) Comments() map[string]string {
	return map[string]string{
		"CreatedAt": "创建时间",
		"DeletedAt": "删除时间",
		"ID":        "自增ID",
		"UpdatedAt": "更新时间",
	}
}

var OSDJobLogTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	OSDJobLogTable = DBCephStatus.Register(&OSDJobLog{})
}

type OSDJobLogIterator struct {
}

func (OSDJobLogIterator) New() interface{} {
	return &OSDJobLog{}
}

func (OSDJobLogIterator) Resolve(v interface{}) *OSDJobLog {
	return v.(*OSDJobLog)
}

func (OSDJobLog) TableName() string {
	return "t_osd_job_log"
}

func (OSDJobLog) ColDescriptions() map[string][]string {
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

func (OSDJobLog) FieldKeyID() string {
	return "ID"
}

func (m *OSDJobLog) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyID())
}

func (OSDJobLog) FieldKeyNodeID() string {
	return "NodeID"
}

func (m *OSDJobLog) FieldNodeID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyNodeID())
}

func (OSDJobLog) FieldKeyOsdJobLogID() string {
	return "OsdJobLogID"
}

func (m *OSDJobLog) FieldOsdJobLogID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyOsdJobLogID())
}

func (OSDJobLog) FieldKeyOSDID() string {
	return "OSDID"
}

func (m *OSDJobLog) FieldOSDID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyOSDID())
}

func (OSDJobLog) FieldKeyContent() string {
	return "Content"
}

func (m *OSDJobLog) FieldContent() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyContent())
}

func (OSDJobLog) FieldKeyLogTime() string {
	return "LogTime"
}

func (m *OSDJobLog) FieldLogTime() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyLogTime())
}

func (OSDJobLog) FieldKeyJobID() string {
	return "JobID"
}

func (m *OSDJobLog) FieldJobID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyJobID())
}

func (OSDJobLog) FieldKeyGlobalID() string {
	return "GlobalID"
}

func (m *OSDJobLog) FieldGlobalID() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyGlobalID())
}

func (OSDJobLog) FieldKeyJobType() string {
	return "JobType"
}

func (m *OSDJobLog) FieldJobType() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyJobType())
}

func (OSDJobLog) FieldKeySize() string {
	return "Size"
}

func (m *OSDJobLog) FieldSize() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeySize())
}

func (OSDJobLog) FieldKeyExpectCost() string {
	return "ExpectCost"
}

func (m *OSDJobLog) FieldExpectCost() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyExpectCost())
}

func (OSDJobLog) FieldKeyOutQueueActualCost() string {
	return "OutQueueActualCost"
}

func (m *OSDJobLog) FieldOutQueueActualCost() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyOutQueueActualCost())
}

func (OSDJobLog) FieldKeyActualCost() string {
	return "ActualCost"
}

func (m *OSDJobLog) FieldActualCost() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyActualCost())
}

func (OSDJobLog) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *OSDJobLog) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyCreatedAt())
}

func (OSDJobLog) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *OSDJobLog) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyUpdatedAt())
}

func (OSDJobLog) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *OSDJobLog) FieldDeletedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return OSDJobLogTable.F(m.FieldKeyDeletedAt())
}

func (OSDJobLog) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *OSDJobLog) IndexFieldNames() []string {
	return []string{
		"CreatedAt",
		"ID",
		"OsdJobLogID",
	}
}

func (m *OSDJobLog) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
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

func (m *OSDJobLog) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *OSDJobLog) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

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

func (m *OSDJobLog) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.DeleteByStruct"),
			),
	)

	return err
}

func (m *OSDJobLog) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.FetchByID"),
			),
		m,
	)

	return err
}

func (m *OSDJobLog) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.UpdateByIDWithMap"),
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

func (m *OSDJobLog) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *OSDJobLog) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OSDJobLog) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.DeleteByID"),
			))

	return err
}

func (m *OSDJobLog) SoftDeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OSDJobLog) FetchByOsdJobLogID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobLogID").Eq(m.OsdJobLogID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.FetchByOsdJobLogID"),
			),
		m,
	)

	return err
}

func (m *OSDJobLog) UpdateByOsdJobLogIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobLogID").Eq(m.OsdJobLogID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.UpdateByOsdJobLogIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByOsdJobLogID(db)
	}

	return nil

}

func (m *OSDJobLog) UpdateByOsdJobLogIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByOsdJobLogIDWithMap(db, fieldValues)

}

func (m *OSDJobLog) FetchByOsdJobLogIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobLogID").Eq(m.OsdJobLogID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.FetchByOsdJobLogIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OSDJobLog) DeleteByOsdJobLogID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("OsdJobLogID").Eq(m.OsdJobLogID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.DeleteByOsdJobLogID"),
			))

	return err
}

func (m *OSDJobLog) SoftDeleteByOsdJobLogID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
					table.F("OsdJobLogID").Eq(m.OsdJobLogID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.SoftDeleteByOsdJobLogID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OSDJobLog) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]OSDJobLog, error) {

	list := make([]OSDJobLog, 0)

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.List"),
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

func (m *OSDJobLog) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("OSDJobLog.Count"),
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

func (m *OSDJobLog) BatchFetchByCreatedAtList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_go_courier_sqlx_v2_datatypes.Timestamp) ([]OSDJobLog, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *OSDJobLog) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]OSDJobLog, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *OSDJobLog) BatchFetchByOsdJobLogIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_liucxer_srv_ceph_status_pkg_tools.SFID) ([]OSDJobLog, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OsdJobLogID").In(values)

	return m.List(db, condition)

}
