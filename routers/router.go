package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"wxadmin/controllers"
)

func init() {
	admin.Run()
	beego.Router("/wx/banner/index", &controllers.BannerController{}, "*:Index")
	beego.Router("/wx/banner/add", &controllers.BannerController{}, "*:Add")
	beego.Router("/wx/banner/update", &controllers.BannerController{}, "*:Update")
	beego.Router("/wx/banner/del", &controllers.BannerController{}, "*:Del")
}
