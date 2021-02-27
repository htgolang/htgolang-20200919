package models

type Host struct {
	ID      int64  `orm:"column(id)"`
	Addr    string `orm:"column(addr);size(256);"`
	Port    int    `orm:"column(port);"`
	User    string `orm:"column(user);size(64);"`
	KeyFile string `orm:"column(key_file);size(1024);"`

	Projects []*Project `orm:"reverse(many);"`
}
