package controllers

import (
	"github.com/astaxie/beego"
	"log"
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
	// 读取文件信息
	f, h, err := c.GetFile("filename")

	if err != nil {
		log.Fatal("读取文件错误", err)
	}

	// 延迟关闭文件
	defer f.Close()

	// 保存文件, 本地文件路径static/upload/上传文件名
	// 需要提前创建好static/upload目录
	c.SaveToFile("filename", "F://Storage/" + h.Filename)
}

