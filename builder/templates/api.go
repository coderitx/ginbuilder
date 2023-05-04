package templates

var ApiTemplate = `package apis

import (
	"{{.pkgname}}/apps/apis/hello"
)

type ApiGroup struct{
	hello.HelloApi
}

var ApiGroupApi = new(ApiGroup)
`

var HelloApiTemplate = `package hello


import (
	"{{.pkgname}}/common/responsex"
	"{{.pkgname}}/apps/service"
	"github.com/gin-gonic/gin"
)

type HelloApi struct{}

func (h *HelloApi) HelloWorld(c *gin.Context) {
	reqx := &service.HelloRequest{}
	responseData,code := reqx.HelloService()
	if code != errorx.SuccessCode{
		responsex.FailWithCode(code,"请求失败",c)
		return
	}
	responsex.OkWithData(responseData, c)
}
`
