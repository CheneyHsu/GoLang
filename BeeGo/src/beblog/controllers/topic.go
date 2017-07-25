package controllers

import (
	"github.com/astaxie/beego"
	"beblog/models"
	"strings"
)

type TopicController struct {
	beego.Controller
}


func (c *TopicController) Get(){
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName="topic.html"
	topics,err:=models.GetAllTopics("","",false)
	if err != nil{
		beego.Error(err)
	}else {
		c.Data["Topics"]=topics
	}

}

func (c *TopicController) Post() {
if !checkAccount(c.Ctx){
	c.Redirect("/login.html",302)
	return
}
//	title := c.Input().Get("title")
//	content := c.Input().Get("content")

	var err error
	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category :=c.Input().Get("category")
	label := c.Input().Get("label")
	if len(tid) == 0{
		err = models.AddTopic(title,category,label,content)
	}else{
		err =models.ModifyTopic(tid,title,category,label,content)
	}
//	err = models.AddTopic(title,content)
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/topic.html",302)
}

func (c *TopicController) Add()  {
	c.TplName="topic_add.html"
//	c.Ctx.WriteString("add")

}

func (c *TopicController) View()  {
	c.TplName="topic_view.html"
	topic,err:=models.GetTopic(c.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
		c.Redirect("/",302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Labels"]=strings.Split(topic.Labels," ")
//	c.Data["Tid"]=c.Ctx.Input.Param("0")
	replies,err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil{
		beego.Error(err)
		return
	}
	c.Data["Replies"]=replies
	c.Data["IsLogin"]=checkAccount(c.Ctx)
}

func (c *TopicController) Modify() {
	c.TplName ="topic_modify.html"
	tid := c.Input().Get("tid")
	topic,err := models.GetTopic(tid)
	if err != nil{
		beego.Error(err)
		c.Redirect("/",302)
		return
	}
	c.Data["Topic"]=topic
	c.Data["Tid"]=tid
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx){
		c.Redirect("/login.html",302)
		return
	}
	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/",302)
}
