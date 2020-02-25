package routers

import (
	"Updater/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload",&controllers.UploadController{})
	beego.Router("/download",&controllers.DownloadController{})
}