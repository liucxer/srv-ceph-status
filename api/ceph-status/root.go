package ceph_status

import (
	"github.com/liucxer/courier/courier"
	"github.com/liucxer/courier/httptransport"
)

var RouterRootCephStatus = courier.NewRouter(httptransport.Group("/ceph-status"))
