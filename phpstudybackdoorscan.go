package main
import (
	"flag"
	"fmt"
	"github.com/jamiealquiza/cidrxpndr"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getipslist(ips string){
	ip, err := cidrxpndr.Expand(ips)
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	for _, i := range ip {
		gotest(i)
	}
}

func gotest(ip string){
	fmt.Println("扫描:"+ ip)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://" + ip, nil)
	if err != nil {
		return
	}
	req.Header.Set("HOST",ip)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("accept-charset","cGhwaW5mbygpOw==")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err !=nil {
		return
	}else {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	str := string(body)
	if strings.Contains(str, "www.php.net") {
		fmt.Println(ip + "存在phpstudy后门")
		ok, _ := os.OpenFile("存在漏洞.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		defer ok.Close()
		ok.Write([]byte(ip + "\n"))
	}
}

func main(){
	var ips string
	var url string
	flag.StringVar(&url,"u","","url:http(s)://shigophilo.github.io")
	flag.StringVar(&ips,"i","","ips : 127.0.0.1/24 /16 /32")
	flag.Parse()
	if (ips == "" && url == ""){
		fmt.Println("please use -u or -i")
		os.Exit(0)
	}else if (ips != "" && url != "") {
		fmt.Println("Do not use -i and --u at the same time")
		os.Exit(0)
	}else if url != ""{
		gotesturl(url)
	}else{
		getipslist(ips)
	}
}

func gotesturl(url string){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("HOST","www.baidu.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("accept-charset","cGhwaW5mbygpOw==")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err !=nil {
		fmt.Println(url+"不存活")
	}else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		str := string(body)
		if strings.Contains(str, "www.php.net") {
			fmt.Println(url + " 存在phpstudy后门")
		}else{
			fmt.Println(url + " 不存在phpstudy后门")
		}
	}

}