package builder

import (
	"fmt"
	"github.com/coderitx/ginbuilder/builder/templates"
	"github.com/coderitx/ginbuilder/builder/tools"
	"os"
	"path"
)

var (
	PkgPath string
	AppName string
)

func InitPkgPath(projectName, pkgpath string) {
	if projectName == "" {
		panic("appName 不能为空")
	} else {
		AppName = projectName
	}
	if pkgpath == "" {
		PkgPath = projectName
	} else {
		PkgPath = pkgpath
	}
	fmt.Printf("[INFO]: ProjectName: %v\n", projectName)
	fmt.Printf("[INFO]: Project Package Name: %v\n", PkgPath)
}

func RenderFile() error {
	gopath := tools.GetGOPath()
	projectpath := path.Join(gopath, "src")
	basepath := path.Join(projectpath, AppName)
	fmt.Println("[INFO]: Project Path: ", basepath)
	_, err := os.ReadDir(basepath)
	if err == nil || !os.IsNotExist(err) {
		fmt.Printf("[ERROR]: [%v] project already exists", basepath)
		return fmt.Errorf("[%v] project already exists", basepath)
	}
	tools.IsExistsDirectoryAndCreate(basepath)
	// go.mod
	tools.WriteToFile(path.Join(basepath, "go.mod"), fmt.Sprintf(templates.GoModTemplate, PkgPath, tools.GetGoVersion()))
	// main.go
	tools.WriteToFile(path.Join(basepath, "main.go"), tools.ReplaceAppNameTemplate(templates.MainTemplate, PkgPath))

	// config
	fs := path.Join(basepath, "config")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "config.go"), tools.ReplaceAppNameTemplate(templates.ConfigBaseTemplate, PkgPath))
	tools.WriteToFile(path.Join(fs, "config.yaml"), tools.ReplaceAppNameTemplate(templates.ConfigYamlTemplate, PkgPath))

	fs = path.Join(fs, "internal_config")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "mysql.go"), tools.ReplaceAppNameTemplate(templates.ConfigMysqlTemplate, PkgPath))
	tools.WriteToFile(path.Join(fs, "redis.go"), tools.ReplaceAppNameTemplate(templates.ConfigRedisTemplate, PkgPath))
	tools.WriteToFile(path.Join(fs, "system.go"), tools.ReplaceAppNameTemplate(templates.ConfigSysTempalte, PkgPath))
	tools.WriteToFile(path.Join(fs, "logger.go"), tools.ReplaceAppNameTemplate(templates.ConfigLoggerTemplate, PkgPath))

	// errorx
	fs = path.Join(basepath, "common/errorx")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "errorx.go"), tools.ReplaceAppNameTemplate(templates.ErrorxTempalte, PkgPath))

	fs = path.Join(basepath, "common/logx")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "logx.go"), tools.ReplaceAppNameTemplate(templates.LoggerxTemplate, PkgPath))

	fs = path.Join(basepath, "common/responsex")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "responsex.go"), tools.ReplaceAppNameTemplate(templates.ResponsexTemplate, PkgPath))

	fs = path.Join(basepath, "global")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "global.go"), tools.ReplaceAppNameTemplate(templates.GlobalFileTempalte, PkgPath))

	// internal
	fs = path.Join(basepath, "internal")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "mysql.go"), tools.ReplaceAppNameTemplate(templates.InitMysqlTemplate, PkgPath))
	tools.WriteToFile(path.Join(fs, "redis.go"), tools.ReplaceAppNameTemplate(templates.InitRedisTemplate, PkgPath))

	// router
	fs = path.Join(basepath, "apps/routers")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "init_router.go"), tools.ReplaceAppNameTemplate(templates.InitRoutersTemplate, PkgPath))
	tools.WriteToFile(path.Join(fs, "hello_router.go"), tools.ReplaceAppNameTemplate(templates.HelloRouterTemplate, PkgPath))

	// api
	fs = path.Join(basepath, "apps/apis")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "api.go"), tools.ReplaceAppNameTemplate(templates.ApiTemplate, PkgPath))
	fs = path.Join(fs, "/hello")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "hello.go"), tools.ReplaceAppNameTemplate(templates.HelloApiTemplate, PkgPath))

	// service
	fs = path.Join(basepath, "apps/service")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "hello.go"), tools.ReplaceAppNameTemplate(templates.HelloServiceTemplate, PkgPath))

	fmt.Printf("[INFO]: %v project create success \n", basepath)
	return nil
}
