package main

import (
	"context"
	"github.com/go-courier/courier"
	"github.com/liucxer/srv-ceph-status/api"
	"github.com/liucxer/srv-ceph-status/cmd/srv-ceph-status/global"
	ceph_status "github.com/liucxer/srv-ceph-status/pkg/mgrs/ceph-status"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
	osd_log "github.com/liucxer/srv-ceph-status/pkg/mgrs/osd-log"
)

func main() {
	go func() {
		ctx := global.WithContext(context.Background())
		go node.ReFlashNodeList(ctx)
		go ceph_status.StartCollectCephStatus(ctx)
		go osd_log.StartCollectOSDLog(ctx)
	}()

	global.App.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})

	global.App.Execute(func(args ...string) {
		courier.Run(api.RootRouter, global.Server())
	})
}
