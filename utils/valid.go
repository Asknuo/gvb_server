package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidMsg(err error, obj any) string {
	//传入obj指针
	getobj := reflect.TypeOf(obj)
	//将err断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			if f, exits := getobj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
