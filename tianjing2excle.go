
//还需要添加端口
package main1

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/exc/excelize"
	"io/ioutil"
	"regexp"
	"strings"
	"os"
)
var dir string
var output string
func main (){
	flag.StringVar(&dir,"p",".","报告路径")
	flag.StringVar(&output,"o","output.xlsx","结果保存文件")
	flag.Parse()
	wirte(output)
	readdirs(dir)
}
func readdirs(dir string){
	dirs,direrr := ioutil.ReadDir(dir)
	if direrr != nil{
		fmt.Println(direrr,"路径错误")
		os.Exit(1)
	}
	for _,file := range dirs {
		if strings.Contains(file.Name(), "main.html") {
			fmt.Println(strings.Replace(file.Name(),"_main.html","",-1))
			readfile(file.Name())
		}
	}
}
func readfile(filename string) {
	//打开html文件
	file, openerr := os.Open(dir + "/" + filename)
	if openerr != nil {
		fmt.Println("打开文件失败" + filename)
		os.Exit(0)
	}
	doc, docerr := goquery.NewDocumentFromReader(file)
	if docerr != nil {
		fmt.Println(docerr)
	}
	ipfile := strings.Replace(filename, "_main.html", "", -1)
	//使用Each获取标签内容,这里应该用Map,因为Map返回一个[]string
	doc.Find("#section_2_content").Find("table").Each(func(i int, selection *goquery.Selection) {
		var text []string
		text = append(text,selection.Text())
		splitstr(text,ipfile)
	})

}

func splitstr(s []string,IP string){
	var vuls []string
	var dengji,miaoshu,jianyi,url int
	for _,arrvul := range s {
		a := strings.Replace(arrvul, "\t", "", -1)
		str := strings.Split(a, "\n")
		for i, vul := range str {
			if vul == "危险级别" {
				dengji = i+4
			}
			if vul == "详细描述" {
				miaoshu = i+4
			}
			if vul == "修补建议" {
				jianyi = i+4
			}
			if vul == "参考网址" {
				url = i+4
			}

		}
		reg := regexp.MustCompile(`【\d+】`)
		biaoti := reg.ReplaceAllString(str[9], "")
		vuls = append(vuls,IP)
		vuls = append(vuls,biaoti)
		vuls = append(vuls,str[dengji])
		vuls = append(vuls,str[miaoshu])
		vuls = append(vuls,str[jianyi])
		vuls = append(vuls,str[url])
	}
	wirteok(vuls)
}
func wirte(file string){
	_, err := os.Stat(file)    //os.Stat获取文件信息
	if err != nil {
		//如果文件存在,就打开
		if os.IsExist(err) {
		}
		//如果文件不存在,就创建
		f := excelize.NewFile()
		index := f.NewSheet("漏洞列表")
		f.SetActiveSheet(index)
		if err := f.SaveAs(file); err != nil {
			fmt.Println(err)
		}
	}
}
func wirteok(results []string){
	//打开文件
	f,err := excelize.OpenFile(output)
	if err != nil{
		fmt.Println(err)
	}
	f.SetSheetRow("漏洞列表", "A2", &results)

	//插入空白行
	err1 := f.InsertRow("漏洞列表",2)
	if err1 != nil{
		fmt.Println(err1)
	}
	f.SetCellValue("漏洞列表", "A1", "IP")
	f.SetCellValue("漏洞列表", "B1", "漏洞名称")
	f.SetCellValue("漏洞列表", "C1", "漏洞等级")
	f.SetCellValue("漏洞列表", "D1", "漏洞描述")
	f.SetCellValue("漏洞列表", "E1", "修复建议")
	f.SetCellValue("漏洞列表", "F1", "参考网址")
	f.SetCellValue("漏洞列表", "G1", "端口")
	if err1 := f.Save(); err1 != nil {
		fmt.Println(err1)
	}
}

/*
type vulbox struct {
	vulname string	//漏洞名称	9
	vulid string	//漏洞编号	19
	vultype string	//漏洞类型	29
	vulgrade string	//危害等级	39
	vulos string	//影响平台	49
	cvss float32	//cvss分值	59
	bugtraq string	//bugtraq编号	69
	cve string		//cve编号	79
	cncve string	//cncve编号	87
	chinavul string	//国家漏洞库	97
	cnvd string		//cnvd编号	107
	sketch string	//简述	115
	describe string	//详细描述	125
	proposal string	//修补建议	135
	vulurl string	//参考网址	145
}

*/
