package pmysql

import (
	"time"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"go.uber.org/zap"

	"github.com/pan-jf/go-utils/plog"
)

type DBConfig struct {
	Source          string   `json:"source"`
	Replicas        string   `json:"replicas"`
	ReplicaTables   []string `json:"replicaTables"`
	MaxIdleConnNum  int      `json:"maxIdleConnNum"`
	MaxOpenConnNum  int      `json:"maxOpenConnNUm"`
	ConnMaxLifeTime int      `json:"connMaxLifeTime"`
}

// InitMysql 初始化mysql
func InitMysql(dbConfig *DBConfig) *gorm.DB {
	db, err := sql.Open("mysql", dbConfig.Source)
	if err != nil {
		plog.Error("InitMysql sql open", zap.Any("dbConfig", dbConfig), zap.Error(err))
		return nil
	}

	db.SetMaxIdleConns(dbConfig.MaxIdleConnNum)
	db.SetMaxOpenConns(dbConfig.MaxOpenConnNum)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifeTime) * time.Minute)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		plog.Error("InitMysql gorm open", zap.Any("dbConfig", dbConfig), zap.Error(err))
		return nil
	}

	if len(dbConfig.Replicas) > 0 && len(dbConfig.ReplicaTables) > 0 {

		err = gormDB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open(dbConfig.Replicas)},
		}, dbConfig.ReplicaTables))

		if err != nil {
			plog.Error("InitMysql gorm Use", zap.Any("dbConfig", dbConfig), zap.Error(err))
			return nil
		}
	}

	return gormDB
}
