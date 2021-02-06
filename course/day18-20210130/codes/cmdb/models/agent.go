package models

import "time"

type Agent struct {
	ID int64 `orm:"column(id);pk;auto;"`

	UUID        string `orm:"column(uuid);size(256);"`
	Addr        string `orm:"column(addr);size(256);"`
	Hostname    string `orm:"column(hostname);size(128);"`
	Description string `orm:"column(description);size(1024);"`

	Config        string `orm:"column(config);type(text);"`
	ConfigVersion int64  `orm:"column(config_version);"`

	Heartbeat *time.Time `orm:"column(heartbeat);"`

	CreatedAt *time.Time `orm:"column(created_at);auto_now_add;"`
	UpdatedAt *time.Time `orm:"column(updated_at);auto_now;"`
	DeletedAt *time.Time `orm:"column(deleted_at);null;"`
}
