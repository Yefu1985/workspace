package built_in

import "fmt"

/*
Go语言中比较牛逼的是可以将func作为一个传递参数传给别的func。
还有一个闭包的概念可以了解一下。
多说无益，上例子吧。
 */

func Hole6Example1() {
	fmt.Println("Welcome to hole #6.")
	// 先说闭包closure吧，就是说可以把一个函数整体直接复制给某个变量，还是举例吧
	add := func(x int, y int) int {
		return x + y
	} // 整个func其实就是把x和y加了一下; 但是，这里的add这个变量可不是int类型了，而是func(int, int) int 类型了
	fmt.Printf("add 的类型是%T \n", add)

	// 这个时候，你就可以直接调用这个函数add(3, 4)了
	fmt.Println(add(3, 4))

	// 即使不赋值给某个变量，一样可以调用, 但是感觉这个方法的可读性不是很好，我不太喜欢；
	addValue := func(x, y int) int { return x + y } (30, 40)
	fmt.Println(addValue)

	// OK, 还有一种牛逼的是把这个作为参数传给其它的func
	fmt.Println(external(add))
}

func external(add func(int, int) int) int {
	fmt.Println("Now, I am in external function.")
	return add(300, 400)
}
