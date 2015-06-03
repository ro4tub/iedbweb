package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/ro4tub/iedbweb/models"
)

type  Context struct{	
	UserId	int64
	NickName string
	Permission int // 0 普通 1 管理员
}

// 是否登录
func (c *Context) IsLogin() bool {
	return (c.UserId != 0)
}

// 是否是管理员
func (c *Context) IsAdmin() bool {
	return (c.Permission == 1)
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
	results, err := models.GetLatest50Items()
	if err != nil {
		this.Abort("500")
	}
	
	gameresults := make([]Game, len(results))
	var game Game
	for k, e := range results {
		if err := json.Unmarshal([]byte(e.Data), &game); err != nil {
			this.Abort("500")
		}
		game.Logo = fmt.Sprintf("/static/upload/%d.png", e.Id)
		gameresults[k] = game
	}
	this.Data["SearchResult"] = gameresults
}
