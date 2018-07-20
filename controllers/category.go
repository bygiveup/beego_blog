package controllers

import(
	"github.com/astaxie/beego"
	"x_blog/models"

)

type CategoryController struct{
	beego.Controller
}

func (this *CategoryController) Get(){
	//获取category.html页面中op的值,然后switch来获取操作
	op := this.Input().Get("op")
	switch op{
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0{
			break
		}
		err := models.AddCategory(name)
		if err != nil{
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0{
			break
		}
	}
	this.Data["IsCategory"] = true  //页面高亮显示
	this.TplName = "category.html"
	var err error
	this.Data["Categories"],err = models.GetAllCategories()
	if err != nil{
		beego.Error(err)
	}
}