package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("start")
	var logFile string
	var url string
	var sleeps int
	flag.StringVar(&logFile, "f", "name.txt", "name file")
	flag.StringVar(&url, "u", "http://127.0.0.1/Application/Runtime/Logs/Home/", "url")
	flag.IntVar(&sleeps, "s", 0, "延时:秒")
	flag.Parse()
	names := list(logFile)
	Httpget(url, names, sleeps)
}
func Httpget(url string, names []string, sleeps int) {
	for _, name := range names {
		fmt.Println("搜索日志 : " + name)
		url := url + name + ".log"
		cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		request.Header.Add("Connection", "close")

		do, err := cli.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer func() {
			_ = do.Body.Close()
		}()
		all, _ := ioutil.ReadAll(do.Body)
		html := string(all)
		//fmt.Println(html)
		if sleeps != 0 {
			time.Sleep(time.Duration(sleeps) * time.Second)
		}
		if strings.Contains(html, "sql") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + ".log     sql \r\n"))
			defer ok.Close()
		}
		if strings.Contains(html, "select") || strings.Contains(html, "SELECT") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + ".log     select\r\n"))
			defer ok.Close()
		}
		if strings.Contains(html, "username") || strings.Contains(html, "USERNAME") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + ".log     username\r\n"))
			defer ok.Close()
		}
		if strings.Contains(html, "password") || strings.Contains(html, "PASSWORD") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + ".log         password\r\n"))
			defer ok.Close()
		}

		if strings.Contains(html, "ADMIN") || strings.Contains(html, "admin") || strings.Contains(html, "token") || strings.Contains(html, "TOKEN") || strings.Contains(html, "管理") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + "    admin.token.管理\r\n"))
			defer ok.Close()
		}
	}
}
func list(urlfile string) []string {
	var url_list []string
	url_file, err := os.Open(urlfile)
	if err != nil {
		fmt.Println("Can't open urlfile")
		os.Exit(0)
	}
	defer url_file.Close()
	reader_Url := bufio.NewReader(url_file)
	for {
		url, err := reader_Url.ReadString('\n')
		url = strings.Replace(url, " ", "", -1)
		url = strings.Replace(url, "\n", "", -1)
		url = strings.Replace(url, "\r", "", -1)
		url_list = append(url_list, url)
		if err == io.EOF {
			break
		}
	}
	return url_list
}
