package controllers

import (
	"github.com/astaxie/beego"
)

type  Context struct{	
	UserId	string
}

func (c *Context) IsLogin() bool {
	return (c.UserId != "")
}


type PageContext struct {
	TitleName string
	Css	[]string
	Javascript []string
	JavascriptCode string
}

type Result struct {
	Ret	string
	Desc string
}

type MainController struct {
	beego.Controller
}

//TODO 使用filter
func (this *MainController) Get() {
    this.Layout = "layout.tpl"
    this.TplNames = "index.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "default_css.tpl"
    this.LayoutSections["JavaScript"] = "default_javascript.tpl"
	context := this.GetSession("CONTEXT")
	if context == nil {
		this.Data["Context"] = &Context{}
	} else {
		this.Data["Context"] = context.(*Context)
	}
	this.Data["PageContext"] = &PageContext{TitleName: "IEDB - 首页"}
}
