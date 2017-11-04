package main

import (
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

//SHA1散列经常用于计算二进制或文本块的短标识。

/*
SHA-1与MD5 的最大区别在于其摘要比MD5 摘要长 32 比特。对于强行攻击，
产生任何一个报文使之摘要等于给定报文摘要的难度：MD5 是2128 数量级的操作，
SHA-1 是2160 数量级的操作。产生具有相同摘要的两个报文的难度：MD5是 264 是数量级的操作，
SHA-1 是280 数量级的操作。因而,SHA-1 对强行攻击的强度更大。但由于SHA-1 的循环步骤比MD5多（80:64）
且要处理的缓存大（160 比特:128 比特），SHA-1 的运行速度比MD5 慢。
*/
func main() {
	//SHA-1
	s := "hello world"
	sha := sha1.New()
	sha.Write([]byte(s))

	bs := sha.Sum(nil)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs) //%x 转换为a hex string
	fmt.Println()

	//MD5
	md := md5.New()
	md.Write([]byte(s))
	bm := md.Sum(nil)
	fmt.Printf(hex.EncodeToString(bm))
}
