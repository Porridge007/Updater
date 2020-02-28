package controllers

import (
	"Updater/models"
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	pwdSalt = "*#1989"
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

	if len(username) < 3 || len(userpwd) < 5 {
		c.Ctx.ResponseWriter.Write([]byte("Invalid parameter"))
		return
	}

	encPasswd := util.Sha1([]byte(userpwd + pwdSalt))

	user := models.User{
		UserName: username,
		UserPwd:  encPasswd,
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

//func (c *UserSignInController) Post() {
//	userName := c.GetString("username")
//	password := c.GetString("password")
//
//
//
//	encPasswd := util.Sha1([]byte(password + pwdSalt))
//}


