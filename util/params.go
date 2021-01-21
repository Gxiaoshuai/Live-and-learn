package util

import (
	"errors"
	"github.com/astaxie/beego/context"
)

type ParamUtil struct{
	Error error

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
			this.Error = errors.New(paramName+"参数必须传");
			return result,false
		}
		result[param] = value
	}
	return result,true
}
///*
//	传入参数 {"roleId":{"from":"roleId","required":"1","name":"技能组id","default":0,"type":"int8"},xxxx}
//*/
//func (this *ParamUtil) GetParamsInt(params map[string]map[string]interface{},ctx *context.Context)(map[string]string,bool){
//	var result = make(map[string]interface{})
//	if ctx.Request.Form == nil {
//		ctx.Request.ParseForm()
//	}
//	var input = ctx.Request.Form
//	input.
//}