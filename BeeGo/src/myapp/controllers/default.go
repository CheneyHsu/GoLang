package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "myapp.com"
	c.Data["Email"] = "xuchenglin@hrbb.com.cn"
	c.TplName = "index.tpl"
//	c.Ctx.WriteString("appname: "+beego.AppConfig.String("appname")+"\nhttpport: "+beego.AppConfig.String("httpport")+"\nrunmode: "+beego.AppConfig.String("runmode"))
	beego.Trace("trace test1")
	beego.Info("info test1")
	beego.SetLevel(beego.LevelInformational)
	beego.Trace("trace test2")
	beego.Info("info test2")
}
