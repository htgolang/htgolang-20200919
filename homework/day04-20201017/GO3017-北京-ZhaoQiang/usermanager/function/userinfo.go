package function

//Users user info
/*
	每个元素
    ID
    名称
    联系方式
    通信地址
*/
var Users = []map[string]string{}

func init() {
	user1 := map[string]string{
		"id":   "1",
		"name": "zhao1",
		"tel":  "152XXXXXXXX",
		"addr": "山西省太原市",
	}
	user2 := map[string]string{
		"id":   "2",
		"name": "zhao2",
		"tel":  "152XXXXXXXX2",
		"addr": "山西省太原市2",
	}

	Users = append(Users, user1, user2)

}
