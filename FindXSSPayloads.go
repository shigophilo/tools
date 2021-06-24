package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"net/http"
	"crypto/tls"
	"io/ioutil"
)
var out string
func main() {
	start()
	var payloadsfile string
	var cookie string
	var url string
	flag.StringVar(&url,"u","","url")
	flag.StringVar(&payloadsfile,"p","payloads.txt","payloads file")
	flag.StringVar(&cookie,"c","","cookie值,多个值请用\"\"(双引号)括起来")
	flag.StringVar(&out,"o","output.txt","保存文件名 txt格式")
	flag.Parse()
	payloads_list := read_payloads(payloadsfile)
	scan(url,cookie,payloads_list)
}


func scan(url string,cookie string,payloads []string) {
	var j = 0
	var urll = ""
		for i,poc := range payloads {
			urll = strings.Replace(url,"*",poc,-1)
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			req, err := http.NewRequest("GET", strings.Replace(urll,"*",poc,-1),nil)
			if err != nil {
				fmt.Println("======失败!" + ":" + url + ":" + err.Error())
				break
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			if cookie != ""{
				req.Header.Set("cookie",cookie)
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("主机可能不存活")
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取失败")
			}
			str := string(body)
			if resp.StatusCode == 200 && strings.Contains(str, poc) {
				j++
				fmt.Printf("正在扫描第%d / %d 个=====已匹配到:%d个\n",i,len(payloads),j)
				wirte(poc)
			}
		}
}

func wirte(poc string){
	ok, _ := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ok.Write([]byte(poc + "\n"))
}
func read_payloads(payloadsfile string) []string {
	var payloads_list []string
	payloads_file, err := os.Open(payloadsfile)
	if err != nil {
		fmt.Print("No Payloads file\n")
		os.Exit(0)
	}
	defer payloads_file.Close()
	readerpayloads := bufio.NewReader(payloads_file)
	for {
		payloads, err := readerpayloads.ReadString('\n')
		payloads = strings.Replace(payloads, "\n", "", -1)
		payloads = strings.Replace(payloads, "\r", "", -1)
		payloads_list = append(payloads_list, payloads)
		if err == io.EOF {
			break
		}
	}
	return payloads_list
}
func start() {
	fmt.Println("================================================================================")
	fmt.Println("-u  url   需要测试的参数值用\"*\"标注 如:http://127.0.0.1/index.php?id=*")
	fmt.Println("-p  可选  payload文件   默认:payloads.txt\n-c  可选  cookie值     多个值请用\"\"(双引号)括起来")
	fmt.Println("-o  可选  output      保存结果(可用payload)的文件名 txt格式  默认 output.txt")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("应该手工先测一下,如将id的值设置为\"11111\",返回包中若匹配到\"11111\"便可以用此脚本跑")
	fmt.Println("======================================================== mrhonest == 20200817 ==")
}