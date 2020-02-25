package controllers

import (
	"Updater/models"
	"Updater/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
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

type DownloadController struct {
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
	defer f.Close()

	// 保存文件操作 equal c.Ctx.SavetoFile
	file, head, err := c.Ctx.Request.FormFile("uploadname")
	defer file.Close()
	openFile, err := os.OpenFile("F://Storage/"+h.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer openFile.Close()
	io.Copy(openFile, file)

	if err != nil {
		fmt.Println("Failed to get data, err:", err.Error())
		return
	}
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
		fmt.Println(id)
		c.Ctx.WriteString("The file has been upload")
		return
	}

	c.Ctx.WriteString("ok")
}

//func DownloadHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	fsha1 := r.Form.Get("filehash")
//	fm := meta.GetFileMeta(fsha1)
//	f, err := os.Open(fm.Location)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	defer f.Close()
//
//	data, err := ioutil.ReadAll(f)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/octect-stream")
//	w.Header().Set("Content-Disposition", "attachment;filename=\""+fm.FileName+"\"")
//	w.Write(data)
//}

func (c *DownloadController) Post() {
	fsha1 := c.GetString("filehash")
	o := orm.NewOrm()
	qs := o.QueryTable("file")

	fmt.Println(qs.Filter("file_sha1", fsha1))
	c.Ctx.WriteString("victory")

}
