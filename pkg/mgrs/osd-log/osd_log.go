package osd_log

import (
	"context"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"time"
)

func StartCollectOSDLog(ctx context.Context)  {
	for {
		time.Sleep(time.Second)
		node.NodeListMutex.Lock()
		var nodeList []models.Node
		nodeList = node.NodeList

		for _, nodeItem := range nodeList {
			go CollectOSDLog(ctx, nodeItem)
		}
		node.NodeListMutex.Unlock()
	}
}

func CollectOSDLog(ctx context.Context, node models.Node) error {
	// cat ceph-osd.7.log > ./ceph-osd.7.log.tmp && echo "" > ceph-osd.7.log
	//cluster := cluster_client.NewCluster()

	// 日志导出，并清空历史日志
	// 日志拉取
	// 日志分析
	// 入库
	return nil
}