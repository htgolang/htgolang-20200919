package models

var Users []map[string]string

func init() {
	/*
	   初始化数据
	   初始化Users，并添加4条基础数据
	   初始化UserPasswd，并转换为十六进制MD5值
	*/
	Users = make([]map[string]string, 0)
	AppendElement("1", "yizuo1", "17612345678", "888@qq.com")
	AppendElement("2", "yizuo2", "17612345678", "888@qq.com")
	AppendElement("3", "yizuo3", "17612345678", "888@qq.com")
	AppendElement("4", "yizuo4", "17612345678", "888@qq.com")

}

func AppendElement(ID, Name, Contact, Address string) {
	/*
	   将传递的值转换为字典并添加至我们的数据中
	*/
	Users = append(Users, GenerateElement(ID, Name, Contact, Address))
}

func GenerateElement(ID, Name, Contact, Address string) (element map[string]string) {
	/*
	   将传递进来的值转换为字典
	*/
	element = make(map[string]string)
	element["ID"], element["Name"], element["Contact"], element["Address"] = ID, Name, Contact, Address
	return
}

func ConvertElementToSlice(element map[string]string) (ret []string) {
	/*
	   字典类型的数据转换为切片类型
	*/
	ret = make([]string, 0)
	ret = append(ret, element["ID"])
	ret = append(ret, element["Name"])
	ret = append(ret, element["Contact"])
	ret = append(ret, element["Address"])
	return
}
