package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"time"
	"web/util"
)

func main(){
	var pathTool = util.Path{}
	var rootPath = pathTool.GetRootPath()
	fmt.Println(rootPath)
	file := kvs.GetCurrentFilePath("../test_config/test.ini",1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	fmt.Println(conf.GetBoolDefault("app.enabled",false))
	var result = conf.GetDefault("mysql.mysqluser","root1")
	var port = conf.GetIntDefault("mysql.mysqlport",3306)
	var timeout = conf.GetDurationDefault("mysql.connectedTimeout",time.Second)
	fmt.Println(result,port,timeout)
}

