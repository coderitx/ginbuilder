package templates

var InitRoutersTemplate = `package routers

import (
	"{{.pkgname}}/global"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitApiRouter() {
	if global.GlobalC.System.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 处理路由分组
	initRouterGroup(r)
	// 处理服务监听
	runServer(r)
}

func initRouterGroup(router *gin.Engine) {
	// 定义公共前缀
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// HelloWorld
	routerGroupApp.ApiHelloGroup()
}

func runServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%d", global.GlobalC.System.Host, global.GlobalC.System.Port),
		Handler: router,
	}
	fmt.Println("[Server Address]: ", fmt.Sprintf("127.0.0.1:%d", global.GlobalC.System.Port))

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
}`

var HelloRouterTemplate = `package routers

import "{{.pkgname}}/apps/apis"

func (r *RouterGroup) ApiHelloGroup(){
	helloApi := apis.ApiGroupApi.HelloApi
	r.GET("hello",helloApi.HelloWorld)
}`
