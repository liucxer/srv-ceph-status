package models

import (
	fmt "fmt"
	time "time"

	github_com_go_courier_sqlx_v2 "github.com/liucxer/courier/sqlx"
	github_com_go_courier_sqlx_v2_builder "github.com/liucxer/courier/sqlx/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/liucxer/courier/sqlx/datatypes"
	github_com_liucxer_srv_ceph_status_pkg_tools "github.com/liucxer/srv-ceph-status/pkg/tools"
)

func (CephStatus) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (CephStatus) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"idx_created_at": []string{
			"CreatedAt",
		},
	}
}

func (CephStatus) UniqueIndexUkCephStatusID() string {
	return "uk_ceph_status_id"
}

func (CephStatus) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"uk_ceph_status_id": []string{
			"CephStatusID",
			"DeletedAt",
		},
	}
}

func (CephStatus) Comments() map[string]string {
	return map[string]string{
		"CreatedAt": "创建时间",
		"DeletedAt": "删除时间",
		"ID":        "自增ID",
		"UpdatedAt": "更新时间",
	}
}

var CephStatusTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	CephStatusTable = DBCephStatus.Register(&CephStatus{})
}

type CephStatusIterator struct {
}

func (CephStatusIterator) New() interface{} {
	return &CephStatus{}
}

func (CephStatusIterator) Resolve(v interface{}) *CephStatus {
	return v.(*CephStatus)
}

func (CephStatus) TableName() string {
	return "t_ceph_status"
}

func (CephStatus) ColDescriptions() map[string][]string {
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

func (CephStatus) FieldKeyID() string {
	return "ID"
}

func (m *CephStatus) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyID())
}

func (CephStatus) FieldKeyNodeID() string {
	return "NodeID"
}

func (m *CephStatus) FieldNodeID() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyNodeID())
}

func (CephStatus) FieldKeyCephStatusID() string {
	return "CephStatusID"
}

func (m *CephStatus) FieldCephStatusID() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyCephStatusID())
}

func (CephStatus) FieldKeyReadBytesSec() string {
	return "ReadBytesSec"
}

func (m *CephStatus) FieldReadBytesSec() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyReadBytesSec())
}

func (CephStatus) FieldKeyReadOpPerSec() string {
	return "ReadOpPerSec"
}

func (m *CephStatus) FieldReadOpPerSec() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyReadOpPerSec())
}

func (CephStatus) FieldKeyRecoveringBytesPerSec() string {
	return "RecoveringBytesPerSec"
}

func (m *CephStatus) FieldRecoveringBytesPerSec() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyRecoveringBytesPerSec())
}

func (CephStatus) FieldKeyWriteBytesSec() string {
	return "WriteBytesSec"
}

func (m *CephStatus) FieldWriteBytesSec() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyWriteBytesSec())
}

func (CephStatus) FieldKeyWriteOpPerSec() string {
	return "WriteOpPerSec"
}

func (m *CephStatus) FieldWriteOpPerSec() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyWriteOpPerSec())
}

func (CephStatus) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *CephStatus) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyCreatedAt())
}

func (CephStatus) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *CephStatus) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyUpdatedAt())
}

func (CephStatus) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *CephStatus) FieldDeletedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return CephStatusTable.F(m.FieldKeyDeletedAt())
}

func (CephStatus) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *CephStatus) IndexFieldNames() []string {
	return []string{
		"CephStatusID",
		"CreatedAt",
		"ID",
	}
}

func (m *CephStatus) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
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

func (m *CephStatus) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *CephStatus) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

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

func (m *CephStatus) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.DeleteByStruct"),
			),
	)

	return err
}

func (m *CephStatus) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.FetchByID"),
			),
		m,
	)

	return err
}

func (m *CephStatus) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.UpdateByIDWithMap"),
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

func (m *CephStatus) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *CephStatus) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *CephStatus) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.DeleteByID"),
			))

	return err
}

func (m *CephStatus) SoftDeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *CephStatus) FetchByCephStatusID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("CephStatusID").Eq(m.CephStatusID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.FetchByCephStatusID"),
			),
		m,
	)

	return err
}

func (m *CephStatus) UpdateByCephStatusIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("CephStatusID").Eq(m.CephStatusID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.UpdateByCephStatusIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByCephStatusID(db)
	}

	return nil

}

func (m *CephStatus) UpdateByCephStatusIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByCephStatusIDWithMap(db, fieldValues)

}

func (m *CephStatus) FetchByCephStatusIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("CephStatusID").Eq(m.CephStatusID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.FetchByCephStatusIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *CephStatus) DeleteByCephStatusID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("CephStatusID").Eq(m.CephStatusID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.DeleteByCephStatusID"),
			))

	return err
}

func (m *CephStatus) SoftDeleteByCephStatusID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
					table.F("CephStatusID").Eq(m.CephStatusID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.SoftDeleteByCephStatusID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *CephStatus) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]CephStatus, error) {

	list := make([]CephStatus, 0)

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.List"),
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

func (m *CephStatus) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("CephStatus.Count"),
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

func (m *CephStatus) BatchFetchByCephStatusIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_liucxer_srv_ceph_status_pkg_tools.SFID) ([]CephStatus, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CephStatusID").In(values)

	return m.List(db, condition)

}

func (m *CephStatus) BatchFetchByCreatedAtList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_go_courier_sqlx_v2_datatypes.Timestamp) ([]CephStatus, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *CephStatus) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]CephStatus, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}
