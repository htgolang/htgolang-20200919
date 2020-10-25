package main

import (
	"userManagementV2/controller"
	"userManagementV2/dao"
	"userManagementV2/manager"
	"userManagementV2/model"
	"userManagementV2/service"
)

func main() {
	db := model.NewUserDB()                                         //初始化数据库
	dbDao := dao.NewUserDao(db)                                     //初始化DAO
	userManager := manager.NewUserManagerImp(dbDao)                 //初始化manager
	userService := service.NewUserService(userManager)              //初始化service
	userController := controller.NewUserController("", userService) //初始化controller
	userController.MainMenu()                                       //启动
}
