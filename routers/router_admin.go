package routers

import (
	"github.com/ro4tub/iedbweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/sync2es", &controllers.AdminController{}, "get:Sync2ES")
}