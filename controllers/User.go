package controllers

import (
	"encoding/json"
	"web/util"
)

type UserController struct{
	BaseController
}
func (this *UserController) Login(){
	var paramUtil  = util.ParamUtil{}
	var paramsMap = map[string]map[string]string{
		"password":{"from":"password","required":"1","name":"密码","default":""},
		"user_name":{"from":"user_name","required":"1","name":"用户名","default":""},
	}
	var params,err = paramUtil.GetParams(paramsMap, this.Ctx)
	msg := "ok"
	if(!err){
		msg = paramUtil.Error;
	}
	temp,_ := json.Marshal(params)
	var result = string(temp[:])
	this.ExitJson(0,msg,result,"")
}