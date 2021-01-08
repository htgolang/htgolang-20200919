package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"user/models"
	"user/utils"
	"strconv"
)

type UserController struct {
	RequireAuth
}

func (c *UserController) Delete() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Query("id"), 10, 64)

	if err := utils.DeleteUserByID(id); err != nil {
		log.Println(err)
	}
	c.Redirect("/", 302)
}

func (c *UserController) Modify() {
	// 解析表单
	c.Ctx.Input.CopyBody(1024)
	data := c.Ctx.Input.RequestBody
	user := &models.User{}
	if err := json.Unmarshal([]byte(data), user); err != nil {
		log.Fatal(err)
	}

	// 查询用户
	userNow, err := utils.GetUserByID(user.ID)
	if err != nil {
		return
	}

	userNow.Name = user.Name
	userNow.Tel = user.Tel
	userNow.Addr = user.Addr
	userNow.Sex = user.Sex

	ormer := orm.NewOrm()
	if _, err := ormer.Update(userNow); err != nil {
		log.Fatal(err)
	}

	c.Ctx.Output.Body([]byte("ok"))
}

func (c *UserController) Add() {
	// 解析表单
	c.Ctx.Input.CopyBody(1024)
	data := c.Ctx.Input.RequestBody
	user := &models.User{}
	if err := json.Unmarshal([]byte(data), user); err != nil {
		log.Fatal(err)
	}

	// 解析参数
	curid := c.Ctx.Input.Query("curid")
	order := c.Ctx.Input.Query("order")

	var Step int64
	var Filter string

	if order == "before" {
		Step = 1
		Filter = "SortID__gte"
	} else {
		Step = -1
		Filter = "SortID__lte"
	}

	// 查询标杆用户
	curidInt, _ := strconv.ParseInt(curid, 10, 64)
	curUser, err := utils.GetUserByID(curidInt)
	if err != nil {
		log.Fatal(err)
	}

	// 查询标杆用户之后所有的用户, NextID + 1
	ormer := orm.NewOrm()
	qs := ormer.QueryTable(&models.User{})
	qs_users := []*models.User{}

	if _, err := qs.Filter(Filter, curUser.SortID).All(&qs_users); err != nil {
		log.Fatal(err)
	}

	// 开启事务
	if err := ormer.Begin(); err != nil {
		log.Fatal(err)
	}
	for _, a := range qs_users {
		a.SortID += Step
		if _, err := ormer.Update(a); err != nil {
			ormer.Rollback()
			log.Fatal(err)
		}
	}

	if err = utils.AddUser(user.Name, "", user.Addr, user.Sex, curUser.SortID); err != nil {
		ormer.Rollback()
		log.Fatal(err)
	}
	if err := ormer.Commit(); err != nil {
		log.Fatal(err)
	}

	c.Redirect("/", 302)
}
