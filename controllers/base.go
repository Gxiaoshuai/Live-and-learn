package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ExitJson(code int,msg string,jsonResult string,other interface{}) {
	var otherString,_ = json.Marshal(other)
	var resultMap = map[string]string{
		"code":strconv.Itoa(code),
		"data":jsonResult,
		"msg":msg,
		"other":string(otherString[:]),
	}
	result,_ := json.Marshal(resultMap);
	this.Ctx.ResponseWriter.Write(result)
}

