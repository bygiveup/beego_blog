package main

import (
	_ "x_blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "x_blog/models"
)


func main() {
	//开启orm调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default",false,true)   //如果false设置成true会一直删除重建,如果不希望每次都打印信息,可以将true设置成false
	beego.Run()
}

