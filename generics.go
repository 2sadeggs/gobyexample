package main

import "fmt"

/*
作为泛型函数的示例，MapKeys 接受任意类型的Map并返回其Key的切片。
这个函数有2个类型参数 - K 和 V； K 是 comparable 类型，
也就是说我们可以通过 == 和 != 操作符对这个类型的值进行比较。
这是Go中Map的Key所必须的。
V 是 any 类型，意味着它不受任何限制 (any 是 interface{} 的别名类型).
*/
/*
函数名 MapKeys
入参(m map[K]V) 也就是说入参是个map类型  但是这个map类型的K和V特殊 不确定是那种类型
于是有了前边的中括号说明
[K comparable, V any] K必须是个可比较的类型（go要求） any是任意类型 (any 是 interface{} 的别名类型).
返回值[]K 也就是K类型的切片
*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	//声明并初始化一个切片 切片类型为K 切片
	r := make([]K, 0, len(m))
	/*
		第二参数指定的是切片的长度，第三个参数是用来指定预留的空间长度，
		例如a := make([]int, 2, 4), 这里值得注意的是返回的切片a的总长度是4，
		预留的意思并不是另外多出来4的长度，其实是包含了前面2个已经切片的个数的。
		所以举个例子当你这样用的时候 a := make([]int, 4, 2)，就会报语法错误。
	*/
	//遍历map的key 追加到切片r
	/*for k range m{
		想当然的错误写法 所以一定不要眼高手低  初期更要多动手
	}*/
	for k := range m {
		r = append(r, k)
	}
	return r
}

//List 是一个具有任意值的单链表
type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

/*
我们可以像在常规类型上一样定义泛型类型的方法 但我们必须保留类型参数。 这个类型是 List[T]，而不是 List
*/
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	/*
		当调用泛型函数的时候, 我们经常可以使用类型推断。 注意，当调用 MapKeys 的时候，
		我们不需要为 K 和 V 指定类型 - 编译器会进行自动推断
	*/
	fmt.Println("keys m: ", MapKeys(m))
	//注意 每次打印出来的may keys 顺序不一定相同

	//… 虽然我们也可以明确指定这些类型
	_ = MapKeys[int, string](m)
	lst := List[int]{}
	fmt.Println("chushi lst: ", lst)
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list: ", lst.GetAll())
}
