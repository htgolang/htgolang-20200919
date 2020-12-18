package services

import (
	"fmt"
	"strconv"
	"time"
	"usermanage/forms"
	"usermanage/model"
)

// 添加用户参数解析
func ParseAddUpdateParams(user forms.User, page *model.AddUpdatePage) error {
	page.Name = user.Name
	if user.Sex {
		page.IsMale = "checked"
		page.IsFeMale = ""
	} else {
		page.IsMale = ""
		page.IsFeMale = "checked"
	}
	page.Addr = user.Addr
	page.Tel = user.Tel
	_, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return err
	}
	page.Birthday = user.Birthday
	page.Passwd = user.Password
	return nil
}

// 查询用户参数解析
func ParseQueryParams(info forms.QueryInfo, page *model.MainPage) error {
	if info.Id != "" {
		_, err := strconv.Atoi(info.Id)
		if err != nil {
			return err
		}
	}
	page.Id = info.Id
	page.Name = info.Name
	if info.Sex == "男" || info.Sex == "女" || info.Sex == ""{
		page.Sex = info.Sex
	} else {
		return fmt.Errorf("wrong sex filter")
	}
	page.Addr = info.Addr
	page.Tel = info.Tel

	if info.Birthday != "" {
		_, err := time.Parse("2006-01-02", info.Birthday)
		if err != nil {
			return err
		}
	}
	page.Birthday = info.Birthday
	return nil
}
