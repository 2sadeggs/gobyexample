package main

import (
	"fmt"
	s "strings" //为strings包起个别名
)

/*
标准库的 strings 包提供了很多有用的字符串相关的函数。 这儿有一些用来让你对 strings 包有一个初步了解的例子。*/
//我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

func main() {
	/*这是一些 strings 中有用的函数例子。
	由于它们都是包的函数，而不是字符串对象自身的方法，
	这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。
	你可以在 strings 包文档中找到更多的函数。*/

	p("Contains:  ", s.Contains("test", "es"))  //是否包含
	p("Count:     ", s.Count("test", "t"))      //计数
	p("HasPrefix: ", s.HasPrefix("test", "te")) //判断是否是前缀
	p("HasSuffix: ", s.HasSuffix("test", "st")) //判断是否是后缀
	p("Index:     ", s.Index("ttttttest", "e"))
	//Index() 函数返回 int 类型的值，如果包含，则返回第一次出现该字符串的索引；反之，则返回 -1。
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //组合两个字符串
	p("Repeat:    ", s.Repeat("a", 5))                //单纯重复字符串5次 无复杂函数
	p("Repeat:    ", s.Repeat("qwer", 3))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	/*在这里，s是原始或给定的字符串，old是您要替换的字符串。 new是替换旧版本的，n是旧替换的次数。
	注意：如果old为空，则它在字符串的开头和每个UTF-8序列之后匹配，最多可对k-rune字符串进行k + 1个替换。
	如果n <0，则替换次数没有限制。*/
	p("Split:     ", s.Split("a-b-c-d-e", "-")) //切割
	p("ToLower:   ", s.ToLower("TEST"))         //转小写
	p("ToUpper:   ", s.ToUpper("test"))         //转大写
	p()
	/*虽然不是 strings 的函数，但仍然值得一提的是， 获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制。*/
	p("Len :", len("hello"))
	p("Char :", "hello"[1])
	/*注意，上面的 len 以及索引工作在字节级别上。
	Go 使用 UTF-8 编码字符串，因此通常按原样使用。
	如果您可能使用多字节的字符，则需要使用可识别编码的操作。
	详情请参考 strings, bytes, runes and characters in Go。*/

}
