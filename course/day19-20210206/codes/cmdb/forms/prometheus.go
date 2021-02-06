package forms

import (
	"cmdb/models"
	"encoding/json"
	"time"
)

type AlertForm struct {
	Fingerprint string
	Status      string
	Labels      map[string]string
	Annotations map[string]string
	StartsAt    *time.Time
	EndsAt      *time.Time
}

func (f *AlertForm) ToModel() *models.Alert {
	var (
		endsAt      *time.Time
		labels      []byte
		annotations []byte
	)
	if f.EndsAt != nil && !f.EndsAt.IsZero() {
		endsAt = f.EndsAt
	}

	labels, _ = json.Marshal(f.Labels)
	annotations, _ = json.Marshal(f.Annotations)

	return &models.Alert{
		Fingerprint: f.Fingerprint,
		Instance:    f.Labels["instance"],
		AlertName:   f.Labels["alertname"],
		Severity:    f.Labels["severity"],
		Status:      f.Status,
		StartsAt:    f.StartsAt,
		EndsAt:      endsAt,
		Summary:     f.Annotations["summary"],
		Description: f.Annotations["description"],
		Labels:      string(labels),
		Annotations: string(annotations),
	}
}

type AlertsForm struct {
	Alerts []AlertForm `json:"alerts"`
}
