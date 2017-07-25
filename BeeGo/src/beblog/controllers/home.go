package controllers

import (
	"github.com/astaxie/beego"
	"beblog/models"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics,err:=models.GetAllTopics(c.Input().Get("cate"),c.Input().Get("label"),true)
	if err != nil{
		beego.Error(err)
	}else {
		c.Data["Topics"]=topics
	}
	categroies,err :=models.GetAllCategories()
	if err != nil{
		beego.Error(err)
	}
	c.Data["Categories"] = categroies
}

