package templates

var HelloServiceTemplate = `package service

import (
	"{{.pkgname}}/common/errorx"
)

func HelloService() errorx.ErrorCode{
	return errorx.SuccessCode
}

`
