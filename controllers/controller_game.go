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
	this.Data["PageContext"] = pagecontext
}


type Game struct {
	Name string `form:"gamename"`
	Genre string `form:"gamegenre"`
	Platform string `form:"gameplatform"`
	Logo string `form:"-"`
	Tags string `form:"gametags"`
	SimpleDesc string `form:"gamesimpledesc"`
	Detail string `form:"gamedetail"`
}

func (this *GameController) CreateGame()  {
	Log.Debug("CreateGame")
    g := Game{}
	if err := this.ParseForm(&g); err != nil {
       Log.Error("ParseForm failed: %v", err)
	   return
    }
	Log.Debug("game name:%s, genre:%s, platform:%s, tags:%s, simpledesc:%s, detail:%s", g.Name, g.Genre, g.Platform, g.Tags, g.SimpleDesc, g.Detail)
	_, _, err := this.GetFile("gamelogo")
	if err != nil {
		Log.Error("GetFile failed: %v", err)
		return
	}
	err = this.SaveToFile("gamelogo", "upload/1.png" /*+ fileheader.Filename*/)
	if err != nil {
		Log.Error("SaveToFile failed: %v", err)
		return
	}
	this.Redirect("/game/天天魔斗士", 302)
}

func (this *GameController) GetGame()  {
	Log.Debug("GetGame")
	name := this.GetString(":name")
	Log.Debug("name:%s", name)
    this.Layout = "layout.tpl"
    this.TplNames = "game.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "default_css.tpl"
    this.LayoutSections["JavaScript"] = "game_javascript.tpl"
	context := this.GetSession("CONTEXT")
	if context == nil {
		this.Data["Context"] = &Context{}
	} else {
		this.Data["Context"] = context.(*Context)
	}
	game := &Game{}
	game.Name = "天天魔斗士"
	game.Genre = "RPG"
	game.Platform = "iOS"
	game.Logo = "/static/upload/1.png"
	game.Tags = "Unity3D,魔幻"
	game.SimpleDesc = "这是简单描述"
	game.Detail = "这是**详细**描述"
	this.Data["Game"] = game
	
	pagecontext := &PageContext{TitleName: "IEDB - " + game.Name}
	this.Data["PageContext"] = pagecontext

}

func (this *GameController) EditGame()  {
	Log.Debug("EditGame")
	name := this.GetString(":name")
	Log.Debug("name:%s", name)
    this.Layout = "layout.tpl"
    this.TplNames = "gameedit.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "gameedit_css.tpl"
    this.LayoutSections["JavaScript"] = "gameedit_javascript.tpl"
	context := this.GetSession("CONTEXT")
	if context == nil {
		this.Data["Context"] = &Context{}
	} else {
		this.Data["Context"] = context.(*Context)
	}
	game := &Game{}
	game.Name = "天天魔斗士"
	game.Genre = "5"
	game.Platform = "2"
	game.Logo = "/static/upload/1.png"
	game.Tags = "Unity3D,魔幻"
	game.SimpleDesc = "这是简单描述"
	game.Detail = "这是**详细**描述"
	this.Data["Game"] = game
	
	pagecontext := &PageContext{TitleName: "IEDB - 编辑 - " + game.Name}
	this.Data["PageContext"] = pagecontext
}

func (this *GameController) SaveGame()  {
	Log.Debug("SaveGame")
    g := Game{}
	if err := this.ParseForm(&g); err != nil {
       Log.Error("ParseForm failed: %v", err)
	   return
    }
	Log.Debug("game name:%s, genre:%s, platform:%s, tags:%s, simpledesc:%s, detail:%s", g.Name, g.Genre, g.Platform, g.Tags, g.SimpleDesc, g.Detail)
	// FIXME gamelogo上传
	// _, _, err := this.GetFile("gamelogo")
	// if err != nil {
	// 	Log.Error("GetFile failed: %v", err)
	// 	return
	// }
	// err = this.SaveToFile("gamelogo", "upload/1.png" /*+ fileheader.Filename*/)
	// if err != nil {
	// 	Log.Error("SaveToFile failed: %v", err)
	// 	return
	// }
	this.Redirect("/game/天天魔斗士", 302)
}

func (this *GameController) GetEditHistory()  {
	Log.Debug("GetEditHistory")
}