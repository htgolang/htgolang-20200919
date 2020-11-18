package models

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
	"yizuo/utils"
)

/*
所有跟数据相关操作在此文件中进行
*/

// 根据传递参数添加用户数据
func ReadUserList(curId, Status, Id, Name, Password, Phone, Address, Birthday string )  {
	/*
		根据传递的用户信息添加至对应的数据库
		u 用户数据信息
			Id			固定值不允许变更，每次新增在现有数据中检索最大值+1作为新用户的ID
			Name		用户名称
			Password	用户密码
			Phone		手机号
			Address		邮箱地址
			Birthday	生日
		users
			curId		唯一值，值与用户的ID值保持一致
			Status		用户状态信息，用于软删除。0位用户不再使用，1位用户正在使用中，默认为1。
			UserData	用户数据信息
	*/
	// 根据输入值插入数据
	var u = &User {
		Id:			cast.ToInt(Id),                     // 用户id
		Name:		Name, 					// 名称
		Password:	Password,               // 密码
		Phone:		Phone, 					// 联系方式
		Address:	Address, 				// 地址
		Birthday:	utils.StrConversionTime(Birthday), // 生日
	}

	var users = Users{
		curId:		cast.ToInt(curId),     // ID
		Status:		cast.ToInt(Status),    // 用户状态。0为软删除，1为在使用
		UserData:	u,         // 传入数据
	}
	UserList = append(UserList,users)
}


// 根据传递参数添加用户数据
func AddUser(Name, Password, Phone, Address, Birthday string )  {
    /*
		根据传递的用户信息添加至对应的数据库
		u 用户数据信息
			Id			固定值不允许变更，每次新增在现有数据中检索最大值+1作为新用户的ID
			Name		用户名称
			Password	用户密码
			Phone		手机号
			Address		邮箱地址
			Birthday	生日
		users
			curId		唯一值，值与用户的ID值保持一致
			Status		用户状态信息，用于软删除。0位用户不再使用，1位用户正在使用中，默认为1。
			UserData	用户数据信息
	*/
	// 根据输入值插入数据
	var u = &User {
		Id:			FindLargestElementID(), // 用户id
		Name:		Name, 					// 名称
		Password:	utils.Md5sum(Password), // 密码
		Phone:		Phone, 					// 联系方式
		Address:	Address, 				// 地址
		Birthday:	utils.StrConversionTime(Birthday), // 生日
	}

	var users = Users{
		curId:		u.Id, // ID
		Status:		1 ,    // 用户状态。0为软删除，1为在使用
		UserData:	u,    // 传入数据
	}
	UserList = append(UserList,users)
}

// 软删除用户数据
func DeleteUser(UserID int) bool {
	/*
	遍历数据找到对应的数据位置
	将Status设置为0软删除用户
	  PS：默认所有的用户操作Status都要为1才进行
	 */
	for k, v := range UserList {
		// 判断如果遍历的数据与需要删除的数据相同s
		if v.UserData.Id == UserID  {
			// 打印根据ID搜索出的条目
			FormatTableOut(v.UserData)
			// 让用户确认是否变更此数据
			confirm := utils.Input("以上是该ID对应的信息，是否删除此条数据。yes/no: ")
			if strings.ToLower(confirm) == "yes" || strings.ToLower(confirm) == "y" {
				UserList[k].Status = 0
				return true
			}
		}
	}
	return false
}

// 根据输入ID，让用户输入数据变更
func ModifyUser(UserID int) bool {
	/*
	遍历数据，判断用户ID是否存在，且用户Status为1
	让用户输入需要变更的数据，变更内存中数据。
	*/
	for _, v := range UserList {
		// 判断如果遍历的数据与需要删除的数据相同s
		if v.UserData.Id == UserID && v.Status == 1 {
			// 打印根据ID搜索出的条目
			FormatTableOut(v.UserData)
			// 让用户确认是否变更此数据
			confirm := utils.Input("以上是该ID对应的信息，是否变更该ID数据。yes/no: ")
			if strings.ToLower(confirm) == "yes" || strings.ToLower(confirm) == "y" {
				// 判断用户是否存在，存在即退出
				name := utils.Input("请输入要变更的用户名：")
				if FindElementUser(name) {
					fmt.Printf("用户%v已存在，已退出。\n",name)
					return false
				}
				v.UserData.Name = name
				v.UserData.Password = utils.Md5sum(utils.Input("请输入变更后的用户密码："))
				v.UserData.Phone = utils.Input("请输入变更后的联系方式（例如：17612345678）：")
				v.UserData.Address = utils.Input("请输入变更后的联系地址（例如：WuHan）：")
				v.UserData.Birthday = utils.StrConversionTime(utils.Input("请输入变更后的生日（例如：1994-04-06 18:08:06）："))
				// 打印变更后的信息
				fmt.Println("变更成功。变更后的信息如下：")
				FormatTableOut(v.UserData)
				return true
			}
			break
		}
	}
	return false
}

// 根据输入的用户，返回对应的数据
func QueryUser(str string) (*User, error) {
	for _,v := range UserList{
		if str == v.UserData.Name {
			return v.UserData, nil
		}
	}
	return nil, nil
}

// 查找所有的用户数据返回，如果后期写入数据库，通过此函数进行
func ListUser() []Users {
	// 调试代码，查看当前所有的数据以及对应用户的状态,方便查看软删除的用户
	fmt.Printf("%T,%v \n",UserList,UserList)

	// 查找所有状态为1的用户，并返回
	var ret []Users
	for _,v := range UserList{
		if v.Status == 1 {
			ret = append(ret, v)
		}
	}
	return ret
}

// 判断输入id的是否在数据库中存在
func FindElementID(id int) bool {
	/*
	判断用户ID是否在数据库中存在
	用户ID存在返回true，用户ID不存在返回false
	 */
	for _,v := range UserList {
		if v.curId == id && v.Status == 1 {
			return true
		}
	}
	return false
}

// 判断输入用户在数据库中是否存在
func FindElementUser(UserName string) bool {
	/*
	判断用户是否在数据库中存在
	用户存在返回true，用户不存在返回false
	 */
	for _,v:= range UserList{
		if UserName == v.UserData.Name {
			return true
		}
	}
	return false
}

// 查看现有数据中ID最大的值，将这个值+1作为我们新用户的ID并返回
func FindLargestElementID() (ID int) {
	var num int
	for _, v := range UserList {
		// 对比大小将ID的值设置为最大的那个
		if num < v.curId {
			num = v.curId
		}
	}
	// 将最后找到的值+1返回
	ID = num + 1
	return
}

// 用户登录信息检测
func UserLoginAuth(user, passwd string) bool {
	for _,v := range UserList {
		if v.UserData.Name == user && v.UserData.Password == utils.Md5sum(passwd){
			return true
		}
	}
	return false
}


// 系统启动用户数据读取
func ReadUsers()  {
	_, err := os.Lstat(UserDataFile)
	// 检查文件是否存在，如果不存在新增默认数据
	if os.IsNotExist(err){
		// 添加三条默认数据
		addDefaultUsers()
		// 初始化的数据落盘
		WritesUsersDataToCsv()
	} else {
		ReadUsersDataToCsv()
	}
}

// 用户数据初始化
func InitAllUser() {
	// 删除所有数据
	deleteAllUsers()
	// 添加三条默认数据
	addDefaultUsers()
	// 现有数据落盘
	WritesUsersDataToCsv()
}

// 添加三条默认数据
func addDefaultUsers()  {
	AddUser("yizuo", "yizuo", "17612345678", "WuHan", "1994-04-06 18:00:00")
	AddUser("admin", "admin", "17612345678", "ShangHai", "1995-04-06 18:00:00")
	AddUser("root", "root", "17612345678", "WuHan", "1996-04-06 18:00:00")
}

// 删除用户所有数据
func deleteAllUsers() {
	UserList = UserList[0:0]
}