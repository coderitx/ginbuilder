# ginbuilder
- 快速创建一个ginweb项目： 目前apps下只有http服务，如果后续有需要的话，会添加上rpc服务，websocket服务
- 后边如果有需要会添加上swagger


# 创建完成的目录结构
```text
├── apps
│  ├── apis  // 所有的apis
│  │  ├── api.go  // api处理入口文件
│  │  └── hello   // hello demo
│  │      └── hello.go
│  ├── routers
│  │ ├── hello_router.go        // 不同的路由处理位置,hello.go 为测试路由
│  │ └── init_router.go         // 路由的初始化，项目优雅启动，优雅停止
│  └── service   // 所有服务的存储位置
│     └── hello.go
├── common                 // 全局包
│  ├── errorx
│  │  └── errorx.go
│  ├── logx
│  │  └── logx.go
│  └── responsex
│      └── responsex.go
├── config                 // 配置文件
│  ├── config.go
│  ├── config.yaml
│  └── internal_config
│      ├── logger.go
│      ├── mysql.go
│      ├── redis.go
│      └── system.go
├── global                 // 公用变量
│  └── global.go
├── go.mod
├── go.sum
├── internal               // 私有依赖
│  ├── mysql.go
│  └── redis.go
├── logs                   // 日志存储位置
│  └── 2023-04-28
│      └── ginbuilder.log
└── main.go                // 项目入口
```


# 使用该工具可以快速创建ginweb服务
###  1. 完成日志的初始化 
- 使用该日志库: "go.uber.org/zap"
- 只需要修改config.yaml中的配置即可修改zap对应的配置

### 2. 封装gin路由
```golang
package routers

import (
	"{{.PkgName}}/global"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func runServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%d", global.GlobalC.System.Host, global.GlobalC.System.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.S().Infoln("Listener Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		zap.S().Infoln("timeout of 3 seconds.")
	}
	zap.S().Infoln("Server exiting")
}
```


### 3. 初始化gorm
- 使用该库: "gorm.io/gorm"

### 4. 初始化redis
- 使用该库: "github.com/go-redis/redis/v8"

### 5. 封装response的基本响应结构
```golang

package responsex

import (
	"{{.PkgName}}/common/errorx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(int(errorx.SuccessCode), data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(int(errorx.SuccessCode), data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(int(errorx.SuccessCode), map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(int(errorx.SuccessCode), map[string]any{}, "成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(int(errorx.FailedCode), data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(int(errorx.FailedCode), map[string]any{}, msg, c)
}

func FailWithCode(code errorx.ErrorCode, msg string, c *gin.Context) {
	msg, ok := errorx.ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(int(errorx.FailedCode), map[string]any{}, msg, c)
}


```


### 6. 简单封装error状态码
```golang
package errorx

type ErrorCode int

const (
	SuccessCode   ErrorCode = 1000 // 成功
	SettingsError ErrorCode = 1001 //系统错误
	ArgumentError ErrorCode = 1002 //参数错误
	FailedCode    ErrorCode = 1999 // 返回失败
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
		SuccessCode:   "成功",
		FailedCode:    "失败",
	}
)
```

# 使用方法
### 1. 安装ginbuilder

```shell
go install github.com/coderitx/ginbuilder
```

### 2. 创建项目
```shell
# 项目会创建在 $GOPATH/src 下
# 如果不指定pkg,则会默认使用app同名
ginbuilder new -project ${project-name} --pkg ${project-package-name}
```

### 3. 启动
```shell
cd ${projeck_path}
go mod tidy
go run main.go
```

### 4. 访问测试
浏览器直接访问：
- [hello](127.0.0.1:9999/api/hello)
```json
{"code":0,"data":"hello ${package name}","msg":"成功"}
```

命令行直接访问
```shell
curl 127.0.0.1:9999/api/hello

{"code":0,"data":"hello ${package name}","msg":"成功"}
```


