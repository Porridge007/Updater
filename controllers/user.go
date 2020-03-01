package controllers

import (
	"Updater/models"
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"time"
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
		fmt.Println(213)
		fmt.Println(id, err.Error())
		c.Ctx.ResponseWriter.Write([]byte("FAILED"))
	}
	c.Ctx.ResponseWriter.Write([]byte("SUCCESS"))
}

func (c *UserSignInController) Get() {
	c.TplName = "signin.html"
}

func (c *UserSignInController) Post() {
	userName := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(userName)
	fmt.Println(password)

	encPasswd := util.Sha1([]byte(password + pwdSalt))
	fmt.Println(encPasswd)
	o := orm.NewOrm()

	count, err := o.QueryTable("User").Filter("UserName", userName).Filter("UserPwd", encPasswd).Count()
	fmt.Println(count)
	if count <= 0 ||err != nil {
		c.Ctx.ResponseWriter.Write([]byte("1231"))
		return
	}
	fmt.Println(count,12)
	c.Ctx.ResponseWriter.Write([]byte("SUCCESS"))

	token := GenToken(userName)
	fmt.Println(token)
	// "replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	_,err = o.Raw("replace into user_token set user_name = ?, token = ?", userName, token).Exec()
	fmt.Println()
	if err != nil {
		fmt.Println(11122)
		c.Ctx.ResponseWriter.Write([]byte("FAILED2"))
		fmt.Println(232323)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}




func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
