package util

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

type ParamUtil struct{
	Error string

}
/*
	传入参数 {"password":{"from":"","required":"1","name":"","default":""},xxxx}
 */
func (this *ParamUtil) GetParams (params map[string]map[string]string, ctx *context.Context)(map[string]string,bool){
	var result = make(map[string]string)
	var paramName string
	var paramKey string
	var defaultVal string
	var value string
	for param := range params{
		paramName = params[param]["name"]
		paramKey = params[param]["from"]
		defaultVal  = params[param]["default"]
		value = defaultVal
		value = ctx.Input.Query(paramKey)
		if value =="" && params[param]["required"] == "1"{
			this.Error = fmt.Sprintf("%s参数必须传",paramName);
			return result,false
		}
		result[param] = value
	}
	return result,true
}