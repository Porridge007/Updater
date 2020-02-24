package controllers

import (
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

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
	devices := QueryDeviceInfo()
	c.Data["devices"] = devices
	c.TplName = "upload.tpl"
}

func (c *UploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	fmt.Println(c.GetString("device"))
	fmt.Println(c.GetString("version"))
	fmt.Println(f, h, err)
	defer f.Close()

	// 保存文件操作 equal c.Ctx.SavetoFile
	file, head, err := c.Ctx.Request.FormFile("uploadname")
	defer file.Close()
	openFile, err := os.OpenFile("F://Storage/"+h.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer openFile.Close()
	io.Copy(openFile, file)

	//// 生成Sha1值
	//buf := bytes.NewBuffer(nil)
	//io.Copy(buf,file)
	//fmt.Println(util.Sha1(buf.Bytes()))
	//c.Ctx.WriteString("ok")

	if err != nil {
		fmt.Println("Failed to get data, err:", err.Error())
		return
	}
	defer file.Close()
	fileMeta := FileMeta{
		FileName: head.Filename,
		Location: "F://Storage/" + head.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		fmt.Println("Failed to create file, err:", err.Error())
		return
	}
	defer newFile.Close()
	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Failed to save data into file, err:", err.Error())
		return
	}
	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)
	fmt.Println(fileMeta)

	c.Ctx.WriteString("ok")
}
