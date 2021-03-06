package routers

import (
	"Updater/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/file/upload", &controllers.UploadController{})
	beego.Router("/file/download", &controllers.DownloadController{})
	beego.Router("/file/list", &controllers.ListController{})
	beego.Router("/query_latest", &controllers.QueryLatestController{})
	beego.Router("/update_latest", &controllers.UpdateLatestController{})
	beego.Router("/update_given", &controllers.UpdateGivenController{})
	beego.Router("/update_latest_by_path",&controllers.QueryLatestPathController{})

	beego.Router("/user/signup", &controllers.UserSignUpController{})
	beego.Router("/user/signin", &controllers.UserSignInController{})
}