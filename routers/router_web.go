package routers

import (
	"github.com/ro4tub/iedbweb/controllers"
	"github.com/astaxie/beego"
)

// web页面的路由
func init() {
    beego.Router("/", &controllers.MainController{})
	// 创建账号
	beego.Router("/signup", &controllers.UserController{}, "post:CreateUser")
	// 登录
	beego.Router("/signin", &controllers.UserController{}, "post:Login")
	// 登录
	beego.Router("/signout", &controllers.UserController{}, "get:Logout")
	// 访问他人
	beego.Router("/user/:name", &controllers.UserController{}, "get:GetUser")
	// 查看自己
	beego.Router("/user", &controllers.UserController{}, "get:GetMyInfo")
	// 创建游戏
	beego.Router("/creategame", &controllers.GameController{}, "get:PreCreateGame")
	beego.Router("/creategame", &controllers.GameController{}, "post:CreateGame")
	// 查看游戏
	beego.Router("/game/:name", &controllers.GameController{}, "get:GetGame")
	// 编辑游戏
	beego.Router("/game/:name/edit", &controllers.GameController{}, "get:EditGame")
	// 保存游戏
	beego.Router("/game", &controllers.GameController{}, "put:SaveGame")
	// 查看游戏的
	beego.Router("/history/:id", &controllers.GameController{}, "get:GetEditHistory")
	beego.Router("/audit", &controllers.UserController{}, "get:GetMyAudit")
}
