package main

import (
	_"beblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beblog/models"
	"beblog/controllers"
)

func init(){
	models.RegisterDB()
}

func main() {
	//debug info
	orm.Debug = true
	//set db-table
	orm.RunSyncdb("default",false,true)

	//set router
	beego.Router("/",&controllers.MainController{})
	beego.Router("/login.html",&controllers.LoginController{})
	beego.Router("/category.html",&controllers.CategoryController{})
	beego.Router("/topic.html",&controllers.TopicController{})
	beego.Router("/reply",&controllers.ReplyController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete")

	beego.AutoRouter(&controllers.TopicController{})
	// start beego
	beego.Run()
}

