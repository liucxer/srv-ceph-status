package node

import (
	"context"
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport/httpx"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
)

func init() {
	RouterRootNode.Register(courier.NewRouter(&CreateNode{}))
}

type CreateNode struct {
	httpx.MethodPost    `summary:"创建节点" path:""`
	node.CreateNodeBody `in:"body"`
}

func (req *CreateNode) Output(ctx context.Context) (interface{}, error) {
	return node.CreateNode(ctx, req.CreateNodeBody)
}
