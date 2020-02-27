package routers

import (
	"Updater/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload",&controllers.UploadController{})
	beego.Router("/download",&controllers.DownloadController{})
	beego.Router("/update_latest",&controllers.UpdateLatestController{})
    beego.Router("/update_given", &controllers.UpdateGivenController{})
    beego.Router("/user/signup",&controllers.UserSignUpController{})
}