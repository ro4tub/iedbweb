package models
import (
	"time"
	"errors"
	// "database/sql"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
	. "github.com/ro4tub/gamedb/util"
)

// 游戏记录
type Item struct {
	Id	int64
	Name string
	AuthorId int64 // 创建者的账号id
	Data string // 存放Json数据
	Version int
	AddTime time.Time
	ModifyTime time.Time
}

type ItemEdit struct {
	Id int64
	ItemId int64
	ItemName string
	AccountId int64
	Data string
	Comment string
	Status byte // 0 审核中 1 通过审核 2 被拒绝
	ReviewId int64
	ReviewComment string
	ReviewTime time.Time
	Version int // Status == 1 的时候才有意义
	AddTime time.Time
	ModifyTime time.Time
}

func init() {
	orm.RegisterModel(new(Item))
	orm.RegisterModel(new(ItemEdit))
}

// 创建编辑流水
func CreateItemEdit(itemedit *ItemEdit) error {
	Log.Debug("CreateItemEdit: %d %d", itemedit.ItemId, itemedit.AccountId)
	myorm := orm.NewOrm()
	itemeditid, err := myorm.Insert(itemedit)
	if err != nil {
		Log.Error("CreateItemEdit failed: %v", err)
		return err
	}
	itemedit.Id = itemeditid
	return nil
}

func UpdateItemEdit(itemedit *ItemEdit) error {
	Log.Debug("UpdateItemEdit:%d %d", itemedit.ItemId, itemedit.AccountId)
	myorm := orm.NewOrm()
	affectedrows, err := myorm.Update(itemedit)
	if err != nil {
		Log.Error("UpdateItemEdit failed: %v", err)
		return err
	}
	if affectedrows != 1 {
		Log.Error("affectedrows != 1")
		return errors.New("affectedrows != 1")
	}
	return nil
}

func CreateItem(item *Item) error {
	Log.Debug("CreateItem: %d", item.AuthorId)
	myorm := orm.NewOrm()
	itemid, err := myorm.Insert(item)
	if err != nil {
		Log.Error("CreateItem failed: %v", err)
		return err
	}
	item.Id = itemid
	return nil
}

func UpdateItem(item *Item) error {
	Log.Debug("UpdateItem: %d", item.AuthorId)
	myorm := orm.NewOrm()
	affectedrows, err := myorm.Update(item, "Data", "Version", "ModifyTime")
	if err != nil {
		Log.Error("UpdateItem failed: %v", err)
		return err
	}
	if affectedrows != 1 {
		Log.Error("affectedrows != 1")
		return errors.New("affectedrows != 1")
	}
	return nil
}

func GetItemById(id int64) *Item {
	Log.Debug("GetItemById: %d", id)
	myorm := orm.NewOrm()
	var item Item
	item.Id = id
	err := myorm.Read(&item)
	if err != nil {
		Log.Error("orm.Read failed: %v", err)
		return nil
	}
	return &item
	
}

func GetItemByName(name string) *Item {
	Log.Debug("GetItemByName: %s", name)
	myorm := orm.NewOrm()
	var item Item
	if err := myorm. Raw("select * from item where name = ?", name).QueryRow(&item); err != nil {
		if err == orm.ErrNoRows {
			Log.Warn("no matched row")
		} else {
			Log.Error("select from item failed", err)
		}
		return nil
	}
	return &item
}

// 获得最新添加的50个
func GetLatest50Items() ([]Item, error) {
	Log.Debug("GetLatest10Items")
	var items []Item
	myorm := orm.NewOrm()
	num, err := myorm.Raw("select * from item order by add_time limit 50").QueryRows(&items)
	if err != nil {
		if err == orm.ErrNoRows {
			Log.Warn("no matched row")
		} else {
			Log.Error("select from item failed", err)
		}
		return nil, err
	}
	Log.Debug("got %d items", num)
	return items, nil
}

// 获得所有的Item， 有性能问题，使用请注意
// TODO 分页查询
func GetItems() ([]Item, error) {
	Log.Debug("GetItems")
	var items []Item
	myorm := orm.NewOrm()
	num, err := myorm.Raw("select * from item").QueryRows(&items)
	if err != nil {
		if err == orm.ErrNoRows {
			Log.Warn("no matched row")
		} else {
			Log.Error("select from item failed", err)
		}
		return nil, err
	}
	Log.Debug("got %d items", num)
	return items, nil
	
}