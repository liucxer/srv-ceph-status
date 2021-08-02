package node

import (
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport"
)

var RouterRootNode = courier.NewRouter(httptransport.Group("/node"))
