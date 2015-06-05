package controllers

import (
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"github.com/ro4tub/iedbweb/models"
	. "github.com/ro4tub/gamedb/util"
)

type GameController struct {
	beego.Controller
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

// 搜索结果
type GameSearchResult struct {
	Name string 
	Genre string
	Platform string
	Logo string
	Tags []string
	SimpleDesc string
}

var (
	platformStr map[int]string
	genreStr map[int]string
)

func (g *Game) PlatformStr() string {
	num, _ := strconv.Atoi(g.Platform)
	return platformStr[num]
}

func (g *Game) GenreStr() string {
	num, _ := strconv.Atoi(g.Genre)
	return genreStr[num]
}


func init() {
	platformStr = map[int]string {
		1: "iOS",
		2: "Android",
		3: "Xbox 360",
		4: "Xbox One",
		5: "PS3",
		6: "PS4",
		7: "PC",
		8: "浏览器",
		99: "其它",
	}
	genreStr = map[int]string {
		1: "动作",
		2: "动作冒险",
		3: "冒险解密",
		4: "休闲益智",
		5: "角色扮演",
		6: "竞速",
		7: "模拟",
		8: "体育竞技",
		9: "策略",
		99: "其它",
	}
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
	pagecontext := &PageContext{TitleName: "IEDb - 创建游戏"}
	this.Data["PageContext"] = pagecontext
}


func (this *GameController) CreateGame()  {
	Log.Debug("CreateGame")
	contextdata := this.GetSession("CONTEXT")
	if contextdata == nil {
		this.Abort("401")
	}
	context , ok := contextdata.(*Context)
	if !ok || context.IsLogin() == false {
		this.Abort("401")
	}
	this.Data["Context"] = context
	
    g := Game{}
	if err := this.ParseForm(&g); err != nil {
       Log.Error("ParseForm failed: %v", err)
	   return
    }
	// trim
	g.SimpleDesc = strings.TrimSpace(g.SimpleDesc)
	g.Detail = strings.TrimSpace(g.Detail)
	Log.Debug("game name:%s, genre:%s, platform:%s, tags:%s, simpledesc:%s, detail:%s", g.Name, g.Genre, g.Platform, g.Tags, g.SimpleDesc, g.Detail)
	var data []byte
	var err error
	data, err = json.Marshal(g)
	if err != nil {
		Log.Error("json.Marshal failed: %v", err)
		this.Abort("500")
	}
	var item models.Item
	if context.IsAdmin() == true {
		item.Name = g.Name
		item.AuthorId = context.UserId
		item.Version = item.Version + 1
		item.AddTime = time.Now()
		item.ModifyTime = time.Now()
		item.Data = string(data)

		if err := models.CreateItem(&item); err != nil {
			Log.Error("CreateItem failed: %v", err)
			this.Abort("500")
		}
		// 更新索引
		gamedoc := models.GameDocument{Id: item.Id, Name: g.Name, Genre: g.Genre, Platform: g.Platform, Tags: strings.Split(g.Tags, ","), SimpleDesc: g.SimpleDesc, Detail: g.Detail}
		if err := models.CreateSearchIndex(gamedoc); err != nil {
			Log.Error("CreateSearchIndex failed: %v", err)
		}
	}
	itemedit := &models.ItemEdit{AccountId:context.UserId, ItemName: g.Name, Data:string(data), Comment:"新建", Status:0, AddTime:time.Now(), ModifyTime:time.Now()}
	if context.IsAdmin() == true {
		itemedit.ItemId = item.Id
		itemedit.Status = 1
		itemedit.Version = 1
		itemedit.ReviewId = context.UserId
		itemedit.ReviewComment = "自动通过"
		itemedit.ReviewTime = time.Now()
	}
	err = models.CreateItemEdit(itemedit)
	if err != nil {
		Log.Error("CreateItemEdit failed: %v", err)
		this.Abort("500")
	}

	// 保存gamelogo
	// FIM 改用${名字}.png，这样可以先写文件，然后写数据库
	_, _, err = this.GetFile("gamelogo")
	if err != nil {
		Log.Error("GetFile failed: %v", err)
	}
	err = this.SaveToFile("gamelogo", fmt.Sprintf("static/upload/%d.png", item.Id))
	if err != nil {
		Log.Error("SaveToFile failed: %v", err)
	}
	this.Redirect("/game/"+g.Name, 302)
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
	item := models.GetItemByName(name)
	if item == nil {
		// 404
		this.Abort("404")
	}
	var game Game
	if err := json.Unmarshal([]byte(item.Data), &game); err != nil {
		this.Abort("500")
	}
	game.Logo = fmt.Sprintf("/static/upload/%d.png", item.Id)
	this.Data["Game"] = game
	
	pagecontext := &PageContext{TitleName: "IEDb - " + game.Name}
	this.Data["PageContext"] = pagecontext
}

func (this *GameController) EditGame()  {
	Log.Debug("EditGame")
	contextdata := this.GetSession("CONTEXT")
	if contextdata == nil {
		this.Abort("401")
	}
	context , ok := contextdata.(*Context)
	if !ok || context.IsLogin() == false {
		this.Abort("401")
	}
	this.Data["Context"] = context
	
	name := this.GetString(":name")
	Log.Debug("name:%s", name)
    this.Layout = "layout.tpl"
    this.TplNames = "gameedit.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "gameedit_css.tpl"
    this.LayoutSections["JavaScript"] = "gameedit_javascript.tpl"
	item := models.GetItemByName(name)
	if item == nil {
		// 404
		this.Abort("404")
	}
	var game Game
	if err := json.Unmarshal([]byte(item.Data), &game); err != nil {
		this.Abort("500")
	}
	this.Data["Game"] = game
	
	pagecontext := &PageContext{TitleName: "IEDb - 编辑 - " + game.Name}
	this.Data["PageContext"] = pagecontext
}

// 普通用户只会生成编辑流水，需要管理员审核后，才会更新到item表；
// 管理员生成编辑流水后，并直接更新到item表
func (this *GameController) SaveGame()  {
	Log.Debug("SaveGame")
	contextdata := this.GetSession("CONTEXT")
	if contextdata == nil {
		this.Abort("401")
	}
	context , ok := contextdata.(*Context)
	if !ok || context.IsLogin() == false {
		this.Abort("401")
	}
	this.Data["Context"] = context
    g := Game{}
	if err := this.ParseForm(&g); err != nil {
       Log.Error("ParseForm failed: %v", err)
	   return
    }
	// trim
	g.SimpleDesc = strings.TrimSpace(g.SimpleDesc)
	g.Detail = strings.TrimSpace(g.Detail)
	Log.Debug("game name:%s, genre:%s, platform:%s, tags:%s, simpledesc:%s, detail:%s", g.Name, g.Genre, g.Platform, g.Tags, g.SimpleDesc, g.Detail)
	
	item := models.GetItemByName(g.Name)
	if item == nil {
		// 404
		this.Abort("404")
	}
	var game Game
	if err := json.Unmarshal([]byte(item.Data), &game); err != nil {
		this.Abort("500")
	}
	var dirty bool
	if game.Name != g.Name {
		game.Name = g.Name
		dirty = true
	}
	
	if game.Genre != g.Genre {
		game.Genre = g.Genre
		dirty = true
	}
	
	if game.Platform != g.Platform {
		game.Platform = g.Platform
		dirty = true
	}
	
	if game.Tags != g.Tags {
		game.Tags = g.Tags
		dirty = true
	}
	
	if _, _, err := this.GetFile("gamelogo"); err == nil {
		// FIXME 通过fileheader.Filename读取文件扩展名
		err = this.SaveToFile("gamelogo", fmt.Sprintf("static/upload/%d.png", item.Id))
		if err != nil {
			Log.Error("SaveToFile failed: %v", err)
			this.Abort("500")
		}
	}
	
	if game.SimpleDesc != g.SimpleDesc {
		game.SimpleDesc = g.SimpleDesc
		dirty = true
	}
	
	if game.Detail != g.Detail {
		game.Detail = g.Detail
		dirty = true
	}
	
	if dirty == true {
		Log.Warn("need to be modify")
		var data []byte
		var  err error
		data, err = json.Marshal(game)
		if err != nil {
			Log.Error("json.Marshal failed: %v", err)
			this.Abort("500")
		}
		// FIXME 增加编辑comment
		itemedit := &models.ItemEdit{AccountId:context.UserId, ItemName: g.Name, Data:string(data), Comment:"修改", Status:0, AddTime:time.Now(), ModifyTime:time.Now()}
		if context.IsAdmin() == true {
			itemedit.ItemId = item.Id
			itemedit.Status = 1
			itemedit.Version = item.Version + 1
			itemedit.ReviewId = context.UserId
			itemedit.ReviewComment = "自动通过"
			itemedit.ReviewTime = time.Now()
		}
		err = models.CreateItemEdit(itemedit)
		if err != nil {
			Log.Error("CreateItemEdit failed: %v", err)
			this.Abort("500")
		}
		
		if context.IsAdmin() == true {
			item.Version = item.Version + 1
			item.ModifyTime = time.Now()
			item.Data = string(data)
			item.Name = g.Name
	
			if err = models.UpdateItem(item); err != nil {
				Log.Error("UpdateItem failed: %v", err)
				this.Abort("500")
			}
			// 更新索引
			gamedoc := models.GameDocument{Id: item.Id, Name: g.Name, Genre: g.Genre, Platform: g.Platform, Tags: strings.Split(g.Tags, ","), SimpleDesc: g.SimpleDesc, Detail: g.Detail}
			if err := models.CreateSearchIndex(gamedoc); err != nil {
				Log.Error("CreateSearchIndex failed: %v", err)
			}
		}
	}

	this.Redirect("/game/"+g.Name, 302)
}

// 搜索
func (this *GameController) Search() {
	Log.Debug("Search")
	text := this.GetString("t")
	Log.Debug("text=%s", text)
    this.Layout = "layout.tpl"
    this.TplNames = "search.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Css"] = "default_css.tpl"
    this.LayoutSections["JavaScript"] = "default_javascript.tpl"
	context := this.GetSession("CONTEXT")
	if context == nil {
		this.Data["Context"] = &Context{}
	} else {
		this.Data["Context"] = context.(*Context)
	}
	// 搜索
	results, err := models.Search(text, 0, 100)
	if err != nil {
		this.Abort("500")
	}
	gameresults := make([]GameSearchResult, len(results))
	for k, e := range results {
		gameresults[k] = GameSearchResult{Name: e.Name, Genre: e.Genre, Platform: e.Platform, Logo: fmt.Sprintf("/static/upload/%d.png", e.Id), Tags: e.Tags, SimpleDesc: e.SimpleDesc}
	}
	this.Data["SearchResult"] = gameresults
	pagecontext := &PageContext{TitleName: "IEDb - 搜索: "+text}
	this.Data["PageContext"] = pagecontext
}

func (this *GameController) GetEditHistory()  {

}