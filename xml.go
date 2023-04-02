package main

import (
	"encoding/xml"
	"fmt"
)

/*Go 通过 encoding.xml 包为 XML 和 类-XML 格式提供了内建支持。*/

/*Plant 结构将被映射到 XML 。 与 JSON 示例类似，字段标签包含用于编码器和解码器的指令。
这里我们使用了 XML 包的一些特性：
XMLName 字段名称规定了该结构的 XML 元素名称；
id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素。*/

//定义一个结构体与xml映射
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	//初始化一个plant 植物 叫coffee 咖啡
	coffee := &Plant{Id: 27, Name: "Coffee"}
	//单独初始化Plant的Origin 产地  是个string切片
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	//我打印一下看看
	fmt.Printf("xianzaishizheyang: %#v\n", coffee)
	//xianzaishizheyang: &main.Plant{XMLName:xml.Name{Space:"", Local:""}, Id:27, Name:"Coffee", Origin:[]string{"Ethiopia", "Brazil"}}

	/*传入我们声明了 XML 的 Plant 类型。 使用 MarshalIndent 生成可读性更好的输出结果。*/
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	/*	<plant id="27">
		<name>Coffee</name>
		<origin>Ethiopia</origin>
		<origin>Brazil</origin>
		</plant>
	*/
	//明确的为输出结果添加一个通用的 XML 头部信息。
	fmt.Println(xml.Header + string(out))
	//	<?xml version="1.0" encoding="UTF-8"?>

	//使用 Unmarshal 将 XML 格式字节流解析到 Plant 结构。 如果 XML 格式错误或无法映射到 Plant 结构，将返回一个描述性错误。
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	//Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Printf("out jiexiedao Plant: %#v\n", p)
	//out jiexiedao Plant: main.Plant{XMLName:xml.Name{Space:"", Local:"plant"}, Id:27, Name:"Coffee", Origin:[]string{"Ethiopia", "Brazil"}}

	//在初始化一个plant 名字叫西红柿 主要是下面的xml 嵌套用的到
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	/*parent>child>plant 字段标签告诉编码器，将 Plants 中的元素嵌套在 <parent><child> 里面。*/

	type Nesting struct { //看样子叫雀巢
		XMLName xml.Name `xml:"nesting"`
		//Plants  []Plant  `xml:"parent>child>plant"` //第一次写错了 导致
		//.\xml.go:69:19: cannot use []*Plant{…} (value of type []*Plant) as type []Plant in assignment
		Plants []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}
