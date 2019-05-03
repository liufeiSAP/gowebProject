package routers

import (
	"gowebProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users/user", &controllers.UserController{}, "Post:AddUser")
	beego.Router("/users/users", &controllers.UserController{}, "Post:AddUsers")
}
