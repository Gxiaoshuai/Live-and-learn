package util

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)
type HttpUtil struct{

}
func (this *HttpUtil) Get(url string,params map[string]string ) (goquery.Document,error) {
	var newDoc goquery.Document
	if url == ""{
		return newDoc,errors.New("url为空")
	}
	if strings.Index(url,"?") == -1 {
		url += "?"
	}
	for key,val := range params{
		url  = url + key +"="+val
	}
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return newDoc,err
	}
	defer resp.Body.Close()
	body,err := html.Parse(resp.Body)
	newDoc = this.Html(body)
	return newDoc,err
}

func (this *HttpUtil) Decode(body,encode,decode string) string{
	srcCoder := mahonia.NewDecoder(encode)
	srcResult := srcCoder.ConvertString(body)
	tagCoder := mahonia.NewDecoder(decode)
	_, newStr, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(newStr)
	return result
}


func (this *HttpUtil) Html(body *html.Node) (goquery.Document){
	doc := goquery.NewDocumentFromNode(body)
	return *doc
}
