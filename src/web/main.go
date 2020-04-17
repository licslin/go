package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "web/routers"
)

func main() {
	//beego.Run()
	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	beego.Run()
}

