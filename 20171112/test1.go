package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

/*
Go语言的crypto里面支持对称加密的高级加解密包有：
crypto/aes包：AES(Advanced Encryption Standard)，又称Rijndael加密法，
是美国联邦政府采用的一种区块加密标准。
crypto/des包：DES(Data Encryption Standard)，是一种对称加密标准，是目前使用最广泛的密钥系统，
特别是在保护金融数据的安全中。曾是美国联邦政府的加密标准，但现已被AES所替代。
*/
func main() {
	str := "hello world"
	//base64加解密
	b := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(b)
	c, _ := base64.StdEncoding.DecodeString(b)
	fmt.Println(string(c))

	//aes
	//填充
	var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	var key = "astaxie12798akljzmknm.ahkjkljl;k"
	fmt.Println(len(key))

	//参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	//,返回了一个cipher.Block接口，这个接口实现了三个功能
	d, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	//加密字符串

	cfb := cipher.NewCFBEncrypter(d, commonIV)
	ciphertext := make([]byte, len(str))
	cfb.XORKeyStream(ciphertext, []byte(str))
	fmt.Printf("%s=>%x\n", str, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(d, commonIV)
	plaintextCopy := make([]byte, len(str))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

}

/*
type Block interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Decrypt(dst, src []byte)
}
*/
