package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"sync"
	"time"
	"web/util"
)
var httpUtil util.HttpUtil
var fileUtilHandle util.FileUtil

func main(){
	var url string = "https://www.cc148.com/44_44321/"
	var link []string = make([]string,1500)
	list := getArticle(url,"#list")
	if &list == nil{
		return
	}
	list.Find("dd a").Each(func (index int, ele *goquery.Selection){
		if index <= 8{
			return
		}
		str ,_ := ele.Html()
		str = httpUtil.Decode(str,"gbk","utf8")
		href,err := ele.Attr("href")
		if err == false{
			return
		}
		//link = append(link,url+href)
		link[index] = url+href
	})
	filePath := "link.txt"
	fileUtilHandle = getFileTool()
	str := strings.Join(link,"\n")
	err := fileUtilHandle.WriteFile(filePath,[]byte(str),0755)
	if !err {
		log.Fatal("%s","写入文件错误")
		return
	}
	var index int  = 0
	var articlePath string = "我本狂婿.txt"
	syncLock := sync.Mutex{}
	var wg sync.WaitGroup
	scapStr := string([]byte{194,160})
	for i:= 0 ;i< 1;i++{
		wg.Add(1)
		go  func(link []string,syncLock sync.Mutex,index *int){
			var content string
			for{
				if *index >= len(link){
					break
				}
				syncLock.Lock()
				url = link[*index]
				*index ++
				syncLock.Unlock()
				txt := getArticle(url,"div.box_con")
				title := txt.Find("h1").Text()
				title = httpUtil.Decode(title,"gbk","utf8")
				content,_ = txt.Find("div#content").Html()
				fmt.Println(*index,"+",url,"+",txt)
				if &txt == nil{
					continue
				}
				content = strings.Replace(content,scapStr,"", -1)
				content = httpUtil.Decode(content,"gbk","utf8")
				content = strings.Replace(content,"<br>","",-1)
				content = strings.Replace(content,"<br/><br/>","\n",-1)
				content = strings.Replace(content,"<br/>","\n",-1)
				content = strings.Replace(content,"笔趣阁手机版阅读网址：m.cc148.com"," ",-1)
				content = strings.Replace(content,"请记住本书首发域名：www.cc148.com"," ",-1)
				httpUtil = getHttpTool()
				content = title +"\n"+content
				res:= appendContentToFile(articlePath,content)
				log.Printf("%s  %s %t",url,title,res)
				time.Sleep(200 * time.Millisecond)
			}
			wg.Done()
		}(link,syncLock,&index)
	}
	wg.Wait()
}

func test(){
	var scapStr = string([]byte{194,160})
	var articlePath string = "我本狂婿.txt"
	txt := getArticle("https://www.cc148.com/44_44321/40112920.html","div.box_con")

	title := txt.Find("h1").Text()
	fmt.Println(title,"title")
	content,_ := txt.Find("div#content").Html()
	fmt.Println(title,content)
	if &txt == nil{
		return
	}
	content = strings.Replace(content,scapStr,"", -1)
	content = httpUtil.Decode(content,"gbk","utf8")
	content = strings.Replace(content,"<br>","",-1)
	content = strings.Replace(content,"<br/><br/>","\n",-1)
	content = strings.Replace(content,"<br/>","\n",-1)
	content = strings.Replace(content,"笔趣阁手机版阅读网址：m.cc148.com"," ",-1)
	content = strings.Replace(content,"请记住本书首发域名：www.cc148.com"," ",-1)
	httpUtil = getHttpTool()
	content = title +"\n"+content
	res:= appendContentToFile(articlePath,content)
	log.Printf("%s  %s %d","https://www.cc148.com/44_44321/40112920.html",content,res)
}
func appendContentToFile(filePath string, data string) bool{
	fileUtilHandle = getFileTool()
	res := fileUtilHandle.WriteFileByAppend(filePath,[]byte(data),0644)
	if res == false{
		return false
	}
	return true
}
func getArticle (url string,nodeFiler string) goquery.Selection{
	var parmas map[string]string
	httpUtil := getHttpTool()
	htmlBody,err := httpUtil.Get(url,parmas)
	var list goquery.Selection
	if err != nil{
		return list
	}
	list  = *(htmlBody.Find(nodeFiler))
	return list
}

func getHttpTool () util.HttpUtil{
	if &httpUtil == nil{
		httpUtil =  ( util.HttpUtil{})
	}
	return httpUtil
}



func getFileTool () util.FileUtil{
	if &fileUtilHandle == nil{
		fileUtilHandle =  util.FileUtil{}
	}
	return fileUtilHandle
}