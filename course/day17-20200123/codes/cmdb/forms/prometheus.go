package forms

type AlertForm struct {
	Alerts []struct {
		Fingerprint string
		Status      string
		Labels      map[string]string
		Annotations map[string]string
		StartsAt    string
		EndsAt      string
	} `json:"alerts"`
}
