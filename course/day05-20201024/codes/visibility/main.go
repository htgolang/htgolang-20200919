package main

import (
	"fmt"
	"visibility/models"
)

func main() {
	publicVar := models.PublicStruct{}

	fmt.Printf("%#v\n", publicVar)
	fmt.Println(publicVar.PublicAttrPu)
	// fmt.Println(publicVar.privateAttr)
	// fmt.Println(models.privateStruct{})

	privateVar := models.NewPrivateStruct()
	fmt.Printf("%#v\n", privateVar)
	fmt.Println(privateVar.PublicAttrPr)
	// fmt.Println(privateVar.privateAttr)

	combindVar := models.CombindStruct{}
	fmt.Println(combindVar.PublicStruct.PublicAttrPu)
	fmt.Println(combindVar.PublicAttrPu)
	// fmt.Println(combindVar.privateStruct.PublicAttrPr)
	fmt.Println(combindVar.PublicAttrPr)

}
