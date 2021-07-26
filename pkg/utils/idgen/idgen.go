package idgen

import (
	"time"

	"github.com/go-courier/snowflakeid"
	"github.com/go-courier/snowflakeid/workeridutil"
)

var startTime, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
var sff = snowflakeid.NewSnowflakeFactory(16, 8, 5, startTime)

func MustNewIDGen() IDGen {
	idgen, err := NewIDGen()
	if err != nil {
		panic(err)
	}
	return idgen
}

func NewIDGen() (IDGen, error) {
	return sff.NewSnowflake(workeridutil.WorkerIDFromIP(ResolveExposedIP()))
}

type IDGen interface {
	ID() (uint64, error)
}
