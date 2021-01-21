package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type FileUtil struct{

}
var pathUtil Path
func (this *FileUtil) WriteFile(filePath string,data []byte,mode os.FileMode) bool{
	err := ioutil.WriteFile(filePath,data,mode)
	if err != nil{
		log.Fatal("%s","文件写入错误:"+err.Error())
		return false
	}
	return true
}
func (this *FileUtil) WriteFileByAppend(filePath string,data []byte,mode os.FileMode)bool{
	pathUtilHandle := getPathTool()
	var f *os.File
	if !pathUtilHandle.IsExist(filePath){
		f,_ = os.Create(filePath)
	}else{
		f,_ = os.OpenFile(filePath, os.O_APPEND, 0644)
	}
	defer f.Close()
	_,err := io.WriteString(f,string(data))
	if err == nil{
		return true
	}
	return false
}
func (this *FileUtil) ReadFile (filePath string ) []byte{
	data,err := ioutil.ReadFile(filePath)
	if err == nil{
		log.Fatal("%s","读取文件错误:"+err.Error())
	}
	return data;
}

func getPathTool () Path{
	if &pathUtil == nil{
		pathUtil =  ( Path{})
	}
	return pathUtil
}