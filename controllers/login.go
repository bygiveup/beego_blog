package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

type LoginController struct{
	beego.Controller
}

func (this *LoginController) Get() {
	//退出删除cookie
	isExit := this.Input().Get("exit") == "true"
	if isExit{
		this.Ctx.SetCookie("uname","",-1,"/")   //-1立即删除cookie
		this.Ctx.SetCookie("pwd","",-1,"/")
		this.Redirect("/",301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
//	this.Ctx.WriteString(fmt.Sprint(this.Input()))  //fmt.Sprint打印出来
//	return
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && 
		beego.AppConfig.String("pwd") == pwd{
		maxAge := 0
		if autoLogin{
			maxAge = 1<<31 -1
		}
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")
	}
	this.Redirect("/",301)	
	return
}


func checkAccount(ctx *context.Context) bool {
	ck,err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	p,err := ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := p.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd


}
