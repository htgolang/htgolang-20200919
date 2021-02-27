package services

import "cmdb/models"

type projectService struct{}

var ProjectService = new(projectService)

func (s *projectService) GetByProjectID(id int64) *models.Project {
	// TODO 数据库获取
	return &models.Project{
		ID:          id,
		ProjectID:   id,
		Name:        "cmdb",
		Description: "",
		GitSSHURL:   "",
		GitHTTPURL:  "",
		AutoBuild:   true,
		BuildScript: `go build .
tar vzcf ${PackageFile} cmdb
		`,
		AutoDeploy: true,
		DeployScript: `
ps -ef | grep gitlabcmdb | grep -v grep | awk '{print $2}'| xargs kill -9
mkdir -p /opt/gitlabcmdb/
tar zvxf ${PackageFile} -C /opt/gitlabcmdb/
nohup /opt/gitlabcmdb/cmdb >/dev/null 2>&1 &
		`,
		Hosts: []*models.Host{
			{
				Addr:    "localhost",
				Port:    22,
				User:    "root",
				KeyFile: "/root/.ssh/deploy",
			},
		},
	}
}
