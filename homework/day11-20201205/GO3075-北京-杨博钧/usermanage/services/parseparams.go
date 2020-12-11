package services

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"usermanage/model"
)

// 添加用户参数解析
func ParseAddUpdateParams(r *http.Request, page *model.AddUpdatePage) error {
	page.Name = r.FormValue("name")
	if r.FormValue("sex") == "1" {
		page.IsMale = "checked"
		page.IsFeMale = ""
	} else {
		page.IsMale = ""
		page.IsFeMale = "checked"
	}
	page.Addr = r.FormValue("addr")
	page.Tel = r.FormValue("phone")
	_, err := time.Parse("2006-01-02", r.FormValue("birthday"))
	if err != nil {
		return err
	}
	page.Birthday = r.FormValue("birthday")
	page.Passwd = r.FormValue("password")
	return nil
}

// 查询用户参数解析
func ParseQueryParams(r *http.Request, page *model.MainPage) error {
	if r.FormValue("id") != "" {
		_, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			return err
		}
	}
	page.Id = r.FormValue("id")
	page.Name = r.FormValue("name")
	if r.FormValue("sex") == "男" || r.FormValue("sex") == "女" || r.FormValue("sex") == ""{
		page.Sex = r.FormValue("sex")
	} else {
		return fmt.Errorf("wrong sex filter")
	}
	page.Addr = r.FormValue("addr")
	page.Tel = r.FormValue("phone")

	if r.FormValue("birthday") != "" {
		_, err := time.Parse("2006-01-02", r.FormValue("birthday"))
		if err != nil {
			return err
		}
	}
	page.Birthday = r.FormValue("birthday")
	return nil
}
