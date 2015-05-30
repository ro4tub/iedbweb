package models

import (
	"crypto/md5"
	"time"
	"errors"
	"encoding/hex"
	"database/sql"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
	. "github.com/ro4tub/gamedb/util"
)

type Account struct {
	Id int64 	// id
	Name string  // 昵称
	Email string // 电子邮箱
	Password string // 密码
	Permission int // 权限 0:普通用户  1:管理员
	AddTime time.Time
}

type InviteCode struct {
	Id  int64
	Code string // 邀请码
	AccountId int // 关联account id
	UsedBy  int // 被使用的id, 0表示没有使用
	AddTime  time.Time
	ModifyTime time.Time
}



func init() {
	orm.RegisterModel(new(Account))
	orm.RegisterModel(new(InviteCode))
}

func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}


// func  CreateAccount(name string, email string, password string, invitecode string) (int, error) {
// 	Log.Debug("CreateAccount: %s, %s, %s", name, email, invitecode)
// 	myorm := orm.NewOrm()
// 	var err error
// 	if err = myorm.Begin(); err != nil {
// 		Log.Error("transaction begin error: %v", err)
// 		return 0, err
// 	}
// 	passwordstored := GetMD5Hash(password)
// 	var ret sql.Result
// 	if  ret, err = myorm.Raw("insert into account values(NULL, ?, ?, ?, 0, now())", name, email, passwordstored).Exec(); err != nil {
// 		Log.Error("insert into account failed", err)
// 		myorm.Rollback()
// 		return 0, err
// 	}
// 	var accountid int64
// 	if accountid, err = ret.LastInsertId(); err != nil {
// 		Log.Error("get LastInsertId failed", err)
// 		myorm.Rollback()
// 		return 0, err
// 	}
// 	Log.Debug("last inser id=%d", accountid)
//
// 	if ret, err = myorm.Raw("update invite_code set used_by=? where code=? and used_by=0", accountid, invitecode).Exec(); err != nil {
// 		Log.Error("update invite_code failed", err)
// 		myorm.Rollback()
// 		return 0, err
// 	}
// 	var rowsaffected int64
// 	if rowsaffected, err = ret.RowsAffected(); err != nil {
// 		Log.Error("get RowsAffected failed", err)
// 		myorm.Rollback()
// 		return 0, err
// 	}
//
// 	if rowsaffected != 1 {
// 		Log.Error("update invite_code failed: rowsaffected=%d", rowsaffected)
// 		myorm.Rollback()
// 		return 0, errors.New("update invite_code failed")
// 	}
//
// 	if err := myorm.Commit(); err != nil {
// 		Log.Error("transaction commit error: %v", err)
// 		myorm.Rollback()
// 		return 0,err
// 	}
// 	return int(accountid),nil
//
// }

func CreateAccount(name string, email string, password string) *Account {
	Log.Debug("CreateAccount: %s, %s", name, email)
	myorm := orm.NewOrm()
	passwordstored := GetMD5Hash(password)
	var (
		ret sql.Result
		err error
	)

	if  ret, err = myorm.Raw("insert into account values(NULL, ?, ?, ?, 0, now())", name, email, passwordstored).Exec(); err != nil {
		Log.Error("insert into account failed", err)
		return nil
	}
	var accountid int64
	if accountid, err = ret.LastInsertId(); err != nil {
		Log.Error("get LastInsertId failed", err)
		return nil
	}
	return &Account{Name:name, Email:email, Id: accountid}
}



func GetAccountByNameOrEmail(name string, email string) *Account {
	Log.Debug("GetAccountByNameOrEmail: %s, %s", name, email)
	myorm := orm.NewOrm()
	var account Account
	if err := myorm.Raw("select * from account where name=? or email=?", name, email).QueryRow(&account); err != nil {
		if err == orm.ErrNoRows {
			Log.Warn("no matched row")
		} else {
			Log.Error("select from account failed", err)
		}
		return nil
	}
	return &account
}


func CheckAccountPassword(password string, account *Account) bool {
	passwordstored := GetMD5Hash(password)
	if account.Password != passwordstored {
		Log.Warn("password error: expected=%s, actual=%s", account.Password, passwordstored)
		return false
	}
	return true
}

func CreateInviteCode(accountid int64, code string) error {
	Log.Debug("CreateInviteCode: %d, %s", accountid, code)
	myorm := orm.NewOrm()
	if  _, err := myorm.Raw("insert into invite_code values(NULL, ?, ?, 0, now(), now())", code, accountid).Exec(); err != nil {
		Log.Error("insert into invite_code failed", err)
		return err
	}
	return nil
}

func UpdateInviteCode(code string, usedby int64) error {
	Log.Debug("UpdateInviteCode: %s, %d", code, usedby)
	myorm := orm.NewOrm()
	var ret sql.Result
	var err error
	if  ret, err = myorm.Raw("update invite_code set used_by=?, modify_time=now() where code=? and used_by=0", usedby, code).Exec(); err != nil {
		Log.Error("insert into invite_code failed", err)
		return err
	}
	var rowsaffected int64
	if rowsaffected, err = ret.RowsAffected(); err != nil {
		Log.Error("get RowsAffected failed", err)
		return err
	}
	if rowsaffected != 1 {
		return errors.New("update invite_code failed: rowsaffected != 1")
	}
	return nil
}

func GetInviteCode(code string) *InviteCode {
	Log.Debug("GetInviteCode: %s", code)
	myorm := orm.NewOrm()
	var invitecode InviteCode
	if err := myorm.Raw("select * from invite_code where code = ?", code).QueryRow(&invitecode); err != nil {
		if err == orm.ErrNoRows {
			Log.Warn("no matched row")
		} else {
			Log.Error("select from invite_code failed", err)
		}
		return nil
	}
	
	return &invitecode
}