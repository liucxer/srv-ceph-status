package node

import (
	"context"
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport/httpx"
	"github.com/liucxer/courier/sqlx-pg/pgbuilder"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
)

func init() {
	RouterRootNode.Register(courier.NewRouter(&ListNode{}))
}

type ListNode struct {
	httpx.MethodGet `summary:"节点列表" path:""`
	pgbuilder.Pager
}

func (req *ListNode) Output(ctx context.Context) (interface{}, error) {
	return node.ListNode(ctx)
}
