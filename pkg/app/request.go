package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gpmgo/gopm/modules/log"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info(err.Key, err.Message)
	}
	return
}