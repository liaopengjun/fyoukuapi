package routers

import (
	"fyoukuapi/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.VideoController{})
	beego.Include(&controllers.BaseController{})
	beego.Include(&controllers.CommentController{})
	beego.Include(&controllers.TopController{})
	beego.Include(&controllers.BarrageController{})
	beego.Include(&controllers.RedisDemoController{})
}
