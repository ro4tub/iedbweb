package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

type GameController struct {
	beego.Controller
}




func (this *GameController) PreCreateGame()  {
	Log.Debug("PreCreateGame")
    this.Layout = "layout.tpl"
    this.TplNames = "gamecreate.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "gamecreate_css.tpl"
    this.LayoutSections["JavaScript"] = "gamecreate_javascript.tpl"
	context := this.GetSession("CONTEXT")
	if context == nil {
		this.Data["Context"] = &Context{}
	} else {
		this.Data["Context"] = context.(*Context)
	}
	pagecontext := &PageContext{TitleName: "IEDB - 创建游戏"}
	pagecontext.Css = []string{"static/css/bootstrap-tagsinput.css", "static/css/fileinput.min.css"}
	
	this.Data["PageContext"] = pagecontext
}

func (this *GameController) CreateGame()  {
	Log.Debug("CreateGame")
	data := this.Data["Form"]
	if data == nil {
		Log.Warn("data == nil")
	} else {
		Log.Warn("data != nil")		
	}
}

func (this *GameController) GetGame()  {
	Log.Debug("GetGame")
}

func (this *GameController) EditGame()  {
	Log.Debug("EditGame")
}

func (this *GameController) SaveGame()  {
	Log.Debug("SaveGame")
}

func (this *GameController) GetEditHistory()  {
	Log.Debug("GetEditHistory")
}