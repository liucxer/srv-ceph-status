package osd_job

import (
	"context"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/liucxer/srv-ceph-status/pkg/constants/types"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
	"github.com/sirupsen/logrus"
)

func CreateOsdJob(ctx context.Context, body models.OSDJobBody) (*models.OSDJob, error) {
	var (
		err error
		job models.OSDJob
	)

	sfid, err := idgen.IDGenFromContext(ctx).ID()
	if err != nil {
		return nil, err
	}

	job.OsdJobID = tools.SFID(sfid)
	job.OSDJobBody = body
	job.OSDJobStatus = types.OSD_JOB_STATUS__WAIT
	err = job.Create(db.DBExecutorFromContext(ctx))
	if err != nil {
		logrus.Errorf("job.Create err:%v", err)
		return nil, err
	}

	return &job, err
}

func PickUpOsdJob(ctx context.Context) (*models.OSDJob, error) {
	var (
		err error
	)

	cond := builder.EmptyCond()
	job := models.OSDJob{}
	limitAddition := builder.Limit(1)
	cond = cond.And(job.FieldOSDJobStatus().In(types.OSD_JOB_STATUS__WAIT, types.OSD_JOB_STATUS__FAILURE))
	jobList, err := job.List(db.DBExecutorFromContext(ctx), cond, limitAddition)
	if err != nil {
		logrus.Errorf("job.List err:%v", err)
		return nil, err
	}

	if len(jobList) == 0 {
		logrus.Debug("len(taskList) == 0")
		return nil, nil
	}

	job.OsdJobID = jobList[0].OsdJobID

	tasks := sqlx.NewTasks(db.DBExecutorFromContext(ctx))
	pickTaskConflict := false
	fn := func(db sqlx.DBExecutor) error {
		err := job.FetchByOsdJobIDForUpdate(db)
		if err != nil {
			logrus.Errorf("job.FetchByOsdJobIDForUpdate err:%v", err)
			return err
		}
		if job.OSDJobStatus != types.OSD_JOB_STATUS__WAIT && job.OSDJobStatus != types.OSD_JOB_STATUS__FAILURE {
			pickTaskConflict = true
			return nil
		}
		job.OSDJobStatus = types.OSD_JOB_STATUS__RUNNING
		err = job.UpdateByIDWithStruct(db)
		if err != nil {
			logrus.Errorf("job.UpdateByIDWithStruct err:%v", err)
			return err
		}
		return nil
	}
	tasks = tasks.With(fn)

	err = tasks.Do()
	if err != nil {
		return nil, err
	}

	if pickTaskConflict {
		return nil, nil
	}
	return &job, nil
}
