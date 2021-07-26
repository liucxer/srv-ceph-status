package ceph_status

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var RouterRootCephStatus = courier.NewRouter(httptransport.Group("/ceph-status"))
