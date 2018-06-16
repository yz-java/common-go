package mysql

import (
	"github.com/go-xorm/xorm"
	"log/syslog"
	"common-go/log"
	"encoding/json"
	"io/ioutil"
	"time"
)

var DbConfig DBConfig

var engine *xorm.Engine

type DBConfig struct {
	UserName string
	Password string
	DbName string
}
//加载mysql配置文件
func LoadDBConfig(uri string) {
	configData, e := ioutil.ReadFile(uri)

	if e != nil {
		panic(uri + " file read fail:" + e.Error())
	}
	dbConfig:=DBConfig{}
	err := json.Unmarshal(configData, &dbConfig)
	if err != nil {
		log.Logger.Error("load json redis config fail:", err)
	}
	DbConfig = dbConfig
}

func CreateEngine() error  {
	if DbConfig == nil {
		panic("please load db config")
	}
	var err error
	engine,err=xorm.NewEngine("mysql", DbConfig.UserName+":"+DbConfig.Password+"@/"+DbConfig.DbName+"?charset=utf8")
	if err != nil {
		return err
	}
	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.Logger.Error("Fail to create xorm system logger: %v\n", err)
		return err
	}
	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	engine.SetLogger(logger)
	return nil
}
//自动重连
func StartAutoConnect() {
	go func() {
		for{
			time.Sleep(10)
			err := engine.Ping()
			if err != nil {
				CreateEngine()
			}
		}
	}()

}
