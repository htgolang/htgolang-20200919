package loganalysis

import (
	"beego_fileserver/process"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"strconv"
)

type LogAnalysis struct {
	beego.Controller
}

func (c *LogAnalysis) Upload()  {
	if c.Ctx.Input.IsGet() {
		Fdata := process.Getfile()
		c.Data["fdata"] = Fdata
		c.TplName = "upload.html"
	} else {
		f,h,err := c.GetFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			c.Redirect("/loganalysis/upload",302)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		id,err := process.Filemeta(h.Filename,h.Size)
		if err != nil {
			fmt.Println(err)
			return
		}
		rder := bufio.NewReader(f)
		for {
			data,err:= rder.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
			if err := process.Loginsert(id,data); err != nil {
				fmt.Println(err)
			}

		}
		c.Redirect("/loganalysis/upload",302)

	}
}

func (c *LogAnalysis) Dataapi()  {
	fid := c.Input().Get("fileid")
	if fid == "" {
		 c.Ctx.WriteString("Id 不能为空")
	}
	id,err := strconv.ParseInt(fid,10,64)
	if err != nil {
		c.Ctx.WriteString("非法Id")
	}
	tdata := process.Readformdb(id)
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
		Ipdata []map[string]string
	}
	var data sdata
	data.Fid = fid
	data.Ipdata = process.Readformdb(id).Databar
	c.Data["data"] = data
	c.TplName = "bar.html"
}