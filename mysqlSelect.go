package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var file string
	var out string
	flag.StringVar(&file, "f", "result.txt", "输入文件,fscan扫描结果格式即可,默认:result.txt")
	flag.StringVar(&out, "o", "", "结果保存文件 默认不保存")
	flag.Parse()
	list(file)
}

func toselect(dbuser string, dbpass string, dbip string, dbport string, db_name string) {

	sqls := map[string]string{
		"version":  "SELECT version()",
		"hostname": "SELECT @@hostname",
		"UUID":     "select UUID()",
	}
	for key, value := range sqls {
		db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbip+":"+dbport+")/"+db_name)
		if err != nil {
			fmt.Println(" -- 连接失败")
			return
		}
		defer db.Close()

		rows, err2 := db.Query(value)
		if err2 != nil {
			fmt.Println(" -- 查询失败")
			return
		}
		defer rows.Close()
		for rows.Next() {
			var out string
			err1 := rows.Scan(&out)
			if err1 != nil {
				fmt.Printf("err: %v\n", err1)
			} else {
				fmt.Printf(key+": %v\n", out)
			}
		}
	}
}
func list(file string) {
	db_file, err := os.Open(file)
	if err != nil {
		fmt.Print("文件打开失败,请确认数据库配置文件的路径,文件名是否正确!\n")
		os.Exit(0)
	}
	defer db_file.Close()
	readdb := bufio.NewReader(db_file)
	for {
		dbs, err := readdb.ReadString('\n')
		dbs = strings.Replace(dbs, "\r\n", "", -1)
		dbs = strings.Replace(dbs, "[+] mysql:", "", -1)
		dbs = strings.Replace(dbs, ":", " ", -1)
		res1 := strings.Split(dbs, " ")
		fmt.Println(res1)
		todbs(res1)
		if err == io.EOF {
			break
		}
	}
}
func todbs(dbsinfo []string) {
	toselect(dbsinfo[2], dbsinfo[3], dbsinfo[0], dbsinfo[1], "information_schema")
}
