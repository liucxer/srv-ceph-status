package api

import (
	ceph_status "github.com/liucxer/srv-ceph-status/api/ceph-status"
	"github.com/liucxer/srv-ceph-status/api/node"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
	"github.com/go-courier/httptransport/openapi"
	"github.com/liucxer/confmiddleware/confhttp"
)

var (
	RootRouter = courier.NewRouter(httptransport.Group("/ceph-status"))
	V0Router   = courier.NewRouter(httptransport.Group("/v0"))
)

func init() {
	RootRouter.Register(openapi.OpenAPIRouter)
	RootRouter.Register(confhttp.LivenessRouter)
	RootRouter.Register(V0Router)
	V0Router.Register(node.RouterRootNode)
	V0Router.Register(ceph_status.RouterRootCephStatus)
}
