package main

import (
	"encoding/json"
	"fmt"
)

type AlertForm struct {
	Alerts []struct {
		Status      string
		Labels      map[string]string
		Annotations map[string]string
		StartsAt    string
		EndsAt      string
	} `json:"alerts"`
}

func main() {
	txt := `
	{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"target is down","instance":"localhost:9100","job":"node","severity":"high"},"annotations":{"description":"Node localhost:9100 is Down","summary":"Nodelocalhost:9100 is Down"},"startsAt":"2021-01-23T10:04:10.139169689Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://centos:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"8137000e5cec9c89"}],"groupLabels":{"alertname":"target is down"},"commonLabels":{"alertname":"target is down","instance":"localhost:9100","job":"node","severity":"high"},"commonAnnotations":{"description":"Node localhost:9100 is Down","summary":"Node localhost:9100 is Down"},"externalURL":"http://centos:9093","version":"4","groupKey":"{}:{alertname=\"target is down\"}","truncatedAlerts":0}
	`
	form := AlertForm{}
	json.Unmarshal([]byte(txt), &form)
	for _, alert := range form.Alerts {
		fmt.Println(alert)
	}

}
