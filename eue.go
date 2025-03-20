package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var OUTPUT string

func main() {
	// SMTP服务器地址和端口,用户名,域名后缀
	var smtpServer, smtpPort, usernameFile, domain string
	var concurrency int
	flag.StringVar(&smtpServer, "s", "", "SMTP服务器地址")
	flag.StringVar(&smtpPort, "p", "25", "SMTP端口")
	flag.StringVar(&usernameFile, "u", "users.txt", "用户名文件")
	flag.StringVar(&domain, "d", "", "域名")
	flag.IntVar(&concurrency, "c", 10, "并发数量")
	flag.Parse()
	OUTPUT = domain + ".txt"
	// 读取文件中的用户名
	usernames, err := readURLsFromFile(usernameFile)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}

	// 控制并发数量的通道
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	for _, name := range usernames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			stmpEnumerate(smtpServer, smtpPort, name, domain)
		}(name)
	}
	wg.Wait()
}

func stmpEnumerate(smtpServer, smtpPort, username, domain string) {
	// 建立SMTP连接
	conn, err := net.Dial("tcp", smtpServer+":"+smtpPort)
	if err != nil {
		//log.Printf("Failed to connect to SMTP server: %v", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// 读取初始响应
	_, err = reader.ReadString('\n')
	if err != nil {
		//log.Printf("Failed to read initial response: %v", err)
		return
	}
	//fmt.Println("Initial response:", initialResponse)
	// 发送HELO命令
	_, err = fmt.Fprintf(conn, "HELO %s\r\n", domain)
	if err != nil {
		//log.Printf("Failed to send HELO command: %v", err)
		return
	}
	// 读取HELO响应
	_, err = reader.ReadString('\n')
	if err != nil {
		//log.Printf("Failed to read HELO response: %v", err)
		return
	}
	//fmt.Println("HELO response:", heloResponse)
	// 发送MAIL FROM命令
	_, err = fmt.Fprintf(conn, "MAIL FROM:<test@example.com>\r\n")
	if err != nil {
		//log.Printf("Failed to send MAIL FROM command: %v", err)
		return
	}
	// 读取MAIL FROM响应
	_, err = reader.ReadString('\n')
	if err != nil {
		//log.Printf("Failed to read MAIL FROM response: %v", err)
		return
	}
	//fmt.Println("MAIL FROM response:", mailFromResponse)
	// 发送RCPT TO命令
	_, err = fmt.Fprintf(conn, "RCPT TO:<%s@%s>\r\n", username, domain)
	if err != nil {
		//log.Printf("Failed to send RCPT TO command: %v", err)
		return
	}
	// 读取RCPT TO响应
	rcptToResponse, err := reader.ReadString('\n')
	if err != nil {
		//log.Printf("Failed to read RCPT TO response: %v", err)
		return
	}
	//fmt.Println("RCPT TO response:", rcptToResponse)
	// 分析RCPT TO响应
	if strings.HasPrefix(rcptToResponse, "250") {
		str := username + "@" + domain
		fmt.Println(str)
		ok, _ := os.OpenFile(OUTPUT, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		defer ok.Close()
		ok.Write([]byte(str + "\r\n"))
	}
	// } else if strings.HasPrefix(rcptToResponse, "550") {
	// 	fmt.Printf("User %s does not exist\n", username)
	// } else {
	// 	fmt.Printf("Unexpected response for user %s: %s", username, rcptToResponse)
	// }

	// 发送QUIT命令
	_, err = fmt.Fprintf(conn, "QUIT\r\n")
	if err != nil {
		//log.Printf("Failed to send QUIT command: %v", err)
		return
	}
}

// 从文件中读取用户名
func readURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var usernames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		usernames = append(usernames, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return usernames, nil
}

func init() {
	fmt.Println(`                                           `)
	fmt.Println(` /$$$$$$$$      /$$   /$$         /$$$$$$$$`)
	fmt.Println(`| $$_____/     | $$  | $$        | $$_____/`)
	fmt.Println(`| $$           | $$  | $$        | $$      `)
	fmt.Println(`| $$$$$ /$$$$$$| $$  | $$ /$$$$$$| $$$$$   `)
	fmt.Println(`| $$__/|______/| $$  | $$|______/| $$__/   `)
	fmt.Println(`| $$           | $$  | $$        | $$      `)
	fmt.Println(`| $$$$$$$$     |  $$$$$$/        | $$$$$$$$`)
	fmt.Println(`|________/      \______/         |________/`)
	fmt.Println(`   Email          User           Enumeration`)
}
