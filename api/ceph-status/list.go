package ceph_status

import (
	ceph_status "github.com/liucxer/srv-ceph-status/pkg/mgrs/ceph-status"
	"context"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
	"github.com/go-courier/sqlx-pg/pgbuilder"
)

func init() {
	RouterRootCephStatus.Register(courier.NewRouter(&ListCephStatus{}))
}

type ListCephStatus struct {
	httpx.MethodGet `summary:"状态列表" path:""`
	ceph_status.ListCephStatusParam
	pgbuilder.Pager
}

func (req *ListCephStatus) Output(ctx context.Context) (interface{}, error) {
	return ceph_status.ListCephStatus(ctx, req.ListCephStatusParam)
}
