package main

import (
	"fmt"
	"sync"
	"web/util"
)


var index int = 0
func main(){
	fileDir := "E:\\golang"
	channel  := make(chan string,1000)
	//fileList := make([]string,0)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	channel <- fileDir
	fileList := make([]string,2000)
	ListDir(&channel,fileList,waitGroup)
	waitGroup.Wait()
	pathUtil := new (util.Path)
	var tempMd5Str string
	var err error;
	md5List := make([]string,0)
	index = 0
	for _,val := range fileList{
		if val == ""{
			continue
		}
		tempMd5Str,err = pathUtil.GetFileMd5(val)
		if tempMd5Str == "" || err != nil{
			continue
		}
		for _,md5 := range md5List{
			if(md5 == tempMd5Str){
				DeleteFileByPath(val)
				continue
			}
		}
		md5List = append(md5List, tempMd5Str)
	}
	fmt.Println(md5List)
}
func ListDir(ch *chan string,fileList []string,waitGroup *sync.WaitGroup){
	pathUtil := new (util.Path)

	defer waitGroup.Done()
	select {
		case dir := <- *ch:
			list,_ := pathUtil.ListDirIncludeSubDir(dir,true)
			for _,file := range list{
				if !pathUtil.IsFile(file){
					*ch <- file
					waitGroup.Add(1)
					go ListDir(ch,fileList,waitGroup)
				}else{
					//fileSliece := []string{dir+"\\"+file}
					fileList[index] = file
					index++
				}
			}
	}
}
func DeleteFileByPath(filePath string)(result bool){
	fmt.Println("delete file "+filePath)
	pathUtil := new (util.Path)
	err := pathUtil.DeleteFile(filePath)
	if(err != nil){
		fmt.Println(err)
		return  false
	}
	return true
}