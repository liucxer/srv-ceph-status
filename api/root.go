package api

import (
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport"
	"github.com/liucxer/courier/httptransport/openapi"
	"github.com/liucxer/confmiddleware/confhttp"
	ceph_status "github.com/liucxer/srv-ceph-status/api/ceph-status"
	"github.com/liucxer/srv-ceph-status/api/node"
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
