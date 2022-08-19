package pmysql

import (
	"fmt"
	"testing"
)

func TestInitMysql(t *testing.T) {
	db := InitMysql(&DBConfig{
		Source:          "user:password@tcp(hostname:port)/dbName?charset=utf8&parseTime=True&loc=Local&timeout=1000ms",
		MaxIdleConnNum:  10,
		MaxOpenConnNum:  10,
		ConnMaxLifeTime: 60,
	})

	fmt.Println(db)
}
