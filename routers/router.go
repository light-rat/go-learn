package routers

import (
	"github.com/astaxie/beego"
	"goTest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Include(&controllers.LoginController{})
}
