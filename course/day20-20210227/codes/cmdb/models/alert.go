package models

import (
	"time"
)

type Alert struct {
	ID          int64      `orm:"column(id);pk;auto;"`
	Fingerprint string     `orm:"size(128);"`
	Instance    string     `orm:"size(256);"`
	AlertName   string     `orm:"size(256);"`
	Severity    string     `orm:"size(32)"`
	Status      string     `orm:"size(32);"`
	StartsAt    *time.Time `orm:"type(datetime);"`
	EndsAt      *time.Time `orm:"type(datetime);null;"`
	Summary     string     `orm:"type(text);"`
	Description string     `orm:"type(text);"`
	Labels      string     `orm:"type(longtext);"`
	Annotations string     `orm:"type(longtext);"`

	CreatedAt *time.Time `orm:"type(datetime);auto_now_add;"`
	UpdatedAt *time.Time `orm:"type(datetime);auto_now;"`
	DeletedAt *time.Time `orm:"type(datetime);null;"`
}
