package ceph_status

import (
	"context"
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport/httpx"
	"github.com/liucxer/courier/sqlx-pg/pgbuilder"
	ceph_status "github.com/liucxer/srv-ceph-status/pkg/mgrs/ceph-status"
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
	return ceph_status.ListCephStatus(ctx, req.ListCephStatusParam, req.Pager)
}
