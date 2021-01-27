package controller

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"usermanager_beego/users"
)



type Alert struct {
	Authcontroller
}

type Alertdata struct {
	Receiver string
	Status string
	Alerts []Onealert
	GroupLabels map[string]string
	CommonLabels map[string]string
	CommonAnnotations map[string]string
	ExternalURL string
	Version string
	GroupKey string
}

type Onealert struct {
	Status string
	Labels map[string]string
	Annotations map[string]string
	StartsAt string
	EndsAt string
	GeneratorURL string
	Fingerprint string
}

func (c *Alert) Prepare()  {
	//c.Islogin()
	//c.CheckPermission()
	//c.GenNav()

}

func (c *Alert) Api()  {
	if c.Ctx.Input.IsPost() {
		str := c.Ctx.Input.RequestBody
		var onedata Alertdata
		err := json.Unmarshal(str,&onedata)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = map[string]string{"status": "500"}
			c.ServeJSON()
			return
		}
		err = procdata(onedata)
		if err != nil {
			Logger.Error("%s\n",err)
			return
		}
		c.Data["json"] = map[string]string{"status": "200"}
		c.ServeJSON()
		return
	}else {
		c.Data["json"] = map[string]string{"status": "500"}
		c.ServeJSON()
		return
	}
}

func (c *Alert) Query()  {
	c.Islogin()
	c.CheckPermission()
	c.GenNav()
	var all []users.AlertInfo
	qs := orm.NewOrm()
	str := c.Input().Get("queystr")
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)
	if str != "" {
		fmt.Println("con")
		qCond := orm.NewCondition()
		qCond = qCond.Or("Instance__icontains", str)
		qCond = qCond.Or("AlertName__icontains", str)
		qCond = qCond.Or("Summary__icontains", str)
		qCond = qCond.Or("Description__icontains", str)
		cond = cond.AndCond(qCond)
		c.Data["queystr"] = str
	}
	_,err := qs.QueryTable(&users.AlertInfo{}).SetCond(cond).All(&all)
	fmt.Println(all)
	if err != nil {
		Logger.Error("%s\n",err)
		return
	}
	c.Data["alerts"] = all
	c.Layout = "base/layout.html"
	c.TplName = "alert.html"

}

func procdata(data Alertdata) error {
	if len(data.Alerts) != 0 {
		for _,v := range data.Alerts {
			stat ,err := time.ParseInLocation(time.RFC3339Nano,v.StartsAt,time.Local)
			if err != nil {
				Logger.Error("%s\n",err)
				return err
			}
			endat, err := time.ParseInLocation(time.RFC3339Nano,v.EndsAt,time.Local)
			if err != nil {
				Logger.Error("%s\n",err)
				return err
			}
			lb ,err := json.Marshal(v.Labels)
			if err != nil {
				Logger.Error("%s\n",err)
				return err
			}
			an ,err := json.Marshal(v.Annotations)
			if err != nil {
				Logger.Error("%s\n",err)
				return err
			}
			ndata := &users.AlertInfo{
				Fingerprint: v.Fingerprint,
				Instance: v.Labels["instance"],
				Severity: v.Labels["severity"],
				AlertName: v.Labels["alertname"],
				Status: data.Status,
				StartsAt: stat,
				EndsAt: endat,
				Labels: string(lb),
				Annotations: string(an),
				Summary: v.Annotations["summary"],
				Description:v.Annotations["description"],
			}
			err = todb(ndata)
			if err != nil {
				Logger.Error("%s\n",err)
				return err
			}

		}
		return nil
	}
	return fmt.Errorf("%s\n","Alerts is null!")

}

func todb(data *users.AlertInfo) error {
	indb := &users.AlertInfo{Fingerprint: data.Fingerprint,Status: "firing"}
	qs := orm.NewOrm()
	if data.Status == "firing" {
		if err := qs.Read(indb,"fingerprint","status"); err != nil {
			if _, err := qs.Insert(data); err != nil{
				Logger.Error("%s\n",err)
				return err
			}
			return nil

		}else {
			return nil
		}

	}
	if data.Status == "resolved" {
		if err := qs.Read(indb,"fingerprint","status"); err == nil {
			indb.Status = data.Status
			indb.EndsAt = data.EndsAt
			_, err = qs.Update(indb)
			if err != nil {
				return err
			}
			return nil
		}else {
			fmt.Println(err)
			return err
		}
	}
	return fmt.Errorf("%s\n","Unkwone error")
}

