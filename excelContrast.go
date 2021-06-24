package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
	"github.com/exc/excelize"
)

//var module string
var exclude []string
var assessment []string
/*
var yingyongvuls []string
var zhujivuls []string
var weifenleivuls []string

*/
func main() {
	start := time.Now()
	printlogo()
	var fileA string
	var fileB string
	var excludefile string
	var assessmentfile string
	//	var vulfile string
	flag.StringVar(&fileA, "new", "a.xlsx", "文件A")
	flag.StringVar(&fileB, "old", "b.xlsx", "文件B")
	flag.StringVar(&assessmentfile, "k", "assessment.txt", "需要考核的漏洞所在文档名(一行一个漏洞名)")
	flag.StringVar(&excludefile, "e", "exclude.txt", "需要排除的IP(不计漏洞)")
	flag.Parse()
	//	flag.StringVar(&vulfile,"v","vuls.txt","漏洞分类文件-只针对总公司分类")
	excludeip(excludefile)
	assessmentarray(assessmentfile)
	//	vultoarray(vulfile)
	f1array := excletoarray(fileA)
	f2array := excletoarray(fileB)
	contrast(&f1array, &f2array)
	end := time.Now()
	fmt.Println("用时:", end.Sub(start), "秒")
}

/*
//漏洞分类
func vultoarray(filename string) {
	vul_file, err := os.Open(filename)
	if err != nil {
		fmt.Print("文件打开失败,请确认存放漏洞分类文件的路径,文件名是否正确!\n")
		os.Exit(0)
	}
	defer vul_file.Close()
	readervul := bufio.NewReader(vul_file)

	for {
		ip, err := readervul.ReadString('\n')
		ip = strings.Replace(ip, " ", "", -1)
		ip = strings.Replace(ip, "\n", "", -1)
		ip = strings.Replace(ip, "\r", "", -1)
		if
		exclude = append(exclude, ip)
		if err == io.EOF {
			break
		}
	}
}
*/
//获取要排除的IP
func excludeip(excludefile string) {
	ip_file, err := os.Open(excludefile)
	if err != nil {
		fmt.Print("文件打开失败,请确认存放IP文件的路径,文件名是否正确!\n本次将没有排除的IP\n")
		//os.Exit(0)
	}
	defer ip_file.Close()
	readerip := bufio.NewReader(ip_file)
	for {
		ip, err := readerip.ReadString('\n')
		ip = strings.Replace(ip, " ", "", -1)
		ip = strings.Replace(ip, "\n", "", -1)
		ip = strings.Replace(ip, "\r", "", -1)
		exclude = append(exclude, ip)
		if err == io.EOF {
			break
		}
	}
	//fmt.Println(exclude) //不展示排除的IP
}

//写入结果 filename是公司名.xlsx  results是传过来的漏洞列表数组
func wirteok(filename string, results []string) {
	var company = strings.Split(filename, ".")[0]
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("正在写入" + company)
	f.SetSheetRow("漏洞列表", "A2", &results)
	//插入空白行
	err1 := f.InsertRow("漏洞列表", 2)
	if err1 != nil {
		fmt.Println(err1)
	}
	f.SetCellValue("漏洞列表", "A1", "IP")
	f.SetCellValue("漏洞列表", "B1", "漏洞名称")
	f.SetCellValue("漏洞列表", "C1", "漏洞等级")
	f.SetCellValue("漏洞列表", "D1", "漏洞地址")
	f.SetCellValue("漏洞列表", "E1", "漏洞描述")
	f.SetCellValue("漏洞列表", "F1", "漏洞危害")
	f.SetCellValue("漏洞列表", "G1", "修复建议")
	f.SetCellValue("漏洞列表", "H1", "是否考核")
	if err1 := f.Save(); err1 != nil {
		fmt.Println(err1)
	}
}

//判断即将写入的文件是否存在
func wirte(vul []string, company string) {
	var filename = company + ".xlsx"
	//判断文件(夹)是否存在
	_, err := os.Stat(filename) //os.Stat获取文件信息
	if err != nil {
		//如果文件存在,就打开
		if os.IsExist(err) {
		}
		//如果文件不存在,就创建
		f := excelize.NewFile()
		index := f.NewSheet("漏洞列表")
		f.SetActiveSheet(index)
		if err := f.SaveAs(filename); err != nil {
			fmt.Println(err)
		}
	}
	wirteok(filename, vul)
}

//比对
func contrast(f1arr *[][]string, f2arr *[][]string) {
	var results [][]string
	for _, vul := range *f1arr {
		var ok = true
		for _, vul1 := range *f2arr {
			if reflect.DeepEqual(vul, vul1) {
				ok = false
				break
			}
		}
		if ok {
			//vul[1] = vul[1] + "\n" + "aaaaaaaaaa"
			results = append(results, vul)
		}
	}
	//fmt.Println("共获取", len(results), "条数据")
	biduivul(results)
}
func biduivul(vuls [][]string){
	var vulsok [][]string
	for _,vul := range vuls{
	for _,vulname :=range assessment{
		if reflect.DeepEqual(vul[1], vulname){
			vul = append(vul,"是")
		}
	}
	vulsok = append(vulsok,vul)
	}
	output(vulsok)
}
//将excle内容转成array,并去掉需要排除的IP
func excletoarray(file string) [][]string {
	var rows [][]string
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	row, _ := f.GetRows("漏洞列表")
	for _, vul := range row {
		var ok = true
		for _, excludeip := range exclude {
			if reflect.DeepEqual(vul[0], excludeip) {
				ok = false
				break
			}
		}
		if ok {
			rows = append(rows, vul[:7])
		}
	}
	return rows
}
//判断是否为考核的漏洞,如果是考核的漏洞就在导出的时候添加颜色
//添加颜色太难了,改成加个字段
func assessmentarray(assessmentfile string){
//获取考核文件中的漏洞名
	assessment_file,err := os.Open(assessmentfile)
	if err != nil{
		fmt.Print("没有考核的漏洞文件文档!不影响程序运行!\n")
		return
	}
	defer assessment_file.Close()
	readassessmentfile := bufio.NewReader(assessment_file)
	for {
		vul,err := readassessmentfile.ReadString('\n')
		vul = strings.Replace(vul, " ", "", -1)
		vul = strings.Replace(vul, "\n", "", -1)
		vul = strings.Replace(vul, "\r", "", -1)
		assessment = append(assessment, vul)
		if err == io.EOF {
			break
		}
	}
}

func printlogo() {
	fmt.Println("<---------- 2021-2-4 新增判断考核漏洞 ---------->")
	fmt.Println("两个excle文件内容对比  -new 新文件   --old 老文件 (比对老文件中没有新文件里的内容)")
	fmt.Println("<--- 也就是第二次扫描比第一次扫描多出来的漏洞(数据)--->")
	fmt.Println("<--- 如果把新文件和老文件掉过来比对,也就是找出第二次扫描比第一次扫描已经修复的漏洞--->")
	fmt.Println("<--- ShiGoPhilo for PICC  2021-01-22--->")
}

//判断归属
func output(results [][]string) {
	for _, vul := range results {
		switch vul[0][0:2] {
		case "10":
			wirte(vul, "总公司")
			break
		case "11":
			wirte(vul, "总公司")
			break
		case "12":
			wirte(vul, "北京")
			break
		case "14":
			wirte(vul, "天津")
			break
		case "16":
			wirte(vul, "河北")
			break
		case "18":
			wirte(vul, "山西")
			break
		case "20":
			wirte(vul, "内蒙古")
			break
		case "22":
			wirte(vul, "辽宁")
			break
		case "24":
			wirte(vul, "大连")
			break
		case "26":
			wirte(vul, "吉林")
			break
		case "28":
			wirte(vul, "黑龙江")
			break
		case "30":
			wirte(vul, "上海")
			break
		case "32":
			wirte(vul, "江苏")
			break
		case "34":
			wirte(vul, "浙江")
			break
		case "35":
			wirte(vul, "浙江")
			break
		case "36":
			wirte(vul, "宁波")
			break
		case "38":
			wirte(vul, "安徽")
			break
		case "40":
			wirte(vul, "福建")
			break
		case "42":
			wirte(vul, "厦门")
			break
		case "44":
			wirte(vul, "江西")
			break
		case "46":
			wirte(vul, "山东")
			break
		case "48":
			wirte(vul, "青岛")
			break
		case "50":
			wirte(vul, "河南")
			break
		case "52":
			wirte(vul, "湖北")
			break
		case "54":
			wirte(vul, "湖南")
			break
		case "56":
			wirte(vul, "广东")
			break
		case "58":
			wirte(vul, "深圳")
			break
		case "60":
			wirte(vul, "广西")
			break
		case "62":
			wirte(vul, "海南")
			break
		case "64":
			wirte(vul, "四川")
			break
		case "66":
			wirte(vul, "重庆")
			break
		case "68":
			wirte(vul, "贵州")
			break
		case "70":
			wirte(vul, "云南")
			break
		case "72":
			wirte(vul, "西藏")
			break
		case "74":
			wirte(vul, "陕西")
			break
		case "76":
			wirte(vul, "甘肃")
			break
		case "78":
			wirte(vul, "青海")
			break
		case "80":
			wirte(vul, "宁夏")
			break
		case "82":
			wirte(vul, "新疆")
			break
		default:
			wirte(vul, "资产不明")
		}
	}
}
