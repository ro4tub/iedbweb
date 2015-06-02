package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ro4tub/iedbweb/models"
	. "github.com/ro4tub/iedbweb/util"
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
	Log.Debug("email:%s, nickname:%s, invitecode:%s", u.Email, u.NickName, u.InviteCode)
	// TODO check email/nickname/invitecode的合法性
	account := models.GetAccountByNameOrEmail(u.NickName, u.Email)
	if account != nil {
		Log.Warn("already use the nickname or email")
		this.Data["json"] = &Result{Ret:"already use the nickname or email"}
		this.ServeJson()
		return
	}
	invitecode := models.GetInviteCode(u.InviteCode)
	if invitecode == nil || invitecode.UsedBy != 0 {
		Log.Warn("not found the invite code or it has been used")
		this.Data["json"] = &Result{Ret:"not found the invite code or it has been used"}
		this.ServeJson()
		return
	}
	// 创建账号
	account = models.CreateAccount(u.NickName, u.Email, u.Password)
	if account == nil {
		Log.Warn("db access error")
		this.Data["json"] = &Result{Ret:"db access error"}
		this.ServeJson()
		return
	}
	err := models.UpdateInviteCode(u.InviteCode, account.Id)
	if err != nil {
		Log.Warn("db access error")
		this.Data["json"] = &Result{Ret:"db access error"}
		this.ServeJson()
		return
	}
	context := &Context{UserId: account.Id, NickName: account.Name}
	this.SetSession("CONTEXT", context)
	// this.Redirect("/", 302)
	this.Data["json"] = &Result{Ret:"ok"}
	this.ServeJson()
}

func (this *UserController) Login()  {
	Log.Debug("Login")
	email := this.GetString("email")
	password := this.GetString("password")
	Log.Debug("email=%s, password=%s", email, password)
	account := models.GetAccountByNameOrEmail(email, email)
	if account == nil {
		Log.Warn("not found the nickname or email")
		this.Data["json"] = &Result{Ret:"not found the nickname or email"}
		this.ServeJson()
		return
	}
	if models.CheckAccountPassword(password, account) == false {
		Log.Warn("incorrect password")
		this.Data["json"] = &Result{Ret:"incorrect password"}
		this.ServeJson()
		return
	}
	context := &Context{UserId: account.Id, NickName: account.Name, Permission:account.Permission}
	this.SetSession("CONTEXT", context)
	// this.Redirect("/", 302)
	this.Data["json"] = &Result{Ret:"ok"}
	this.ServeJson()
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