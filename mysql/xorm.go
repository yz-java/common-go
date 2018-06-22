package mysql

import (
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
	"common-go/log"
	"encoding/json"
	"io/ioutil"
	"time"
	"os"
)

var DbConfig *DBConfig = nil

var Engine *xorm.Engine = nil

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
	DbConfig = &dbConfig
}

func CreateEngine() error  {
	if DbConfig == nil {
		panic("please load db config")
	}
	var err error
	Engine,err=xorm.NewEngine("mysql", DbConfig.UserName+":"+DbConfig.Password+"@/"+DbConfig.DbName+"?charset=utf8")
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	//if err != nil {
	//	log.Logger.Error("Fail to create xorm system logger: %v\n", err)
	//	return err
	//}
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return err
	}
	//logger := xorm.NewSimpleLogger(logWriter)
	logger := xorm.NewSimpleLogger(f)
	logger.ShowSQL(true)
	Engine.SetLogger(logger)
	return nil
}
//自动重连
func StartAutoConnect() {
	go func() {
		for{
			time.Sleep(10)
			err := Engine.Ping()
			if err != nil {
				CreateEngine()
			}
		}
	}()

}
