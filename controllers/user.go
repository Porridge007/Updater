package controllers

import (
	"Updater/common"
	"Updater/models"
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"strconv"
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

	fmt.Println(username, userpwd)

	if len(username) < 3 || len(userpwd) < 5 {
		ret := common.RetData{
			Code: 400,
			Msg:  "Username is too short",
			Data: nil,
		}
		c.Data["json"] = &ret
		c.ServeJSON()


	}

	encPasswd := util.Sha1([]byte(userpwd + pwdSalt))
	fmt.Println(encPasswd)

	user := models.User{
		UserName: username,
		UserPwd:  encPasswd,
	}

	fmt.Println(user)

	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		ret := common.RetData{
			Code: 300,
			Msg:  "DB error",
			Data: user.UserName,
		}
		c.Data["json"] = &ret
		c.ServeJSON()
	}else{
		ret := common.RetData{
			Code: 200,
			Msg:  "Register succeeded",
			Data: user.UserName,
		}
		c.Data["json"] = &ret
		c.ServeJSON()
	}
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
	if count <= 0 || err != nil {
		c.Ctx.ResponseWriter.Write([]byte("FAILED"))
		return
	}
	fmt.Println(count, 12)
	c.Ctx.ResponseWriter.Write([]byte("SUCCESS"))

	token := GenToken(userName)
	fmt.Println(token)
	// "replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	_, err = o.Raw("replace into user_token set user_name = ?, token = ?", userName, token).Exec()
	fmt.Println()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("FAILED2"))
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

func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	//  判断token的时效性，是否过期
	tsNow := time.Now().Unix()
	tokenTime, _ := strconv.ParseInt(token[31:], 16, 64)
	if tsNow-tokenTime > 86400 {
		fmt.Println(tsNow - tokenTime)
		return false
	}
	// 从数据库表tbl_user_token查询username对应的token信息
	o := orm.NewOrm()
	if !o.QueryTable("user_token").Filter("token", token).Exist(){
		return false
	}

	return true
}
