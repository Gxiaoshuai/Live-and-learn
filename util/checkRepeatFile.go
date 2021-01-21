package util

import (
	"errors"
	"fmt"
)
type CheckRepeatFile struct{

}

func (this *CheckRepeatFile) CheckRepeat(dirPath string) (filList []string,err error){

	pathUtil := new(Path)
	fileMd5StrList := make([]string,0);
	var result bool = pathUtil.IsExist(dirPath)
	if !result{
		return fileMd5StrList,errors.New("文件夹不存在")
	}
	list,_ := pathUtil.ListDirIncludeSubDir(dirPath,true)
	fmt.Println(list)

	md5Str := ""
	for _,file := range list{
		fmt.Println(file)
		if  !pathUtil.IsFile(file) {
			tempList,err := this.CheckRepeat(file)
			if err != nil{
				fmt.Println(err)
				continue
			}
			fileMd5StrList = append(fileMd5StrList,tempList...)
			continue
		}
		md5Str,err = pathUtil.GetFileMd5(file)
		if err !=nil{
			fmt.Println(err)
			continue
		}
		for _,val := range fileMd5StrList{
			if val == md5Str{
				err := pathUtil.DeleteFile(file)
				fmt.Println(file,err)
				continue
			}
		}
		fileMd5StrList = append(fileMd5StrList,md5Str)

	}
	return fileMd5StrList,nil
}
