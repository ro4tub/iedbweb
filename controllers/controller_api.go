package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

type ApiController struct {
	beego.Controller
}



func (this *ApiController) CreateItem()  {
	Log.Info("CreateItem")
	// item := models.Item{}
	//     if err := json.Unmarshal(this.Ctx.Input.RequestBody, &item); err != nil {
	// 	Log.Error("json.Unmarshal failed: %v", err)
	// 	this.Data["json"] = err
	//      } else {
	// 	 Log.Info("create item info: %s", item.ToString())
	// 	 // models.CreateItem(g)
	// 	 this.Data["json"] = g
	//      }
	this.ServeJson()
}

func (this *ApiController) DeleteItem()  {
	this.ServeJson()
}

func (this *ApiController) UpdateItem()  {
	Log.Info("UpdateItem")
	// item := models.Item{}
	//     if err := json.Unmarshal(this.Ctx.Input.RequestBody, &item); err != nil {
	// 	Log.Error("json.Unmarshal failed: %v", err)
	// 	this.Data["json"] = err
	//      } else {
	// 	 Log.Info("save item info: %s", item.ToString())
	// 	 // models.SaveItem(g)
	// 	 this.Data["json"] = g
	//      }
	this.ServeJson()
}

func (this *ApiController) GetItem()  {
	this.ServeJson()
}


