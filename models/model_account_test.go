package models

import (
	"github.com/astaxie/beego/orm"
	"testing"
)

func clearDB(tablename string, t *testing.T) {
	myorm := orm.NewOrm()
	// FIXME 为何用占位符不行呢
	//if _, err := myorm.Raw("TRUNCATE ?", tablename).Exec(); err != nil {
	if _, err := myorm.Raw("TRUNCATE " + tablename).Exec(); err != nil {
		t.Fatal("clearDB %s failed: %v", tablename, err)
		return
	}
}

func initDB(t *testing.T) {
	orm.RegisterDataBase("default", "mysql", "root:ib3@tcp(10.211.55.3:3306)/iedb_test?charset=utf8", 30)
	orm.Debug = true
	// orm.DebugLog = orm.NewLog(w)
	// 初始化数据库
	clearDB("account", t)
	clearDB("invite_code", t)
}

func TestAccount(t *testing.T) {
	initDB(t)
	var account *Account
	account = CreateAccount("kaifhuang", "ro4tub@gmail.com", "ttmds2014")
	if account == nil {
		t.Error("CreateAccount failed")
		return
	}
	account = CreateAccount("kaifhuang", "ro4tub@gmail.com", "ttmds2014")
	if account != nil {
		t.Error("oops CreateAccount with duplicate name and email")
		return
	}
	account = GetAccountByNameOrEmail("kaifhuang", "")
	if account == nil {
		t.Error("not get account")
	}
	
	account = GetAccountByNameOrEmail("", "ro4tub@gmail.com")
	if account == nil {
		t.Error("not get account")
	}
	
	account = GetAccountByNameOrEmail("ooxx", "ooxx@gmail.com")
	if account != nil {
		t.Error("oops, get an account")
	}
}


func TestInviteCode(t *testing.T) {
	initDB(t)
	var invitecode *InviteCode
	invitecode = GetInviteCode("ooxx")
	if invitecode != nil {
		t.Error("oops, get an invite code: ooxx")
		return
	}
	if err := CreateInviteCode(10001, "ooxx"); err != nil {
		t.Error("CreateInviteCode failed")
		return
	}
	
	invitecode = GetInviteCode("ooxx")
	if invitecode == nil {
		t.Error("not found invite code: ooxx")
		return
	}
	
	if invitecode.Code != "ooxx" {
		t.Error("oops expected code: ooxx, actual code: %s", invitecode.Code)
		return
	}
	
	err := UpdateInviteCode("ooxx", 10002)
	if err != nil {
		t.Error("UpdateInviteCode failed: %v", err)
		return
	}
	invitecode = GetInviteCode("ooxx")
	if invitecode == nil {
		t.Error("not found invite code: ooxx")
		return
	}
	if invitecode.UsedBy != 10002 {
		t.Error("UpdateInviteCode failed: invitecode.UsedBy != 10002")
		return
	}
}