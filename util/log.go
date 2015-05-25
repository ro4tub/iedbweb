package util

import (
    "github.com/astaxie/beego/logs"
)

var (
	Log *logs.BeeLogger
)

func init()  {
	Log = logs.NewLogger(10000)
	Log.SetLogger("console", "")
	Log.SetLogger("file", `{"filename":"test.log"}`)
}