package ceph_status

import (
	"context"
	"github.com/go-courier/httptransport/__examples__/server/pkg/errors"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/go-courier/sqlx/v2/datatypes"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
)

func CreateCephStatus(ctx context.Context, nodeID tools.SFID, body models.CephStatusBody) (*models.CephStatus, error) {
	taskID, err := idgen.IDGenFromContext(ctx).ID()
	if err != nil {
		return nil, err
	}

	object := models.CephStatus{}
	object.CephStatusID = tools.SFID(taskID)
	object.NodeID = nodeID
	object.CephStatusBody = body
	if err = object.Create(db.DBExecutorFromContext(ctx)); err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return nil, err
		}
		return nil, errors.InternalServerError
	}

	return &object, nil
}

type ListCephStatusParam struct {
	NodeID    tools.SFID          `name:"nodeID" in:"query"`
	StartTime datatypes.Timestamp `name:"startTime,omitempty" in:"query"`
	EndTime   datatypes.Timestamp `name:"endTime,omitempty" in:"query"`
}

func ListCephStatus(ctx context.Context, listParam ListCephStatusParam) (*[]models.CephStatus, error) {
	var (
		err  error
		list []models.CephStatus
	)

	object := models.CephStatus{}

	cond := builder.EmptyCond()
	if listParam.NodeID != 0 {
		cond = cond.And(object.FieldNodeID().Eq(listParam.NodeID))
	}
	if !listParam.StartTime.IsZero() {
		cond = cond.And(object.FieldCreatedAt().Gte(listParam.StartTime))
	}
	if !listParam.EndTime.IsZero() {
		cond = cond.And(object.FieldCreatedAt().Lte(listParam.EndTime))
	}

	if list, err = object.List(db.DBExecutorFromContext(ctx), cond); err != nil {
		return nil, errors.InternalServerError
	}

	return &list, nil
}
