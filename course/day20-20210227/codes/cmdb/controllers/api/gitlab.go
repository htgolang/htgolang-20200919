package api

import (
	base "cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"
	"encoding/json"
)

type GitlabController struct {
	base.ApiController
}

func (c *GitlabController) Webhook() {
	// fmt.Println(c.Ctx.Input.Context.Request.Header)
	var form forms.GitlabForm

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		if form.EventName == "tag_push" && form.Commit != "" {
			project := services.ProjectService.GetByProjectID(form.Project.ID)
			// 不存在
			if project == nil {
				// 数据库创建项目信息(auto_build=false, auto_deploy=false)
			} else {
				// 存在
			}

			if project.AutoBuild {
				// 记录，检查
				// 临时设置(原因数据库中未设置)
				project.Namespace = form.Project.Namespace
				project.Name = form.Project.Name
				project.GitHTTPURL = form.Project.GitHTTPURL
				project.GitSSHURL = form.Project.GitSSHURL

				// 构建过程
				packageFile, err := services.BuilderService.Build(project, form.TagName(), form.Commit)
				// 更新编译状态，文件

				if err == nil && project.AutoDeploy {
					// 自动部署
					services.DeployerService.Deploy(project, form.TagName(), packageFile)
				}
			}
		}
	}
	c.Data["json"] = map[string]int{
		"code": 200,
	}
	c.ServeJSON()
}
