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
	"github.com/gin-gonic/gin"
)

type HelloApi struct{}

func (h *HelloApi) HelloWorld(c *gin.Context) {
	responsex.OkWithData("hello {{.pkgname}}", c)
}
`
