package routers

import (
	"x_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.Router("/login", &controllers.LoginController{})
}
