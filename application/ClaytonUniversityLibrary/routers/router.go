package routers

import (
	"ClaytonUniversityLibrary/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.IndexController{})
}
