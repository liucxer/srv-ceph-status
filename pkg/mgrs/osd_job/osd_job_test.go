package osd_job_test

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/liucxer/srv-ceph-status/cmd/srv-ceph-status/global"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/osd_job"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOsdJob(t *testing.T) {
	ctx := global.WithContext(context.Background())
	body := models.OSDJobBody{
		Runtime: 200,
	}
	/*
		RecoveryLimit float64            `db:"f_recovery_limit" json:"recoveryLimit"`
	*/
	/*
		diskTypes := []types.DiskType{types.DISK_TYPE__HDD, types.DISK_TYPE__SSD}
		opTypes := []types.RWType{
			types.RW_TYPE__READ,
			types.RW_TYPE__WRITE,
			types.RW_TYPE__RANDWRITE,
			types.RW_TYPE__RANDREAD,
			types.RW_TYPE__RW,
			types.RW_TYPE__RANDRW,
		}
		blockSizes := []string{"4M", "2M", "1M", "256k", "64k", "32k", "4k"}
		ioDepthes := []int64{512, 256, 128, 64, 32, 16, 8, 1}
		recoveryLimits := []float64{2000, 1500, 1000, 500, 316, 158, 1}

	*/
	job, err := osd_job.CreateOsdJob(ctx, body)
	require.NoError(t, err)
	spew.Dump(job)
}

func TestPickUpOsdJob(t *testing.T) {
	ctx := global.WithContext(context.Background())
	job, err := osd_job.PickUpOsdJob(ctx)
	require.NoError(t, err)
	spew.Dump(job)
}
