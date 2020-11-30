package modules

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
	"usermanage/utils"
)

//SetDbNameAndMode 只设置DB文件的名字和格式，初始化接口对象
func SetDbNameAndMode(mode string) error {
	// --init gob/csv/json
	switch mode {
	case "gob":
		DbFilePath = "userDB.gob"
		var tmpf GobFile
		SwapDataX = tmpf
	case "csv":
		DbFilePath = "userDB.csv"
		var tmpf CsvFile
		SwapDataX = tmpf
	case "json":
		DbFilePath = "userDB.json"
		var tmpf JSONFile
		SwapDataX = tmpf
	default:
		return fmt.Errorf("请在以下参数选择输入: gob / csv / json")
	}
	return nil
}

//DistinguishFileType 根据文件后缀判断文件编码类型
func DistinguishFileType(fileType string) error {
	switch strings.ToLower(fileType) {
	case ".gob":
		var tmpf GobFile
		SwapDataX = tmpf
		return nil
	case ".csv":
		var tmpf CsvFile
		SwapDataX = tmpf
		return nil
	case ".json":
		var tmpf JSONFile
		SwapDataX = tmpf
		return nil
	default:
		fmt.Println("不能识别的文件格式")
		return fmt.Errorf("不能识别的文件格式")
	}
}

//LoadDbFile  将path文件内容解析为SliceU加载到内存中
func LoadDbFile(fpath string) error {
	if utils.IsFileExist(fpath) {
		fp, _ := os.Open(fpath)
		defer fp.Close()
		fileType := path.Ext(fpath)
		DistinguishFileType(fileType)
		SwapDataX.LoadFromFile(&SliceU, fp)

		//SliceUID为空,需要重新生成
		for _, u := range SliceU {
			SliceUID = append(SliceUID, u.ID)
		}

		return nil
	}
	return fmt.Errorf("DB文件不存在")
}

//SaveDbFile 组合命名path文件,把SliceU写入path中
func SaveDbFile(saveType string) error {
	errD := DistinguishFileType(saveType)
	if errD != nil {
		return errD
	}
	// fmt.Println("即将保存的数据库文件格式后缀为：", saveType) //debug
	now := time.Now().Format("20060102150405")
	savePath := "userDB_" + now + saveType //userDB_20060102150405.gob

	// fmt.Println("正在执行文件保存：", savePath) //debug
	dbf, errC := os.Create(savePath)
	if errC != nil {
		return errC
	}
	defer dbf.Close()
	errS := SwapDataX.SaveToFile(SliceU, dbf)
	if errS != nil {
		return errS
	}
	fmt.Println(savePath, "保存完毕!")
	return nil
}

//SaveOrNot ...
func SaveOrNot() {
	fmt.Print("是否保存修改到文件? Y/n: ")
	var ifsave string
	fmt.Scan(&ifsave)
	if strings.ToLower(ifsave) == "n" {
		fmt.Println("不保存退出!")
		return
	}
	saveType := "." + utils.Input("请输入保存的文件格式[gob/json/csv]: ")
	err := SaveDbFile(saveType)
	if err != nil {
		fmt.Println(err)
	}
}
