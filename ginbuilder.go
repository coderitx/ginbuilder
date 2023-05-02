package main

import (
	"flag"
	"fmt"
	"github.com/coderitx/ginbuilder/builder"
)

func main() {
	projectName := flag.String("project", "", "gin project name")
	packageName := flag.String("pkg", "", "gin project package name")
	flag.Parse()

	if *projectName == "" && *packageName == "" {
		fmt.Println("- Please project name and package name,if the package name is empty, it defaults to the project name")
		fmt.Println("- Create project path $GOPATH/src")
		fmt.Println("--------------example----------------")
		fmt.Println("ginbuilder -project test-server -pkg github.com/conderitx/test-server")
		fmt.Println("----------------end------------------")
		return
	}

	if *projectName == "" {
		fmt.Printf("[INFO]: - please input the project name")
		panic("Project name cannot be empty")
	}
	if *packageName == "" {
		fmt.Println("[ERROR]: please input package name is empty,default package name is project name")
		*packageName = *projectName
	}
	// 初始化包名
	builder.InitPkgPath(*projectName, *packageName)
	err := builder.RenderFile()
	if err != nil {
		fmt.Errorf("[ERROR]: create project [%v] failed", *projectName)
		fmt.Errorf("[ERROR]: render file failed %v", err.Error())
		return
	}

	fmt.Println("[SUCCESS]: successful")
}
