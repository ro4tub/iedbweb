package bootstrap

import (
	"strings"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

func join(in []string)(out string){
    out = strings.Join(in, ",")
    return
}

func split(in string)(out []string){
    out = strings.Split(in, ",")
    return
}




func init() {
	Log.Info("bootstrap starting ...")
	// 初始化数据库
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dburl"), 30)
	// 增加模板函数
	beego.AddFuncMap("join", join)
	beego.AddFuncMap("split", split)
	Log.Info("bootstrap end ...")
}
