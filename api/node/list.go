package node

import (
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
	"context"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
	"github.com/go-courier/sqlx-pg/pgbuilder"
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