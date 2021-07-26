package utils

import (
	"github.com/liucxer/srv-ceph-status/pkg/tools"
	"sync"
	"time"
)

func IDMapToIDArray(sources map[tools.SFID]interface{}) (IDs tools.SFIDs) {
	for k := range sources {
		IDs = append(IDs, k)
	}
	return
}

// 求两个sync.map的差集keys
func DifferenceMap(totalMap, subMap *sync.Map) (differenceKeys []uint64) {
	totalMap.Range(func(key, value interface{}) bool {
		if _, ok := subMap.Load(key); !ok {
			differenceKeys = append(differenceKeys, key.(uint64))
		}
		return true
	})
	return
}

func NowDistanceTodayBegin() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return currentTime.Unix() - startTime
}
