package controllers

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego"
	"github.com/ro4tub/iedbweb/models"
	. "github.com/ro4tub/gamedb/util"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Index() {
	Log.Debug("Index")
	contextdata := this.GetSession("CONTEXT")
	if contextdata == nil {
		this.Abort("401")
	}
	context , ok := contextdata.(*Context)
	if !ok || context.IsAdmin() == false {
		this.Abort("401")
	}
	this.Data["Context"] = context
	
    this.Layout = "layout.tpl"
    this.TplNames = "admin.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "admin_css.tpl"
    this.LayoutSections["JavaScript"] = "admin_javascript.tpl"
	
	pagecontext := &PageContext{TitleName: "IEDb - 后台管理"}
	this.Data["PageContext"] = pagecontext
}

// /admin/publish
func (this *AdminController) Publish()  {
	Log.Info("admin publish")
	this.ServeJson()
}

// 把db数据同步到es
func (this *AdminController) Sync2ES() {
	Log.Debug("Sync2ES")
	contextdata := this.GetSession("CONTEXT")
	if contextdata == nil {
		this.Abort("401")
	}
	context , ok := contextdata.(*Context)
	if !ok || context.IsAdmin() == false {
		this.Abort("401")
	}
	this.Data["Context"] = context
	
	result := &Result{Ret: "ok"}
	items, err := models.GetItems()
	if err != nil {
		result.Ret = err.Error()
		this.Data["json"] = result
		this.ServeJson()
		return
	} 
	
	var g Game
	for _, item := range items {
		if err := json.Unmarshal([]byte(item.Data), &g); err != nil {
			Log.Error("json.Unmarshal failed: %v", err)
			result.Ret = err.Error()
			this.Data["json"] = result
			this.ServeJson()
			return
		}
		// 更新索引
		gamedoc := models.GameDocument{Id: item.Id, Name: g.Name, Genre: g.Genre, Platform: g.Platform, Tags: strings.Split(g.Tags, ","), SimpleDesc: g.SimpleDesc, Detail: g.Detail}
		if err := models.CreateSearchIndex(gamedoc); err != nil {
			Log.Error("CreateSearchIndex failed: %v", err)
			result.Ret = err.Error()
			this.Data["json"] = result
			this.ServeJson()
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}


