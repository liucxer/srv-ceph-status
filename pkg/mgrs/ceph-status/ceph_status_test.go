package ceph_status_test

import (
	"github.com/liucxer/ceph-tools/pkg/cluster_client"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCollectStatus(t *testing.T) {
	cluster, err := cluster_client.NewCluster([]string{"10.0.20.27"})
	require.NoError(t, err)

	for i :=0; i < 100000;i++ {
		_, err = cluster.CephStatus(1)
	}

	require.NoError(t, err)
}
