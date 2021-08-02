package node

import (
	"context"
	"encoding/json"
	"github.com/liucxer/courier/sqlx"
	"github.com/liucxer/srv-ceph-status/pkg/constants/errors"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/tools"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
	"sync"
	"time"
)

var (
	NodeList      []models.Node
	NodeListMutex sync.Mutex
)

func ReFlashNodeList(ctx context.Context) {
	for {
		time.Sleep(3 * time.Second)
		nodeList, err := ListNode(ctx)
		if err != nil {
			continue
		}

		NodeListMutex.Lock()

		NodeList = []models.Node{}
		for _, nodeItem := range *nodeList {
			NodeList = append(NodeList, nodeItem)
		}
		NodeListMutex.Unlock()
	}
}

type CreateNodeBody struct {
	Address     string   `db:"f_address" json:"address"`
	AddressList []string `db:"f_address_list" json:"addressList"`
}

func CreateNode(ctx context.Context, body CreateNodeBody) (*models.Node, error) {
	taskID, err := idgen.IDGenFromContext(ctx).ID()
	if err != nil {
		return nil, err
	}

	object := models.Node{}
	object.NodeID = tools.SFID(taskID)
	object.Address = body.Address
	bts, err := json.Marshal(body.AddressList)
	if err != nil {
		return nil, err
	}
	object.AddressList = string(bts)

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
