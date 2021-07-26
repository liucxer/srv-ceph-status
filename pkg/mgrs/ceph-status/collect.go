package ceph_status

import (
	"context"
	"encoding/json"
	"github.com/ceph/go-ceph/rados"
	"github.com/google/uuid"
	"github.com/liucxer/srv-ceph-status/pkg/mgrs/node"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var (
	NodeList      []models.Node
	NodeListMutex sync.Mutex
)

func ReFlashNodeList(ctx context.Context) error {
	for {
		time.Sleep(3 * time.Second)
		nodeList, err := node.ListNode(ctx)
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
	return nil
}

func StartCollectStatus(ctx context.Context) error {
	for {
		time.Sleep(time.Second)
		NodeListMutex.Lock()
		var nodeList []models.Node
		nodeList = NodeList

		for _, nodeItem := range nodeList {
			go CollectStatus(ctx, nodeItem)
		}
		NodeListMutex.Unlock()
	}
	return nil
}

func GetTmpCephConfFile(monIpAddr string) (string, error) {
	tmpPath := os.TempDir() + uuid.New().String() + ".conf"
	content := `[global]
mon_host = `+ monIpAddr +`
`
	err := ioutil.WriteFile(tmpPath, []byte(content), os.ModePerm)
	if err != nil {
		logrus.Errorf("ioutil.WriteFile err:%v", err)
		return "", err
	}
	return tmpPath, nil
}

type CephConn struct {

}

var (
	cephConnCache = map[string]*rados.Conn{}
)

func CephConnInit(ipAddr string) error {
	cephConfPath, err := GetTmpCephConfFile(ipAddr)
	if err != nil {
		return err
	}
	defer func() {_ = os.Remove(cephConfPath)}()

	conn, err := rados.NewConn()
	if err != nil {
		return err
	}

	err = conn.ReadConfigFile(cephConfPath)
	if err != nil {
		return err
	}

	err = conn.Connect()
	if err != nil {
		return err
	}

	cephConnCache[ipAddr] = conn
	return nil
}

func CollectStatus(ctx context.Context, node models.Node) error {
	var (
		conn *rados.Conn
		isExist bool
		err error
	)
	if conn, isExist = cephConnCache[node.Address]; !isExist {
		err = CephConnInit(node.Address)
		if err != nil {
			return err
		}
		conn = cephConnCache[node.Address]
	}

	command, err := json.Marshal(
		map[string]string{"prefix": "status", "format": "json"})
	if err != nil {
		return err
	}
	buf, _, err := conn.MonCommand(command)
	if err != nil {
		return err
	}

	type CephStatusResp struct {
		PGMap struct {
			ReadBytesSec          float64 `json:"read_bytes_sec"`
			ReadOpPerSec          float64 `json:"read_op_per_sec"`
			RecoveringBytesPerSec float64 `json:"recovering_bytes_per_sec"`
			WriteBytesSec         float64 `json:"write_bytes_sec"`
			WriteOpPerSec         float64 `json:"write_op_per_sec"`
		} `json:"pgmap"`
	}

	var resp CephStatusResp
	err = json.Unmarshal(buf, &resp)
	if err != nil {
		return err
	}

	body := models.CephStatusBody{}
	body.ReadBytesSec = resp.PGMap.ReadBytesSec
	body.ReadOpPerSec = resp.PGMap.ReadOpPerSec
	body.RecoveringBytesPerSec = resp.PGMap.RecoveringBytesPerSec
	body.WriteBytesSec = resp.PGMap.WriteBytesSec
	body.WriteOpPerSec = resp.PGMap.WriteOpPerSec
	_, err = CreateCephStatus(ctx, node.NodeID, body)
	if err != nil {
		return err
	}

	return nil
}
