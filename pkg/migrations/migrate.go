package migrations

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" //import mysql driver
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"pgxs.io/chassis"
	xlog "pgxs.io/chassis/log"

	"pgxs.io/chassis/config"

	"pgxs.io/gopkg/pkg/iconst"
)

const (
	//EnvPgTestDataFile import test data sql file
	EnvPgTestDataFile = "PG_TEST_DATA_FILE"
	//AppRunEnvLocal app run env "local"
	AppRunEnvLocal = "local"
	//AppRunEnvTest app run env "test"
	AppRunEnvTest = "test"
	//AppRunEnvProd app run env "prod"
	AppRunEnvProd = "prod"
)

//Migrate start migrations
func Migrate() error {
	//数据库版本管理
	return runMigrations(AssetNames(), func(name string) ([]byte, error) {
		return Asset(name)
	}, config.Database().DSN)
}

//Run new bindataInstance and UP
func runMigrations(assetNames []string, afn bindata.AssetFunc, dbURL string) error {

	log := xlog.New().Service(iconst.AppName).Category("DB-Migrations")

	// wrap assets in Resource
	s := bindata.Resource(assetNames, afn)

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build instance error")
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, "mysql://"+dbURL)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build bindata instance error")
	}

	//IF ENV NOT PROD IMPORT TEST DATA
	if !isProd() {
		if err := m.Down(); err != nil {
			log.Error(err)
		}
	}

	upErr := m.Up() // run your migrations
	if upErr != nil && upErr != migrate.ErrNoChange {
		log.Errorf("数据库迁移异常,错误:%s", upErr.Error())
		return errors.New("DB migrations UP error " + upErr.Error())
	}

	//IF ENV NOT PROD IMPORT TEST DATA
	if !isProd() {
		fileName := os.Getenv(EnvPgTestDataFile)
		log.Info(fileName)
		if fileName != "" {
			if file, err := os.Open(fileName); err == nil {
				// count := 0
				if data, err := ioutil.ReadAll(file); err == nil {
					chassis.DB().Exec(string(data))
				}
			} else {
				log.Error(err)
			}
		} else {
			log.Error("导入数据文件未配置")
		}
	}
	return nil
}

func isProd() bool {
	return !(config.App().Env == "test" || config.App().Env == "local")
}
