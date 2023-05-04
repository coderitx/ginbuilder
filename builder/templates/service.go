package templates

var HelloServiceTemplate = `package service

import (
	"{{.pkgname}}/common/errorx"
)

type HelloRequest struct{}

type HelloResponse struct{
	Message string #json:"message"#
}

func (r *HelloRequest) HelloService() (*HelloResponse,errorx.ErrorCode){
	return &HelloResponse{
		Message: "hello {{.pkgname}}",
	},errorx.SuccessCode
}

`
