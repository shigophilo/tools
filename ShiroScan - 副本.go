package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/fatih/color"
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
	var ipFile string
	var url string
	flag.StringVar(&ipFile,"f","url.txt","urls file")
	flag.StringVar(&url,"u","http://127.0.0.1","url")
	flag.Parse()

	if len(os.Args) == 1{
		color.Red("use: " +os.Args[0] +  " " + "-f urlfile"+ "   批量检测")
		color.Green("use: " + os.Args[0] + " " + "-u url" + "   单独url检测")
		os.Exit(0)
	}else if ipFile != ""{
		array := list(ipFile)
		num := len(array) - 1
		for i,v := range array {
			color.Blue(strconv.Itoa(i) + "/" + strconv.Itoa(num) + "   "+ "Cracking: "+ v)
			ShiroScan("GET", v)
			go ShiroScan("POST",v)
		}
	} else if url != ""{
		color.Blue("Cracking: "+ url)
		ShiroScan("GET", url)
		go ShiroScan("POST",url)
	}else {
		fmt.Println("")
		color.Red("usage: " +os.Args[0] +  " " + "-f urlfile"+ "   批量检测")
		color.Green("usage: " + os.Args[0] + " " + "-u url" + "   单独url检测")
		os.Exit(0)
		os.Exit(0)
	}
	end := time.Now()
	fmt.Println("用时:", end.Sub(start), "秒")
}


func ShiroScan(metods string,url string){
	var req *http.Request
	cookie := &http.Cookie{Name:"rememberMe",Value: "1"}
	post := "rememberMe=1"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if metods == "POST"{
		req,_ = http.NewRequest("POST",url,strings.NewReader(post))
	}else{
		req,_ = http.NewRequest("GET",url,nil)
	}
	req.AddCookie(cookie)
	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(url + "-- request fail")
		return
	}
	defer resp.Body.Close()
//	fmt.Println(resp.Header)
	ok, _ := os.OpenFile("shiro.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	for _,v := range resp.Header{
		//fmt.Println(i,v)
		for _,vv := range v {
			if strings.Contains(vv, "rememberMe") {
				color.Red(metods + " : "+ url + "--" + resp.Status + "     Shiro")
				ok.Write([]byte(url + "\r\n"))
			}
		}
	}

/*
response,errread := ioutil.ReadAll(resp.Body)

	if errread != nil{
		fmt.Println(url + "--read Body fail")
	}
	str := string(response)
	fmt.Println(str)
*/
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
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}