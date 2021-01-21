package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
type Path struct {
	error string
}
func (this *Path) GetRootPath() (string){
	var path,_ = os.Getwd();
	return path
}
// 是否为文件
func (this *Path)IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}
func (this *Path)IsExist(f string) bool{
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func (this *Path)ListDirIncludeSubDir(dirPth string,includeSub bool) ([]string,error){
	var list = make([]string,0)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil{
		this.error = "打开目录失败"
		return list,err
	}
	for _,file := range dir{
		if !includeSub && file.IsDir(){
			continue
		}
		list = append(list, dirPth+"\\"+file.Name())
	}
	return list,nil;
}
func (this *Path)ListDir(dirPth string) ([]string,error){
	return this.ListDirIncludeSubDir(dirPth,false)
}
func (this *Path) DeleteFile(file string) (err error){
	delError  := os.Remove(file)
	return delError
}

func (this *Path) GetFileMd5(filePath string) (md5String string,err error){
	if !this.IsExist(filePath) || !this.IsFile(filePath){
		return "",errors.New("文件不存在或者类型不正确")
	}
	f, err := os.Open(filePath)
	if err != nil {
		return "",err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "复制出错",err
	}
	md5hash.Sum(nil)
	return  fmt.Sprintf("%x", md5hash.Sum(nil)),nil
}
