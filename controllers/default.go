package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"fmt"
)

type MainController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UploadController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "upload.tpl"
}


func (c *UploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	fmt.Println(f,h,err)
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("uploadname", "F://Storage/" + h.Filename)
	c.Ctx.WriteString("ok")
}
