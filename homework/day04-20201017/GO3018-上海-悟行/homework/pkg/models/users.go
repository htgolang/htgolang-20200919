package models

var Users []map[string]string

func init() {
	//初始化数据
	Users = make([]map[string]string,0)
	AppendElement("1","a1","11111111111","Shanghai")
	AppendElement("2","a2","22222222222","shenzhen")
	AppendElement("3","a3","33333333333","beijing")
	AppendElement("4","a4","44444444444","Guangzhou")
	AppendElement("5","a5","55555555555","Nanjing")
	AppendElement("6","a6","66666666666","Chongqing")
}

func AppendElement(ID,Name,Contact,Address string) {
	//将传入的值转为字典并添加到数据中
	Users = append(Users,GenerateElement(ID,Name,Contact,Address))
}

func GenerateElement(ID,Name,Contact,Address string) (element map[string]string) {
	//将值转为字典
	element = make(map[string]string)
	element["ID"],element["Name"],element["Contact"],element["Address"] = ID,Name,Contact,Address
	return
}

func ConvertElementToSlice(element map[string]string) (ret []string) {
	//  字典类型的数据转换为切片类型
	ret = make([]string, 0)
	ret = append(ret, element["ID"])
	ret = append(ret, element["Name"])
	ret = append(ret, element["Contact"])
	ret = append(ret, element["Address"])
	return
}