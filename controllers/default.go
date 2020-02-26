package controllers

import (
	"Updater/models"
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type MainController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}

type DownloadController struct {
	beego.Controller
}

type UpdateLatestController struct {
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
	file, head, err := c.GetFile("uploadname")
	defer file.Close()

	fileInfo := models.File{
		File_name: head.Filename,
		File_addr: "F://Storage/" + head.Filename,
		Device:    c.GetString("device"),
		Version:   c.GetString("version"),
	}

	newFile, err := os.Create(fileInfo.File_addr)
	if err != nil {
		fmt.Println("Failed to create file, err:", err.Error())
		return
	}
	defer newFile.Close()

	fileInfo.File_size, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Failed to save data into file, err:", err.Error())
		return
	}
	newFile.Seek(0, 0)
	fileInfo.File_sha1 = util.FileSha1(newFile)

	o := orm.NewOrm()
	id, err := o.Insert(&fileInfo)
	if err != nil {
		logs.Error(id,err.Error())
		c.Ctx.WriteString("The file has been upload")
		return
	}

	c.Ctx.WriteString("ok")
}


func (c *DownloadController) Post() {
	fsha1 := c.GetString("filehash")

	o := orm.NewOrm()
	fileMeta := models.File{File_sha1:fsha1}
	err := o.QueryTable("file").Filter("file_sha1",fsha1).One(&fileMeta)

	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	f, err := os.Open(fileMeta.File_addr)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octect-stream")
	c.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.File_name+"\"")
	c.Ctx.ResponseWriter.Write(data)

	c.Ctx.WriteString("victory")

}

func (c *UpdateLatestController) Get() {
	device := c.GetString("device")
	o := orm.NewOrm()
	fileMeta := models.File{Device:device}

	err := o.QueryTable("file").OrderBy("-created").One(&fileMeta)
	if err != nil{
		return
	}
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}

	f, err := os.Open(fileMeta.File_addr)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octect-stream")
	c.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.File_name+"\"")
	c.Ctx.ResponseWriter.Write(data)

	c.Ctx.WriteString("victory")
}


