package controllers

import (
	"Updater/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserSignUpController struct {
	beego.Controller
}

type UserSignInController struct {
	beego.Controller
}

func (c *UserSignUpController) Get() {
	c.TplName = "signup.tpl"
}

func (c *UserSignUpController) Post() {
	username := c.GetString("username")
	userpwd := c.GetString("password")

	user := models.User{
		UserName: username,
		UserPwd:  userpwd,
	}

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println(id, err.Error())
		c.Ctx.ResponseWriter.Write([]byte("FAILED"))
	}
	c.Ctx.ResponseWriter.Write([]byte("SUCCESS"))
}

func (c *UserSignInController) Get() {
	c.TplName = "signin.tpl"
}
