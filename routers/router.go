package routers

import (
	"Updater/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/file/upload",&controllers.UploadController{})
	beego.Router("/file/download",&controllers.DownloadController{})
    beego.Router("/file/list",&controllers.ListController{})
	beego.Router("/update_latest",&controllers.UpdateLatestController{})
    beego.Router("/update_given", &controllers.UpdateGivenController{})
    beego.Router("/user/signup",&controllers.UserSignUpController{})
	beego.Router("/user/signin",&controllers.UserSignInController{})
}