package controller

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	"io"
	"strconv"
	"strings"
	"usermanager_beego/users"
)

type LogAnalysis struct {
	Authcontroller
}

type Dateone struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

//type IpOne struct {
//	IpAddr string
//	Total int
//}
type TotalData struct {
	Datapie []Dateone
	Databar []Dateone
}

func (c *LogAnalysis) Prepare()  {
	c.Islogin()
	c.CheckPermission()
	c.GenNav()

}

func (c *LogAnalysis) Upload() {
	if c.Ctx.Input.IsGet() {
		var fdata []users.FileInfo
		qs := orm.NewOrm()
		_,err := qs.QueryTable("FileInfo").All(&fdata)
		if err != nil {
			log.Error(err)
		}
		c.Data["fdata"] = fdata
		c.Layout = "base/layout.html"
		c.TplName = "proclog/upload.html"
		return

	} else {
		f,h,err := c.GetFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			c.Redirect(beego.URLFor("LogAnalysis.Upload"),302)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		qs := orm.NewOrm()
		finfo := &users.FileInfo{
			FileName: h.Filename,
			Size: h.Size,
		}
		fid,err := qs.Insert(finfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		rder := bufio.NewReader(f)
		for {
			line,err:= rder.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
			data := strings.Split(line," ")
			code,err := strconv.Atoi(strings.Trim(data[8],"\""))
			if err != nil {
				fmt.Println(err)
				return
			}
			recorder := &users.Recorder{
				FileId: fid,
				IpAddr: strings.Trim(data[0],"\""),
				Method: strings.Trim(data[5],"\""),
				Status: code,
			}
			if _,err := qs.Insert(recorder); err != nil {
				fmt.Println(err)
				return
			}

		}
		c.Redirect(beego.URLFor("LogAnalysis.Upload"),302)
		return

	}

}


func (c *LogAnalysis) Dataapi()  {
	fid := c.Input().Get("fileid")
	if fid == "" {
		c.Ctx.WriteString("Id 不能为空")
		return
	}
	id,err := strconv.ParseInt(fid,10,64)
	if err != nil {
		c.Ctx.WriteString("非法Id")
		return
	}

	tdata := Readformdb(id)
	c.Data["json"] = &tdata
	c.ServeJSON()
}

func (c *LogAnalysis) Showweb()  {
	fid := c.Input().Get("fileid")
	if fid == "" {
		c.Ctx.WriteString("Id 不能为空")
	}
	id,err := strconv.ParseInt(fid,10,64)
	if err != nil {
		c.Ctx.WriteString("非法Id")
	}
	type sdata struct {
		Fid string
		Ipdata []Dateone
	}
	var data sdata
	data.Fid = fid
	data.Ipdata = Readformdb(id).Databar
	c.Data["data"] = data
	c.Layout = "base/layout.html"
	c.TplName = "proclog/bar.html"
}

func Readformdb(id int64) TotalData {
	var  tdata TotalData
	qs := orm.NewOrm()
	ipsql := "select ip_addr as name,count(ip_addr) as value from recorder where file_id=? group by ip_addr order by value desc limit 10"
	codesql := "select status as name,count(status) as value from recorder where file_id=? group by status"
	var ips []Dateone
	_,err := qs.Raw(ipsql,id).QueryRows(&ips)
	if err != nil {
		fmt.Println(err)
		return tdata
	}
	tdata.Databar = ips
	var codes []Dateone
	_,err = qs.Raw(codesql,id).QueryRows(&codes)
	if err != nil {
		fmt.Println(err)
		return tdata
	}
	tdata.Datapie = codes
	return tdata
}