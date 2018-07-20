package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true    //默认高亮,在T.navbar.tpl中调用
	c.TplName = "home.html"
}
