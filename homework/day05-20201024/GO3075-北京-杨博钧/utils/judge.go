package utils
//判断一个用户是否存在,存在时返回切片下标
func isUserExists(name string) (int, bool) {
	for i, user := range UsersList {
		if name == user.Name {
			return i, true
		}
	}
	return 0, false
}