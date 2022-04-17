package routers

import (
	"ClaytonUniversityLibrary/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.IndexController{})
	//beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/users/login", &controllers.UserLoginController{})
	beego.Router("/users/logout", &controllers.UserLogoutController{})
	beego.Router("/users/register", &controllers.UserRegisterController{})
	beego.Router("/users/update", &controllers.UserUpdateController{})
	beego.Router("/papers", &controllers.PapersController{})
	beego.Router("/papers/upload", &controllers.UploadPaperController{})
	//beego.Router("/papers/peerreview", &controllers.PeerReviewController{})
}
