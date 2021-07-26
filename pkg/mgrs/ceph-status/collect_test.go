package ceph_status_test

import (
	"context"
	ceph_status "github.com/liucxer/srv-ceph-status/pkg/mgrs/ceph-status"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCollectStatus2(t *testing.T) {
	node := models.Node{}
	node.Address = "10.0.8.44"
	err := ceph_status.CollectStatus(context.Background(), node)
	require.NoError(t, err)
}
