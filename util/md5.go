package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Md5Util struct{
	Error error
}
func (this *Md5Util) Md5Str(str string) string{
	w := md5.New()
	io.WriteString(w, str)   //将str写入到w中
	md5str2  := fmt.Sprintf("%x", w.Sum(nil))
	return string(md5str2)
}