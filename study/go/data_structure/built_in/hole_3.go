package built_in

import "fmt"
/*
指针其实是C语言体系里面才有的概念，JAVA, Python中都没有，很多新手都不太喜欢这个概念，感觉不是很好掌握。
其实也还好，我一般都这样理解，真正的数据data住在一个大屋子里，你可以用一个变量直接指向它，这没问题，但是，如果我希望有好几个变量都和它挂钩呢，那咋办呢？
所以就会有一个指针的概念，这个指针其实就是这个大屋子的门牌号，或者叫做内存中的地址，然后把这个门牌号存进另外一个大屋子，所以和这个存了门牌号的大屋子挂钩的就叫指针。
不废话了，干就完了。
 */


func Hole3Example1() {
	fmt.Println("Welcome to hole #3.")

	// 先建立一个string变量, 它的值就是"a"，先看下它的门牌号是多少
	eString := "a"
	fmt.Println(&eString) // 这里打印出来的就是eString这个变量的内存地址，也就是门牌号；

	var ePointer *string // 定义了一个指针变量，名字叫做ePointer，说明它的地址指向的数据类型会是string
	ePointer = &eString // 这样可以把这个ePointer指向eString
	fmt.Println(ePointer)

	// 在编程中有个很重要的概念是pass by value, or pass by reference。随便在网上找个文章看下就好了。
	// Go里面是pass by value的，也就是说，如果把一个array作为参数传给func的话，那么在func里面会copy by value，也就是会生成一个副本，所有的操作都会在array的副本上执行的。
	// 但是，有很多情况下，我们并不希望这样。主要有两个原因: 1. 数据本身很大，再搞一个copy出来，内存吃不消啊，亲； 2. 我们希望这个func中对array的修改会生效。
	// 在这种情况下，我们有两个方法实现: 1. 尽量用reference型的数据结构，比如说不要用array，而用slice; 2. 用指针。
	// 用指针的话，只会把门牌号，也就是内存地址传进去，所以，对这个门牌号里面的元素的修改，一样会生效的。
	// ok，好了，还是继续用例子说明吧。

	arr := [5]int{0, 1, 2, 3, 4} // 先搞一个arr再说
	// case #1: 把arr传进func, 然后对arr进行修改，看看arr是否会有变化
	case1(arr)
	fmt.Println(arr) // 这个func根本不会对arr进行修改，因为它copy一个备份;[0 1 2 3 4]

	// case #2: 用slice，而不要用array了
	slice := arr[:]
	case2(slice)
	fmt.Println(arr) // 如我们预料，把slice传进去，其实就是把地址传进去了，这样对slice的修改也会到arr上的；[100 1 2 3 4]

	// case #3: 用指针，而不用array
	case3(&arr)
	fmt.Println(arr) // 嗯，用指针也是这样的效果: [100 200 2 3 4]

}

func case1(arr [5]int) {
	arr[0] = 100
}
func case2(slice []int) {
	slice[0] = 100
}
func case3(arrPointer *[5]int) {
	arrPointer[1] = 200
}