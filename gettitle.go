/*
	读取存放url的文件
	获取状态码
	状态码如果是200或者301,302 获取title
	url,状态码,title写入文件CSV格式
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main(){
	var filename string
	flag.StringVar(&filename,"f","url","url存放的文件")
	flag.Parse()
	getstatus(readfile(filename))
}

func gettitle(url string,status int){
	var utf8title string
	date,err := goquery.NewDocument(url)
	//fmt.Println(url)
	if err != nil{
		fmt.Println(url + ": 读取title失败!")
		return
	}
	title := date.Find("title").Text()
	utf8title, err = iconv.ConvertString(title, "GB18030", "utf-8")
		if err != nil {
			fmt.Println("处理字符编码失败")
		}

	fmt.Println(utf8title)
	wirter(url,status,utf8title)
}

func getstatus(urls []string) {
	for i:=0;i<len(urls);i++{
		req,err := http.Head(urls[i])
		if err != nil{
			fmt.Println(urls[i]+" 主机不存活")
			wirter(urls[i],0,"主机不存活")
			continue
		}else if req.StatusCode == 200 {
			fmt.Print(urls[i] + "  " + req.Status + " ")
			gettitle(urls[i],200)
		}else if req.StatusCode == 301 || req.StatusCode == 302 {
			fmt.Println(urls[i] + "  " + req.Status)
			gettitle(urls[i],302)

		}else{
			fmt.Println(urls[i] + "  " + req.Status)
			wirter(urls[i],req.StatusCode,"")
		}
	}
}

func wirter(url string , status int , title string){
	file, _ := os.OpenFile("结果.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	file.Write([]byte(url + "\t" + strconv.Itoa(status) + "\t" + title + "\n"))
}

func readfile(file string) []string {
	var url_list []string
	filename,err := os.Open(file)
	if err != nil {
		fmt.Println("文件打开出错")
		os.Exit(0)
	}
	defer filename.Close()
	readurl := bufio.NewReader(filename)
	for {
		urls,err := readurl.ReadString('\n')
		urls = strings.Replace(urls,"\n","",-1)
		url_list = append(url_list,urls)
		//文件末尾,退出循环
		if err != nil {
			return url_list
		}
	}

}
