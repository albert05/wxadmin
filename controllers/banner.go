package controllers

import (
	"github.com/beego/admin/src/rbac"
	"wxadmin/models"
	"wxadmin/util/mysql"
)

type BannerController struct {
	rbac.CommonController
}

func (this *BannerController) Index() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	//sort := this.GetString("sort")

	banners, count := models.FindBannerList(page, page_size, "id desc")
	if this.IsAjax() {
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &banners}
		this.ServeJSON()
		return
	}

	this.TplName = this.GetTemplatetype() + "/wx/banner.tpl"
}

func (this *BannerController) AddBanner() {
	d := mysql.MapModel{}
	if err := this.ParseForm(&d); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}

	banner := models.InsertBanner(d)
	if banner.Id > 0 {
		this.Rsp(true, "Success")
		return
	}

	this.Rsp(false, "Failed")

}

func (this *BannerController) UpdateBanner() {
	banner := models.Banner{}
	if err := this.ParseForm(&banner); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}

	ret := banner.Update(nil)
	if ret {
		this.Rsp(true, "Success")
		return
	}

	this.Rsp(false, "Failed")
}

func (this *BannerController) DelBanner() {
	Id := this.GetString("Id")

	banner, err := models.FindBanner(Id)
	if err != nil {
		this.Rsp(false, err.Error())
		return
	}

	ret := banner.Update(mysql.MapModel{
		"status": 0,
	})
	if ret {
		this.Rsp(true, "Success")
		return
	}

	this.Rsp(false, err.Error())
}
