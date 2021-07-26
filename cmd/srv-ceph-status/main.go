package main

import (
	"github.com/liucxer/srv-ceph-status/api"
	"github.com/liucxer/srv-ceph-status/cmd/srv-ceph-status/global"
	ceph_status "github.com/liucxer/srv-ceph-status/pkg/mgrs/ceph-status"
	"context"
	"github.com/go-courier/courier"
)

func main() {
	go func() {
		ctx := global.WithContext(context.Background())
		go ceph_status.ReFlashNodeList(ctx)
		go ceph_status.StartCollectStatus(ctx)
	}()

	global.App.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})

	global.App.Execute(func(args ...string) {
		courier.Run(api.RootRouter, global.Server())
	})
}
