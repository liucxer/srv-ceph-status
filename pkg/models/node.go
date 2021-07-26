package models

import "github.com/liucxer/srv-ceph-status/pkg/tools"

//go:generate codegen gen model2 Node --database DBCephStatus --with-comments
// @def primary ID
// @def unique_index uk_node_id NodeID
type Node struct {
	PrimaryID
	RefNodeID
	NodeBody
	OperationTimesWithDeletedAt
}

type RefNodeID struct {
	NodeID tools.SFID `db:"f_node_id" json:"nodeID"`
}

type NodeBody struct {
	Address string `db:"f_address" json:"address"`
}
