package services

import (
	"bytes"
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

type builderService struct {
	mutex    *sync.RWMutex
	projects map[int64]bool
}

func newBuilderService() *builderService {
	return &builderService{
		mutex:    &sync.RWMutex{},
		projects: make(map[int64]bool),
	}
}

var BuilderService = newBuilderService()

func (s *builderService) Build(project *models.Project, branch string, commit string) (string, error) {
	// 检查项目相关任务是否执行(限制一个项目同时只能执行一个)
	s.mutex.RLock()
	if _, ok := s.projects[project.ID]; ok {
		s.mutex.RUnlock()
		return "", fmt.Errorf("running")
	}
	s.mutex.RUnlock()

	// 标记项目正在执行
	s.projects[project.ID] = true
	defer func() {
		s.mutex.Lock()
		delete(s.projects, project.ID)
		s.mutex.Unlock()
	}()

	now := time.Now().Format("2006-01-02_15-04-05")
	// /tmp/xxxx/{ID}/{branch}/{now}

	dir, err := filepath.Abs(filepath.Join(
		beego.AppConfig.DefaultString("builder::BuildDir", "/tmp/build/"),
		strconv.FormatInt(project.ID, 10),
		branch,
		now,
	))
	if err != nil {
		return "", err
	}

	// 密钥处理
	cmds := []string{
		fmt.Sprintf(`mkdir -p "%s"`, dir),
		fmt.Sprintf(`cd "%s" && git clone "%s" "%s"`, dir, project.GitSSHURL, branch),
		fmt.Sprintf(`cd "%s/%s" && git checkout -b "%s"`, dir, branch, branch),
		fmt.Sprintf(`cd "%s/%s" && git reset --hard "%s"`, dir, branch, commit),
	}
	for _, cmd := range cmds {
		fmt.Println("================")
		fmt.Println(cmd)
		output, err := utils.Run(cmd)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		fmt.Println(string(output))
	}

	// tag.gz
	// dir/build.sh
	packageName := fmt.Sprintf("%s_%s.tar.gz", branch, now)
	packageFile := filepath.Join(dir, packageName)
	scriptFile := filepath.Join(dir, "build.sh")

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(`#!/bin/bash
BuildDir="%s/%s"
PackageFile="%s"

ID="%d"
Name="%s"
Namespace="%s"
Branch="%s"
Tag="%s"

cd "%s/%s"

`, dir, branch, packageFile, project.ProjectID, project.Name, project.Namespace, branch, branch, dir, branch))

	buffer.WriteString(project.BuildScript)

	err = ioutil.WriteFile(scriptFile, buffer.Bytes(), os.ModePerm)
	if err != nil {
		return "", err
	}
	output, err := utils.RunFile(scriptFile)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(output))
	// 编译
	// build.sh
	// cd /dir
	// buildDir=
	// packageFile=xxxx
	// ProjectName
	// ProjectID
	// TagName

	// ${packageFile}
	//
	// project.BuildScript // 写入 => sh
	// bash -c "path"

	// 检查packageFile 是否存在

	// 移动到发布目录
	packageDir, err := filepath.Abs(beego.AppConfig.DefaultString("builder::PackageDir", "/tmp/package/"))
	if err != nil {
		return "", nil
	}
	err = utils.Mkdir(packageDir)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	distPackageFile := filepath.Join(packageDir, packageName)

	err = os.Rename(packageFile, distPackageFile)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(distPackageFile)
	return distPackageFile, nil
}
