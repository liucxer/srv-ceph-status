package node

import (
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"context"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterRootNode.Register(courier.NewRouter(&CreateNode{}))
}

type CreateNode struct {
	httpx.MethodPost `summary:"创建节点" path:""`
	models.NodeBody  `in:"body"`
}

func (req *CreateNode) Output(ctx context.Context) (interface{}, error) {
	return node.CreateNode(ctx, req.NodeBody)
}
