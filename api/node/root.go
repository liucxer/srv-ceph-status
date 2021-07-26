package node

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var RouterRootNode = courier.NewRouter(httptransport.Group("/node"))
