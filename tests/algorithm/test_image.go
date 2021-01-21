package main

import (
	"fmt"
	"os"
	"web/util"
)

func main (){
	var pathUtil = util.Path{}
	var path = pathUtil.GetRootPath()+string(os.PathSeparator)+"tests"+string(os.PathSeparator)+"img";
	fmt.Println(path)
	pathList,err := pathUtil.ListDir(path)
	if err!=nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	var imageUtil = new (util.ImageUtil)
	for _,file := range pathList{
		imageUtil.RepeatImage(path+"\\"+file,path+"\\re"+file,2,2)
	}
}