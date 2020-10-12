package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/User/Login", &controllers.UserController{},"Post:Login")
}
