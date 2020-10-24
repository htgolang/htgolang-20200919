package models

// 结构体名称大写 属性名称大写
// 结构体名称大写 属性名称小写
type PublicStruct struct {
	privateAttrPu string
	PublicAttrPu  string
}

//结构体名称小写  属性名称大写
//结构体名称小写 属性名称小写
type privateStruct struct {
	privateAttrPr string
	PublicAttrPr  string
}

func NewPrivateStruct() *privateStruct {
	return &privateStruct{}
}

/*
type CombindStruct struct {
	PublicAttr  PublicStruct
	PrivateAttr privateStruct
}
*/
type CombindStruct struct {
	PublicStruct
	privateStruct
}
