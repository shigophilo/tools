package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	pingTest("192.168.1.1")
	pingTest("8.8.8.8")
	conn, err := PingConn("8.8.8.8")
	fmt.Println(conn, err)
}

func PingConn(addr string) (bool, error) {
	Command := fmt.Sprintf("ping -c 1 -W 3 %s > /dev/null && echo true || echo false", addr)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	return string(output) == "true\n", err
}

func pingTest(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Debug = true
	pinger.OnFinish = func(statistics *ping.Statistics) {
		fmt.Printf("OnFinish: %#v\n", statistics)
	}
	pinger.OnRecv = func(packet *ping.Packet) {
		fmt.Printf("OnRecv: %#v\n", packet)
	}
	pinger.Timeout = time.Second * 3
	pinger.Count = 3
	pinger.Run() // blocks until finished
}
