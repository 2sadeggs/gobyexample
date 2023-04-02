package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

/*
SHA256 散列（hash） 经常用于生成二进制文件或者文本块的短标识。
例如，TLS/SSL 证书使用 SHA256 来计算一个证书的签名。
这是 Go 中如何进行 SHA256 散列计算的例子。*/

//Go 在多个 crypto/* 包中实现了一系列散列函数

func main() {
	//定义一个字符串
	s := "sha256 this string"

	//这里我们从一个新的散列开始。
	h := sha256.New()

	//写入要处理的字节。如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组。
	h.Write([]byte(s))

	//Sum 得到最终的散列值的字符切片。Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做。
	bs := h.Sum(nil)

	//SHA256 值经常以 16 进制输出，例如在 git commit 中。 我们这里也使用 %x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	/*sha256 this string
	1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
	*/
	fmt.Println(bs)
	//[26 241 223 168 87 191 29 136 20 254 26 248 152 60 24 8 0 25 146 46 85 127 21 168 160 211 219 115 157 119 170 203]

	/*你可以使用和上面相似的方式来计算其他形式的散列值。
	例如，计算 SHA512 散列， 引入 crypto/sha512 并使用 sha512.New() 方法。*/

	ss := "this is a SHA512 string"
	hh := sha512.New()
	hh.Write([]byte(ss))
	bbs := hh.Sum(nil)
	fmt.Println(ss)
	fmt.Printf("%x\n", bbs)
	/*this is a SHA512 string
	2560aa09ebb6e124f412c8ddc463b365148b4f40c592dd7e55acd6ad509754520b4052e0558790223b491b86896e6a4f2cab9e414d711c5b41782b6de1d63a8d
	*/
}

/*注意，如果你需要密码学上的安全散列，你需要仔细的研究一下 加密散列函数。*/
