package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

/*
获取所有进程id,以数组返回
*/

func ProcessId() (pid []int32) {

	pids, _ := process.Pids()
	for _, p := range pids {
		pid = append(pid, p)
	}
	return pid
}

/*
获取所有进程名，以数组返回
*/

func ProcessName() (pname []string) {

	pids, _ := process.Pids()
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pName, _ := pn.Name()
		pname = append(pname, pName)
	}
	return pname
}

func main() {
	pName := ProcessName()
	for _, v := range pName {
		fmt.Println("进程名:", v)
	}
}
