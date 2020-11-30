package modules

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//SwapData mem-file 数据转换接口,适应三种不同的文件类型
type SwapData interface {
	SaveToFile([]Users, *os.File) error
	LoadFromFile(*[]Users, *os.File) error
}

//SwapDataX 全局接口变量,由init和load函数初始化
var SwapDataX SwapData

//GobFile ...
type GobFile struct {
	// Suffix string
}

//SaveToFile GobFile
func (t GobFile) SaveToFile(sliU []Users, f *os.File) (err error) {
	// fmt.Println("数据文件 格式 即将 编码为：gob!") //debug
	// fmt.Println("用户列表：", sliU) //debug
	//将空Users{}对象注册到gob管理器,告诉它按这种结构类型来编码
	gob.Register(Users{})
	//创建一个编码对象用来保存后的数据,并指定writer为文件指针f
	encdr := gob.NewEncoder(f)
	//将SliceU编码后保存到encdr
	err = encdr.Encode(sliU)
	if err == nil {
		// fmt.Println("编码成功!!") //debug
		return nil
	}
	return err
}

//LoadFromFile GobFile
func (t GobFile) LoadFromFile(pSU *[]Users, f *os.File) (err error) {
	// fmt.Println("即将 解码 Gob 格式数据文件!") //debug
	decdr := gob.NewDecoder(f)
	err = decdr.Decode(pSU)
	if err == nil {
		// fmt.Println("Gob解码成功!!") //debug
	}
	return err
}

//CsvFile ...
type CsvFile struct{}

//SaveToFile CsvFile
func (t CsvFile) SaveToFile(sliU []Users, f *os.File) (err error) {
	// fmt.Println("数据文件 格式 即将 编码为：CSV!") //debug
	nwt := csv.NewWriter(f)
	for _, u := range sliU {
		nwt.Write([]string{strconv.Itoa(u.ID), u.Name, u.Password, u.Tel, u.Addr, u.Birthday.Format("2006-01-02"), strconv.FormatBool(u.Deleted), strconv.FormatBool(u.Ifadmin)})
	}
	nwt.Flush()
	// fmt.Println("CSV编码成功!!") //debug
	return err
}

//LoadFromFile CsvFile
func (t CsvFile) LoadFromFile(pSU *[]Users, f *os.File) (err error) {
	// fmt.Println("即将 读取 CSV 格式数据文件!") //debug
	nrdr := csv.NewReader(f)
	var u Users

	ctx, err := nrdr.ReadAll()
	if err == nil {
		return err
	}
	for _, line := range ctx {
		u.ID, _ = strconv.Atoi(line[0])
		u.Name = line[1]
		u.Password = line[2]
		u.Tel = line[3]
		u.Addr = line[4]
		u.Birthday, _ = time.Parse("2006-01-02 15:04:05", line[5])
		u.Deleted, _ = strconv.ParseBool(line[6])
		u.Ifadmin, _ = strconv.ParseBool(line[7])
		*pSU = append(*pSU, u)
	}
	// fmt.Println("成功加载CSV文件到内存") //debug
	return nil
}

//JSONFile ...
type JSONFile struct{}

//SaveToFile 将一个变量的值格式化为字符串保存为JSONFile
func (t JSONFile) SaveToFile(sliU []Users, f *os.File) (err error) {
	// fmt.Println("数据文件 格式 即将 编码为：JSON!") //debug

	slibyte, err := json.Marshal(sliU)
	// slibyte, err := json.MarshalIndent(sliU, "", "\t")
	if err != nil {
		fmt.Println("Json 编码错误", err.Error())
	}
	f.Write(slibyte)
	// fmt.Println("Json编码成功\n", string(slibyte)) //debug
	return err
}

//LoadFromFile 将JSONFile内容加载到一个变量,传参为变量的指针
func (t JSONFile) LoadFromFile(pSU *[]Users, fp *os.File) error {
	// fmt.Println("即将 解码 Json 格式数据文件!") //debug
	defer fp.Close()

	slicebyte, _ := ioutil.ReadAll(fp)
	// fmt.Println(slicebyte)

	//判断是否为json格式
	ok := json.Valid(slicebyte)
	if ok {
		// fmt.Println("CSV解码成功!!") //debug
		json.Unmarshal(slicebyte, pSU)
		// fmt.Println(pSU, &SliceU) //debug
		return nil
	}
	return fmt.Errorf("数据并非json格式")
}
