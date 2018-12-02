package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"wxadmin/controllers"
	"wxadmin/services"
)

func init() {
	services.InitConfig()
	admin.Run()
	beego.Router("/wx/banner/index", &controllers.BannerController{}, "*:Index")
	beego.Router("/wx/banner/add", &controllers.BannerController{}, "*:AddBanner")
	beego.Router("/wx/banner/update", &controllers.BannerController{}, "*:UpdateBanner")
	beego.Router("/wx/banner/del", &controllers.BannerController{}, "*:DelBanner")
}
