package models

import (
	"wxadmin/util/mysql"
	"fmt"
)

const BANNER_TABLE = "wx"

type Banner struct {
	Id       	int		`db:"id"`
	ImgUrl  	string	`db:"img_url"`
	Title  		string	`db:"title"`
	Jumpurl  	string	`db:"jump_url"`
	Status 		int  	`db:"status"`
	CreatedAt  	int32	`db:"created_at"`
	UpdatedAt  	int32	`db:"updated_at"`
}

func (u *Banner) GetTableName() string {
	return BANNER_TABLE
}

func FindBannerList(page int64, pageSize int64, sort string) ([]mysql.MapModel, error) {
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	sql := fmt.Sprintf("select * from %s order by %s limit %d, %d", BANNER_TABLE, sort, pageSize, offset)
	return mysql.FindAll(sql)
}

func FindBanner(id string) (*Banner, error) {
	var info Banner
	err := mysql.FindCond(&info, map[string]string{"id": id}, "*")

	return &info, err
}

func (u *Banner) Update(data mysql.MapModel) bool {
	data.Load(u)
	return mysql.Update(u)
}

func InsertBanner(data mysql.MapModel) *Banner {
	var info Banner
	data.Load(&info)

	if ok := mysql.Insert(&info); ok {
		return &info
	}

	return nil
}
