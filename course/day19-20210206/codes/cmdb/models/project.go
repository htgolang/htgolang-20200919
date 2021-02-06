package models

type Project struct {
	ID           int64   `orm:"column(id);pk;auto;"`
	ProjectID    int64   `orm:"column(project_id);"`
	Name         string  `orm:"size(128);"`
	Namespace    string  `orm:"column(namespace);size(128);"`
	Description  string  `orm:"size(1024);"`
	GitSSHURL    string  `orm:"column(git_ssh_url);size(1024);"`
	GitHTTPURL   string  `orm:"column(git_http_url);size(1024);"`
	AutoBuild    bool    `orm:"column(auto_build);"`
	BuildScript  string  `orm:"column(build_script);type(text);"`
	AutoDeploy   bool    `orm:"column(auto_deploy);"`
	DeployScript string  `orm:"column(deploy_script);type(text);"`
	Hosts        []*Host `orm:"rel(m2m);"`
}
