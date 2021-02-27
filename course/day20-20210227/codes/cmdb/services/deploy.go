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

type deployerService struct {
	mutex    *sync.RWMutex
	projects map[int64]bool
}

func newDeployerService() *deployerService {
	return &deployerService{
		mutex:    &sync.RWMutex{},
		projects: make(map[int64]bool),
	}
}

var DeployerService = newDeployerService()

func (s *deployerService) Deploy(project *models.Project, branch string, packageFile string) error {
	// 检查项目相关任务是否执行(限制一个项目同时只能执行一个)
	s.mutex.RLock()
	if _, ok := s.projects[project.ID]; ok {
		s.mutex.RUnlock()
		return fmt.Errorf("running")
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

	dir, err := filepath.Abs(filepath.Join(
		beego.AppConfig.DefaultString("deployer::DeployDir", "/tmp/deploy/"),
		strconv.FormatInt(project.ID, 10),
		branch,
		now,
	))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := utils.Mkdir(dir); err != nil {
		fmt.Println(err)
		return err
	}

	deployPackageName := filepath.Base(packageFile)
	deployScriptName := "deploy.sh"
	localDeployScriptFile := filepath.Join(dir, deployScriptName)

	deployTempDir := filepath.Join(
		"/tmp/deploy2/",
		strconv.FormatInt(project.ID, 10),
		branch,
		now,
	)
	deployPackageFile := filepath.Join(deployTempDir, deployPackageName)
	deployScriptFile := filepath.Join(deployTempDir, deployScriptName)

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(`#!/bin/bash
TempDir="%s"
PackageFile="%s"

ID="%d"
Name="%s"
Namespace="%s"
Branch="%s"
Tag="%s"

cd "%s"

`, deployTempDir, deployPackageFile, project.ProjectID, project.Name, project.Namespace, branch, branch, deployTempDir))

	buffer.WriteString(project.DeployScript)
	err = ioutil.WriteFile(localDeployScriptFile, buffer.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, host := range project.Hosts {
		s.DeployHost(host, packageFile, localDeployScriptFile, deployTempDir, deployPackageFile, deployScriptFile)
	}

	return nil
}

func (s *deployerService) DeployHost(host *models.Host, localPackage, localScript, tempDir, remotePackage, remoteScript string) error {
	// 创建文件夹
	// scp 文件
	// deployscript => deploy.sh
	// scp deploy.sh
	// 执行deploy.sh
	// 清理文件

	cmds := []string{
		fmt.Sprintf(`ssh "%s"@"%s" -p %d -i "%s" "%s"`, host.User, host.Addr, host.Port, host.KeyFile, fmt.Sprintf(`mkdir -p "%s"`, tempDir)),
		fmt.Sprintf(`scp -P %d -i "%s" "%s" "%s"@"%s":"%s"`, host.Port, host.KeyFile, localPackage, host.User, host.Addr, remotePackage),
		fmt.Sprintf(`scp -P %d -i "%s" "%s" "%s"@"%s":"%s"`, host.Port, host.KeyFile, localScript, host.User, host.Addr, remoteScript),
		fmt.Sprintf(`ssh "%s"@"%s" -p %d -i "%s" "%s"`, host.User, host.Addr, host.Port, host.KeyFile, fmt.Sprintf(`bash "%s"`, remoteScript)),
		// fmt.Sprintf(`ssh "%s"@"%s" -p %d -i "%s" "%s"`, host.User, host.Addr, host.Port, host.KeyFile, fmt.Sprintf(`/bin/rm "%s"`, tempDir)),
	}

	for _, cmd := range cmds {
		fmt.Println(cmd)
		output, err := utils.Run(cmd)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}
