package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

type AdminController struct {
	beego.Controller
}

// /admin/publish
func (this *AdminController) Publish()  {
	Log.Info("admin publish")
	this.ServeJson()
}


