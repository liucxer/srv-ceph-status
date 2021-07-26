package models

import (
	fmt "fmt"
	time "time"

	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
	github_com_liucxer_srv_ceph_status_pkg_tools "github.com/liucxer/srv-ceph-status/pkg/tools"
)

func (Node) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Node) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"idx_created_at": []string{
			"CreatedAt",
		},
	}
}

func (Node) UniqueIndexUkNodeID() string {
	return "uk_node_id"
}

func (Node) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"uk_node_id": []string{
			"NodeID",
			"DeletedAt",
		},
	}
}

func (Node) Comments() map[string]string {
	return map[string]string{
		"CreatedAt": "创建时间",
		"DeletedAt": "删除时间",
		"ID":        "自增ID",
		"UpdatedAt": "更新时间",
	}
}

var NodeTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	NodeTable = DBCephStatus.Register(&Node{})
}

type NodeIterator struct {
}

func (NodeIterator) New() interface{} {
	return &Node{}
}

func (NodeIterator) Resolve(v interface{}) *Node {
	return v.(*Node)
}

func (Node) TableName() string {
	return "t_node"
}

func (Node) ColDescriptions() map[string][]string {
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

func (Node) FieldKeyID() string {
	return "ID"
}

func (m *Node) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyID())
}

func (Node) FieldKeyNodeID() string {
	return "NodeID"
}

func (m *Node) FieldNodeID() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyNodeID())
}

func (Node) FieldKeyAddress() string {
	return "Address"
}

func (m *Node) FieldAddress() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyAddress())
}

func (Node) FieldKeyAddressList() string {
	return "AddressList"
}

func (m *Node) FieldAddressList() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyAddressList())
}

func (Node) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Node) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyCreatedAt())
}

func (Node) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Node) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyUpdatedAt())
}

func (Node) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *Node) FieldDeletedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return NodeTable.F(m.FieldKeyDeletedAt())
}

func (Node) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Node) IndexFieldNames() []string {
	return []string{
		"CreatedAt",
		"ID",
		"NodeID",
	}
}

func (m *Node) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
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

func (m *Node) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Node) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

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

func (m *Node) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.DeleteByStruct"),
			),
	)

	return err
}

func (m *Node) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Node) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("Node.UpdateByIDWithMap"),
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

func (m *Node) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Node) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("Node.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Node) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.DeleteByID"),
			))

	return err
}

func (m *Node) SoftDeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("Node.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Node) FetchByNodeID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("NodeID").Eq(m.NodeID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.FetchByNodeID"),
			),
		m,
	)

	return err
}

func (m *Node) UpdateByNodeIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("NodeID").Eq(m.NodeID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.UpdateByNodeIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByNodeID(db)
	}

	return nil

}

func (m *Node) UpdateByNodeIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByNodeIDWithMap(db, fieldValues)

}

func (m *Node) FetchByNodeIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("NodeID").Eq(m.NodeID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.FetchByNodeIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Node) DeleteByNodeID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("NodeID").Eq(m.NodeID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.DeleteByNodeID"),
			))

	return err
}

func (m *Node) SoftDeleteByNodeID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

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
					table.F("NodeID").Eq(m.NodeID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Node.SoftDeleteByNodeID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Node) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Node, error) {

	list := make([]Node, 0)

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Node.List"),
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

func (m *Node) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_go_courier_sqlx_v2_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Node.Count"),
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

func (m *Node) BatchFetchByCreatedAtList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_go_courier_sqlx_v2_datatypes.Timestamp) ([]Node, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *Node) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Node, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Node) BatchFetchByNodeIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []github_com_liucxer_srv_ceph_status_pkg_tools.SFID) ([]Node, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("NodeID").In(values)

	return m.List(db, condition)

}
