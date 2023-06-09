package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Go 的 math/rand 包提供了伪随机数生成器。*/
//注意 是伪随机 伪随机！

func main() {
	/*例如，rand.Intn 返回一个随机的整数 n，且 0 <= n < 100。*/
	fmt.Print(rand.Intn(100), ", ")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	//81, 87
	//一直是81，87  所以感觉有点假
	/*PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	*/

	/*rand.Float64 返回一个64位浮点数 f，且 0.0 <= f < 1.0*/
	fmt.Println(rand.Float64())
	//0.6645600532184904  也是假的 一直不变
	/*PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	0.6645600532184904
	PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	0.6645600532184904
	PS D:\GolandProjects\gobyexample> go run .\random-numbers.go
	81, 87
	0.6645600532184904
	*/

	/*这个技巧可以用来生成其他范围的随机浮点数， 例如，5.0 <= f < 10.0*/
	fmt.Print((rand.Float64()*5)+5, ", ")
	fmt.Print((rand.Float64())*5 + 5)
	//7.1885709359349015, 7.123187485356329
	//也是假的
	fmt.Println()

	/*默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
	要产生不同的数字序列，需要给定一个不同的种子。
	注意，对于想要加密的随机数，使用此方法并不安全，
	应该使用 crypto/rand。*/
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	/*调用上面返回的 rand.Rand，就像调用 rand 包中函数一样。*/
	fmt.Print(r1.Intn(100), " , ")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	//这样随机不重样了

	/*如果使用相同种子生成的随机数生成器，会生成相同的随机数序列。*/
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), " , ")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	//结果一直是5，87

	//3系列也用42 做种子
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), " , ")
	fmt.Print(r3.Intn(100))
	fmt.Println()
	//预测跟上边结果一样 因为种子相同
	/*5 , 87
	5 , 87
	*/

}
