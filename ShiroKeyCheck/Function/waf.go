package Function

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//===================================================================================
func StrToUnicode(str string) string {
	DD := []rune(str) //需要分割的字符串内容，将它转为字符，然后取长度。
	finallStr := ""
	for i := 0; i < len(DD); i++ {
		if unicode.Is(unicode.Scripts["Han"], DD[i]) {
			textQuoted := strconv.QuoteToASCII(string(DD[i]))
			finallStr += textQuoted[1 : len(textQuoted)-1]
		} else {
			h := fmt.Sprintf("%x", DD[i])
			finallStr += "\\u" + isFullFour(h)
		}
	}
	return finallStr
}

func isFullFour(str string) string {
	if len(str) == 1 {
		str = "000" + str
	} else if len(str) == 2 {
		str = "00" + str
	} else if len(str) == 3 {
		str = "0" + str
	}
	return str
}
func ToUnicode(str string) string {
	//	fmt.Println("传进来得:" + str)
	uncodeStr := StrToUnicode(str)
	//	fmt.Println("編碼后的:" + uncodeStr)
	return uncodeStr
}

//===================================================================================
func Annotator(str string) string {
	annotatorStr := strings.Replace(str, `\u00`, `/**/\u00`, -1)

	return annotatorStr
}
