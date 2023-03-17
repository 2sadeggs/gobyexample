package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	//可以单独range key 值
	//range can also iterate over just the keys of a map.

	for i, c := range "go" {
		fmt.Println(i, c)
	}
	//输出：0 103 1 111  第一个为索引值 第二个为Unicode值
	//range on strings iterates over Unicode code points.
	//The first value is the starting byte index of the rune and the second the rune itself.
	//See Strings and Runes for more details.
}
