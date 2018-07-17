package mysql

import (
	"common-go/log"
	"github.com/go-xorm/xorm"
)

var baseDao *BaseDao

type BaseDao struct {
	engine *xorm.Engine
}

func GetBaseDaoInstance() *BaseDao {
	if baseDao==nil {
		baseDao=&BaseDao{engine:Engine}
	}
	return baseDao
}

func (this *BaseDao)Save(bean interface{}) bool {
	affected,err:=this.engine.Insert(bean)
	if err == nil{
		if affected > 0 {
			return true
		}
	}
	return false
}

func (this *BaseDao)BatchSave(bean *interface{}) bool {
	affected,err:=this.engine.Insert(bean)
	log.Logger.Error(err)
	if err == nil{
		if affected > 0 {
			return true
		}
	}
	return false
}



func (this *BaseDao)SaveWithSession(session *xorm.Session,beans...interface{}) bool {
	affected,err:=session.Insert(beans)
	if err == nil{
		if affected > 0 {
			return true
		}
	}

	return false
}

func (this *BaseDao) Update(beans interface{}) bool {
	affected,err:=this.engine.Update(beans)
	if err == nil{
		if affected > 0 {
			return true
		}
	}

	return false
}

func (this *BaseDao) DeleteById(id int64,bean interface{}) bool {
	affected,err:=this.engine.ID(id).Delete(&bean)
	if err == nil{
		if affected > 0 {
			return true
		}
	}
	return false
}

func (this *BaseDao)GetEngine() *xorm.Engine  {
	return this.engine
}

func (this *BaseDao)GetSession() *xorm.Session  {
	return this.engine.NewSession()
}

func (this *BaseDao)SessionCommit(session *xorm.Session) {
	session.Commit()
}



