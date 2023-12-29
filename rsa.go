package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {

	fmt.Println(RsaDecode("kTGcpW783n7J4/dO8514pA=="))
}
func RsaDecode(str string) string {

	plainText, err := base64.StdEncoding.DecodeString(str)
	//私钥
	privateKey :=
		`-----BEGIN PUBLIC KEY-----
		MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCl7oKSljmXRR1BRcytdXeSIQk4UVCcPUsw/QAkLniEs+cjehe0AraBYTfDI/gzFJyCpS0hEcIplWaWRS4lHWeJWIUj+aEWg2UHUoluxhPeOMCZAYoMFA7o0EivX/MeVMmLg2tJ0+d2UljqfADjUxKIm8zi/p2czzD9L8qYjji0CMMrDFMSXoW93qZw2Yvv/iavyWBV32jvevrEw19qmxuUhdsOCAWwgULqTG1wl3bIRgIjCAiJ3q5wtX6ouZie+QX3bM31shfzx2IwA7o5W2NIXMcqZV/los5g36hnGFzg1br/0jwS+8oi52AZoBQ+jQVDPGg2ztnYxdiQ/TUlaC4HAgMBAAECggEAPZ+ViIEYBFDbq4kTmXvdmzx/oxW33T56NUhQjB9iDM6PUcKfGPBE/Umwnt016SqHcZGMcBKbTDw67CXnHEuOTxi74X9szuKfW5fQiH4xlL748Bn9Pb0ZSGdKBy+YJ7SPYSu+ZZ24AV5VvUeuQL5yTJ+n+MlcmYbtB6p/ECcJcfqPE0MSKJRsoxXq1DKhyQObEqnOuaPw3/JVCuADW+D541bw2VcaJxUyIk3IvnUfQL7koDZWmkWRLFBhVs7kKhWZjcUhGhSPcmrzUU9RJoEI0+2vHOIoBctbxuL1v3uWSIqi8OVx6+uSwd+lh10fycEyfW/su/g0qQTqjx1qs7WRkQKBgQD6m99njtQMLCjQWL6yrxivWYi5ZiBGrUVBlbZWSNZAtEIblA2XxIf6umZj4XjZwBOb0FfKRgufq4f4mRONhf6zturOB9GE3+w95BTwA78AXbvV089sYocp61TxYQzuqK2p+Nmp7I2MoGs0xIqY6FL3HkODs4671Vz1JwqyD3wlLwKBgQCpgE/V7IFijNsKKwWnpLLPYY+OBmKUZEaa47WN3xdZ039aHwH557s2X+zSLcFEsSBjiNtmozNmF700XXWGDez/qPSzcHvpKGzcv/DTV0xkjBYtUSUilMY9boohOF7HRlXzr8hLdEJkykRNrIDkGS4NMO1MxOpjW1W/9P8wM6H+qQKBgDbZtuPrlu2zHJu9UL+7IyP97Lbna1kMw4O/SNFJz9sJmQ6AMRluonR2J8Lkrpa6O5B24UAcZJ7l5DS9DFKFaWHcPzgo7eqvN/2z1gFSUJO78Ei0u0l1py+9mjHalAIIWTpChSFz7OdClTRPcCZXcxEzRoOk4TCrWlUt3mFqkjlxAoGAT7lbU6pG/Px1diXgJgbQV7xykgGs0lEL6IRb+5vH1uckilX1Tv7QznV15THCKrYJHufbEj6GyWk3A+9FfMUGCQYF+nRYw8TQ10+sTStNcil1sODuz4Icb/6TF6b05VPOz6yJT/wh0lhUohgWE7NINsohgy9Hc24dOglZEennMukCgYEAtO7SR+aoCKrT6brPb2wk1taPcBsnR5Ecavvsvp2BWYJ2Xmju0Q8jeOSdrv7g4SGAiwTHO03Lij3WnKwrVgXGCCkSOCTKbWRUAm6FncIE7j7fD30p7xXltICPea7DsOjJu5c9XlQ6TiPj83J+Mv7NoamHwSKE7tM5bos69iWBtdo=
	-----END PUBLIC KEY-----`
	bytePrivateKey := []byte(privateKey)
	priBlock, _ := pem.Decode(bytePrivateKey)
	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		fmt.Println("1111111111111")
		return ""
	}
	decryptText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, plainText)
	if err != nil {
		fmt.Println("222222222222")
		return ""
	}
	return string(decryptText)
}

// 加密
func RsaEncode(plain string) string {
	msg := []byte(plain)
	//公钥
	publicKey :=
		`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsjXCa4kb3qjHvPjOWSPf
9dkETfqjPSwBvzJNCxf8QXU9eCC57wILWunBvceOAPx1tLUREes2Ny/CScQUjW9I
MXvSuxOWnt4pn2A0TSHWaxZ8LdLnY/1EyBFKG4buXIqD2IgPu7RRel9IiAVMv7kp
aQgWXU5VfN4jqf8lH8Bu1Fm4fDdW6EA0aFOVWJmvbSFhN8ndahQ0AJY+2vgvEhKG
DKYfSKE8+qlLugftYeUucggO7Tc5na5cZ5uAqMyKq9zWahITPPRBbGwucyIGHIDs
5zHPPpp6k2wv5wAy7gvJunwdl82067T76UTHSJ/ZUq+Q5KrFvRN//Hp5Z0GBRrD/
dQIDAQAB
-----END PUBLIC KEY-----`
	//解码公钥
	pubBlock, _ := pem.Decode([]byte(publicKey))
	//读取公钥
	pubKeyValue, _ := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	pub := pubKeyValue.(*rsa.PublicKey)
	encryptText, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	//转为base64
	return base64.StdEncoding.EncodeToString(encryptText)
}
