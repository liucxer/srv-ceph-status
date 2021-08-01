package global

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/sqlx/v2/migration"
	"github.com/liucxer/confmiddleware/confhttp"
	"github.com/liucxer/confmiddleware/conflogger"
	"github.com/liucxer/confmiddleware/confpostgres"
	"github.com/liucxer/confmiddleware/scaffold/pkg/appinfo"
	"github.com/liucxer/srv-ceph-status/pkg/models"
	"github.com/liucxer/srv-ceph-status/pkg/utils/db"
	"github.com/liucxer/srv-ceph-status/pkg/utils/idgen"
	"github.com/sirupsen/logrus"
)

var (
	server = &confhttp.Server{}

	postgres = &confpostgres.Postgres{
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
		Log      *conflogger.Log
		Server   *confhttp.Server
		Postgres *confpostgres.Postgres
	}{
		Postgres: postgres,
		Server:   server,
	}

	App.ConfP(config)
	confhttp.RegisterCheckerFromStruct(config)
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
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := migration.Migrate(postgres, nil); err != nil {
		panic(err)
	}
}
