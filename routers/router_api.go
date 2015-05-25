package routers

import (
	"github.com/ro4tub/iedbweb/controllers"
	"github.com/astaxie/beego"
)

// 后台API的路由，都返回json数据
func init() {
	// 创建游戏条目
	beego.Router("/api/game", &controllers.ApiController{}, "post:CreateItem")
	// 删除游戏条目
	beego.Router("/api/game", &controllers.ApiController{}, "delete:DeleteItem")
	// 修改游戏条目
	beego.Router("/api/game", &controllers.ApiController{}, "put:UpdateItem")
	// 获得游戏条目
	beego.Router("api/game/:name", &controllers.ApiController{}, "get:GetItem")
}