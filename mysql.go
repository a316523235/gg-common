package gg_common

import (
	"fmt"
	"github.com/a316523235/gg-common/conf"
	"github.com/go-xorm/xorm"
)

var mysqlEn = map[string]*xorm.Engine{}

func GetGinStudy()  *xorm.Engine {
	connKey := "gin_study"
	return getMysql(connKey)
}

func getMysql(connKey string) *xorm.Engine {
	if mysqlEn[connKey] == nil {
		conn := conf.Conn["mysql"][connKey];
		user, pwd, host, port, dbname := conn["user"], conn["pwd"], conn["host"], conn["port"], conn["dbname"]
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, dbname)
		en, err := xorm.NewEngine("mysql", dataSource)
		if err != nil {
			fmt.Println("mysql error : "+err.Error())
			panic(err.Error())
		}
		mysqlEn[connKey] = en
	}

	return mysqlEn[connKey]
}