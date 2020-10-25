package service

import (
	"userManagementV2/manager"
	"userManagementV2/model"
)

//具体业务逻辑处理：数据合法性校验，并调用管理

var _ UserService = (*userService)(nil)

type userService struct {
	userManager manager.UserManager
}

func NewUserService(userManager manager.UserManager) *userService {
	return &userService{
		userManager: userManager,
	}
}

func (s *userService) CreateUser(user model.User) error {
	err := s.userManager.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) ListUser() []*model.User {
	res, err := s.userManager.GetUserList()
	if err != nil {
		return nil
	}
	return res
}

//todo 服务
//func (s *UserService) GetUserNum() int {
//	return s.userNum
//}
//
//func (s *UserService) HasUser(userId int) (int, error) {
//	for i, v := range s.users {
//		if v.Id == userId {
//			return i, nil
//		}
//	}
//	return -1, errors.New("查无此用户")
//}
//
//func (s *UserService) GetUser(idx int) []model.User {
//	return []model.User{s.users[idx]}
//}
//
//func (s *UserService) Add(user model.User) bool {
//	s.userNum++
//	s.curId++
//	user.Id = s.curId
//	s.users = append(s.users, user)
//	return true
//}
//
//func (s *UserService) Modify(idx int, user model.User) bool {
//	s.users[idx].Name = user.Name
//	s.users[idx].Phone = user.Phone
//	s.users[idx].Address = user.Address
//	return true
//}
//
//func (s *UserService) Delete(idx int) bool {
//	s.users = append(s.users[:idx], s.users[idx+1:]...)
//	s.userNum--
//	return true
//}
//
//func (s *UserService) List() []model.User {
//	return s.users
//}
//
//func (s *UserService) Query(k string) []model.User {
//	re := regexp.MustCompile(k)
//	if re == nil {
//		fmt.Println("regexp err")
//		return nil
//	}
//	users := []model.User{}
//	for _, v := range s.users {
//		if re.MatchString(v.Name) {
//			users = append(users, v)
//			continue
//		}
//		if re.MatchString(v.Address) {
//			users = append(users, v)
//			continue
//		}
//		if re.MatchString(v.Phone) {
//			users = append(users, v)
//			continue
//		}
//	}
//	return users
//}
