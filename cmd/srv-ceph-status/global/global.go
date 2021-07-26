package global

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/sqlx/v2/migration"
	"github.com/liucxer/confmiddleware"
	"github.com/liucxer/confmiddleware/appinfo"
	"github.com/liucxer/confmiddleware/confhttp"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
	"github.com/sirupsen/logrus"
)

var (
	server = &confhttp.Server{}

	postgres = &confmiddleware.Postgres{
		Database: models.DBCephStatus,
		PoolSize: 50, // 连接池
	}

	App = appinfo.New(
		appinfo.WithName("srv-ceph-status"),
		appinfo.WithMainRoot(".."),
	)
	idGen = idgen.MustNewIDGen()
)

func init() {

	config := &struct {
		Server   *confhttp.Server
		Postgres *confmiddleware.Postgres
	}{
		Postgres: postgres,
		Server:   server,
	}

	App.ConfP(config)
}

// 通过上下文注入依赖
var WithContext = confhttp.WithContextCompose(
	db.WithDBExecutor(postgres),
	idgen.WithIDGen(idGen),
)

func Server() courier.Transport {
	return server.WithContextInjector(WithContext)
}

func Migrate() {
	logrus.SetReportCaller(false)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := migration.Migrate(postgres, nil); err != nil {
		panic(err)
	}
}
