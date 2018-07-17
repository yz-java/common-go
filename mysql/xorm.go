package mysql

import (
	"github.com/go-xorm/xorm"
	//_"github.com/go-sql-driver/mysql"
	"common-go/log"
	"encoding/json"
	"io/ioutil"
	"time"
	"log/syslog"
	"fmt"
)

var DbConfig *DBConfig = nil

var Engine *xorm.Engine = nil

type DBConfig struct {
	UserName string
	Password string
	DbName string
	Address string
	Port string
	MaxIdel int
	MaxOpen int
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
	log.Logger.Info("database config : %v",DbConfig)
}

func CreateEngine() error  {
	if DbConfig == nil {
		panic("please load db config")
	}
	var err error
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DbConfig.UserName, DbConfig.Password, DbConfig.Address, DbConfig.DbName)
	Engine,err=xorm.NewEngine("mysql",mysqlUrl)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.Logger.Error("Fail to create xorm system logger: %v\n", err)
		return err
	}
	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	Engine.SetLogger(logger)
	Engine.SetMaxIdleConns(DbConfig.MaxIdel)
	Engine.SetMaxOpenConns(DbConfig.MaxOpen)
	return nil
}
//自动重连
func StartAutoConnect() {
	go func() {
		for{
			time.Sleep(10*time.Second)
			err := Engine.Ping()
			if err != nil {
				CreateEngine()
			}
		}
	}()

}
