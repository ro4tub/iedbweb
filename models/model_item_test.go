package models

import (
	"testing"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
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
	clearDB("item", t)
	clearDB("item_edit", t)
}

type Game struct {
	Name string 
	Genre string 
	Platform string 
	Logo string 
	Tags string 
	SimpleDesc string 
	Detail string
}

func TestItem(t *testing.T) {
	initDB(t)
	// 新建
	game := Game{Name:"天天魔斗士", Genre:"ARPG", Platform:"Android", Logo:"/static/upload/1.png", Tags:"Unity3D,冒险", SimpleDesc:"腾讯天美工作室出品的好玩游戏", Detail:"这是一款**好玩**的游戏"}
	data, err := json.Marshal(game)
	if err != nil {
		t.Error("json.Marshal Failed: %v", err)
		return
	}
	itemedit := &ItemEdit{AccountId:1, ItemName: game.Name, Data:string(data), Comment:"首次递交", Status:0, AddTime:time.Now(), ModifyTime:time.Now()}
	err = CreateItemEdit(itemedit)
	if err != nil {
		t.Error("CreateItemEdit failed: %v", err)
		return
	}
	
	// 审核通过
	item := &Item{AuthorId:itemedit.AccountId, Name: itemedit.ItemName, Data:itemedit.Data, Version:1, AddTime:time.Now(), ModifyTime:time.Now()}
	err = CreateItem(item)
	if err != nil {
		t.Error("CreateItem failed: %v", err)
		return
	}

	// 审核通过
	itemedit.ItemId = item.Id
	itemedit.Status = 1
	itemedit.Version = 1
	itemedit.ReviewId = 2
	itemedit.ReviewComment = "通过"
	itemedit.ReviewTime = time.Now()
	itemedit.ModifyTime = time.Now()
	err = UpdateItemEdit(itemedit)
	if err != nil {
		t.Error("UpdateItemEdit failed: %v", err)
		return
	}
	
	items, err := GetLatest10Items()
	if err != nil {
		t.Error("GetLatest10Items failed: %v", err)
	}
	if len(items) != 1 {
		t.Error("GetLatest10Items failed")
	}
}