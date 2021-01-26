package services

import (
	"cmdb/models"

	"github.com/astaxie/beego/orm"
)

type alertService struct{}

var AlertService = new(alertService)

func (s *alertService) Notify(alert *models.Alert) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Alert{})
	queryset = queryset.Filter("Fingerprint", alert.Fingerprint)
	queryset = queryset.Filter("DeletedAt__isnull", true)
	queryset = queryset.Filter("Status", "firing")

	if alert.Status == "firing" {
		if cnt, err := queryset.Count(); err == nil && cnt == 0 {
			ormer.Insert(alert)
		}
	} else {
		queryset.Update(orm.Params{
			"EndsAt": alert.EndsAt,
			"Status": alert.Status,
		})
	}
}

func (s *alertService) Query(q string) []*models.Alert {
	var alerts []*models.Alert
	queryset := orm.NewOrm().QueryTable(&models.Alert{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)
	if q != "" {
		qCond := orm.NewCondition()
		qCond = qCond.Or("Instance__icontains", q)
		qCond = qCond.Or("AlertName__icontains", q)
		qCond = qCond.Or("Summary__icontains", q)
		qCond = qCond.Or("Description__icontains", q)
		cond = cond.AndCond(qCond)
	}

	queryset.SetCond(cond).All(&alerts)
	return alerts
}
