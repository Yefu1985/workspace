package built_in

import "fmt"

/*
这个坑主要是整理一下array和slice的关系。其实挺简单的，array就是一个固定大小的数组，slice就是类似于python中的切片，底层还是array。
array其实用的不多，slice倒是挺多的。
在slice中，比较重要的概念是len()和cap().
还是通过例子来慢慢理解吧。
 */

func Hole2Example1() {
	fmt.Println("Welcome to hole #2.")
	//首先定义一个数组array，有这么几种方式
	arr1 := [5]int{0,1,2,3,4} //这个就是最常用的初始化，里面的那个数字5就是数组长度；
	arr2 := [...]int{0,1,2,3,4} //这个也可以，就是不用告诉它长度，Go自己会去算；
	arr3 := new([5]int) //这个应该很少用吧，用new来创建一个数组arr3; 这里其实只是创建了一个指针而已；
	_, _ = arr1, arr2
	fmt.Println(arr3) // 这里会返回 &[0 0 0 0 0]; 看吧，只是指针而已;

	// 再定义一个slice, 有这么几种方式
	slice1 := []int{1, 2, 3, 4} // 最基本的方式，跟array很像，只是不需要再specify长度了;
	_ = slice1
	slice2 := arr1[1:4] //这就是把arr1 {1, 2, 3, 4, 5}里面的index从1到3的都拿出来了; {2, 3, 4}
	slice3 := make([]int, 3, 10) // 这也挺常用的, 用make方法来做，第一个3是长度，第二个10是capacity;第二个capacity也可以不定义，没关系；
	// new 和 make一个很大的区别在于，new会返回指针，make会分配空间，并且初始化所有的内存为零值;

	//先来打印一下slice2，看看它的信息
	fmt.Printf("slice2的值为: %v\n", slice2)
	fmt.Printf("slice2 的length为%d, capacity为%d.\n", len(slice2), cap(slice2))
	// slice2的长度为3，这个很好理解，slice2的capacity为4，这个是因为slice2是从arr1中切片过来的，从index = 1开始,直到最后，一共有{1,2,3,4}
	// 四个value，所以它的capacity是4;

	//再来看一下slice3是什么鬼
	fmt.Printf("slice3的值为: %v\n", slice3)
	fmt.Printf("slice3 的length为%d, capacity为%d.\n", len(slice3), cap(slice3))
	// slice3的打印值为0, 0, 0，很简单，因为make初始化了所有的element为默认零值，对于int类型来说，那就是0啊，亲，所以是3个0


	// 下面说一下slice的几个小坑
	// 1. slice只是切片而已，对slice的修改一样会修改它对应的底层array；举个例子更容易理解；
	slice2[2] = 100 //这个对slice2里面的index = 2的位置做出修改,slice2本身是{1,2,3}, 也就是说会把3变成100;
	fmt.Println(slice2) //这里会打印成1, 2, 100，很好理解；
	fmt.Println(arr1) //这里也会打印成0, 1, 2, 100, 4，因为对slice2的修改也会反映到arr1上;

	// 2. append的坑
	// 这个append和python中的有一点点不一样，在slice里面的append，如果append的时候还没超过底层array的长度的话，那就会对array进行修改；
	//如果超过了的话，那slice就会新建一个array，把slice里面所有的元素都复制过去。 嗯，还是举例说明.
	slice2 = append(slice2, 200)
	fmt.Println(slice2) // 如我们所预测，slice2会是 {1, 2, 100, 200}
	fmt.Println(arr1) // 因为slice2即使加上一个200，也还没超过arr1的范围，所以，arr1就会变成{0, 1, 2, 100, 200}

	// 再继续添加元素300
	slice2 = append(slice2, 300)
	fmt.Println(slice2) // 这个肯定没问题，slice2会变成{1, 2, 100, 200, 300}
	fmt.Println(arr1) // Look, 因为300已经加不进arr1了，所以，arr1只会是{0, 1, 2, 100, 200}；
	fmt.Printf("the len of slice2 is %d, the capacity of slice2 is %d. \n", len(slice2), cap(slice2))
	// 嗯，看一下slice2的长度，已经变成了5了，很好理解，因为有1,2,100,200,300; 但是capacity会变成8，是因为当append的时候，
	// 如果capacity不够了，就会自动扩展capacity;比如说double一下，或者是增加50%之类的；

	// 我刚才是逐步将arr1里面的元素通过append on slice2的方法替换掉了，如果我一口气就append很多元素(多余arr1的长度)在slice2上呢，会怎样呢？
	slice4 := arr1[1:4]
	slice4 = append(slice4, 101, 201, 301, 401)
	fmt.Println(slice4)
	fmt.Println(arr1) //当然，如果一口气就append了那么多的话，系统是会detect到的，会直接开辟一个新的arr，而不会再去修改原先底层的arr了
}
