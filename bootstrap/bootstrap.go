package bootstrap

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

func init() {
	Log.Info("bootstrap starting ...")
	// 初始化数据库
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dburl"), 30)
	Log.Info("bootstrap end ...")
}
