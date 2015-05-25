package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Email string 	`form:"email"`
	NickName string	`form:"nickname"`
	InviteCode string `form:"invitecode"`
	Password string `form:"password"`
}

func (this *UserController) CreateUser()  {
	Log.Debug("CreateUser")
    u := User{}
	if err := this.ParseForm(&u); err != nil {
       Log.Error("ParseForm failed: %v", err)
	   return
    }
	Log.Debug("email:%s, nickname:%s, invitecode:%s, password:%s", u.Email, u.NickName, u.InviteCode, u.Password)
	context := &Context{UserId:"0123456789"}
	this.SetSession("CONTEXT", context)
	// this.Data["json"] = &Result{Ret:"ok"}
	// this.ServeJson()
	this.Redirect("/", 302)
}

func (this *UserController) Login()  {
	Log.Debug("Login")
	email := this.GetString("email")
	password := this.GetString("password")
	Log.Debug("email=%s, password=%s", email, password)
	context := &Context{UserId:"0123456789"}
	this.SetSession("CONTEXT", context)
	this.Redirect("/", 302)
}

func (this *UserController) Logout()  {
	Log.Debug("Logout")
	this.DelSession("CONTEXT")	
	this.Redirect("/", 302)
}

func (this *UserController) GetUser()  {
	Log.Debug("GetUser")
}

func (this *UserController) GetMyInfo()  {
	Log.Debug("GetMyInfo")
}

func (this *UserController) GetMyAudit()  {
	Log.Debug("GetMyAudit")
}