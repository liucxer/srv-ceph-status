package node

import (
	"context"
	"github.com/go-courier/httptransport/__examples__/server/pkg/errors"
	"github.com/go-courier/sqlx/v2"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
)

func CreateNode(ctx context.Context, body models.NodeBody) (*models.Node, error) {
	taskID, err := idgen.IDGenFromContext(ctx).ID()
	if err != nil {
		return nil, err
	}

	object := models.Node{}
	object.NodeID = tools.SFID(taskID)
	object.NodeBody = body
	if err = object.Create(db.DBExecutorFromContext(ctx)); err != nil {
		if sqlx.DBErr(err).IsConflict() {
			return nil, err
		}
		return nil, errors.InternalServerError
	}

	return &object, nil
}

func ListNode(ctx context.Context) (*[]models.Node, error) {
	object := models.Node{}
	res, err := object.List(db.DBExecutorFromContext(ctx), nil)
	return &res, err
}
