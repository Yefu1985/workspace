package built_in

import (
"fmt"
"reflect"
)

/*
There is no char type in Go. Actually, rune is char type. So, a char will occupy 4 bytes (32 bits).
汉字是用utf-8编译的，一个汉字占用3个byte;
一个英文字符占用两个byte;
 */
func CharExample() {
	r := 's' // 这里用的是单引号，所以其实定义的是一个rune类型，本质上也就是int32
	fmt.Printf("%q has 4 bytes, and its binary value is %b\n", r, r)
	rType := reflect.TypeOf(r).Kind() // 这个语句可以知道r的实际类型，其实就是int32;
	fmt.Printf("the real type of r is %s\n", rType)

	char := '我'
	fmt.Printf("%q has 4 bytes, and its binary value is %b\n", char, char)
	charType := reflect.TypeOf(char).Kind() // 这个语句可以知道r的实际类型，其实就是int32;
	fmt.Printf("the real type of char is %s\n", charType)
}

func StringExample() {
	r := "s" // 这里用的是双引号,就是一个字符串了，一个英文字母就是一个byte;
	fmt.Printf("%q has %b byte.\n", r, len(r))
	rType := reflect.TypeOf(r).Kind() // 这个语句可以知道r的实际类型，当然是string;
	fmt.Printf("the real type of r is %s\n", rType)

	char := "我" // 一个汉字需要三个byte
	fmt.Printf("%q has %v bytes.\n", char, len(char))
	charType := reflect.TypeOf(char).Kind() // 这个语句可以知道r的实际类型，当然是string;
	fmt.Printf("the real type of char is %s\n", charType)
}

func CountLength() {


}