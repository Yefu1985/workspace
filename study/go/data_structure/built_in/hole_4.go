package built_in

import "fmt"

/*
这个坑主要是关于struct的。
用过C的话，对struct应该有一定认识，如果用过JAVA或者是Python的话，struct就和class很类似。
但素，Go中的OOD的概念很淡薄，所以没有那些继承，重载，构造，等等方法，但是可以有各种方式实现类似功能。所以，也就没有那些extend，implement之类的用法。
这个坑里面主要展示在Go里面怎么玩struct.
 */

// 我们可以假设建立一个小一点的struct，比如说dog, 它总有一些变量;

type Dog struct {
	Name string
	Age int
	Breed string
	Color string
	DogSize Size
}

// 设置一个size struct, 再把dogSize放进dog里面;
type Size struct {
	Height float32
	Length float32
	Weight float32
}

func Hole4Example1() {
	fmt.Println("Welcome to hole #4.")
	// 先实例化一个Dog吧,有好几个方法可以实例化的；
	dog := Dog{Name:"puppy", Age:4, Breed:"labrador", Color:"black"} // 方法1;
	//dog := new(Dog) //这样的初始化会让所有的成员变量都是默认零值；
	fmt.Println(dog)

	// 很自然的，我们想给dog这个type加一些方法，比如说，rename();
	dog.Rename1()
	fmt.Println(dog.Name) // 如果是用Dog作为传递参数的话，pass by value; 会创建dog的副本，并不会对本身的dog实例产生影响;
	dog.Rename2()
	fmt.Println(dog.Name) // 这个是用指针作为传参的，那么对这个dog的修改会影响到指针指向的dog;

	// 下面再来一个struct嵌套的话题，比如说dog还会有一个属性，就是size，这个size本身又可以作为struct，里面包含height, length, weight
	dog2 := Dog{"nick", 3, "Chihuahua", "white", Size{1.0,2.0,3.0}}
	fmt.Println(dog2.DogSize.Height) // 这里很有意思，如果在Dog里面，你用了DogSize作为Size的变量名的话，那么这里的DogSize就不能省略；
	// 如果在Dog里面的定义没有写DogSize这个变量名的话，这里的DogSize就可以省略了；这个其实很好理解，这就是一个结构体，为啥还要给它取个名呢；

	// 还可以再试一下，如果当前的struct和内嵌的struct的变量名有重合的呢，该咋办；
	// 可以通过cat这个struct来验证，在cat这个struct里面，也有weight，在Size这个struct里面也有weight，试下看
	cat := new(Cat) // 这是第二种实例化方法;
	cat.Name = "Kitty"
	cat.Age = 2
	cat.Breed ="persian"
	cat.Color = "white"
	cat.Weight = 1.0
	cat.Size = Size{2.0,3.0,4.0}
	fmt.Println(cat.Weight)
	fmt.Println(cat.Size.Weight)
	// 嗯，结论出来了，并不会报错，会先打印最高级别的value，也就是1.0, 而不是Size里面的4.0
}

// Go里面给struct加func比较特殊，是用一种叫做receiver的概念来实现的; 比如下面的Bark前面的(dog Dog)就叫receiver;
// 这个receiver可以用Dog这样的type，也可以用*Dog这样的指针type； 我们写两个Rename函数，看一下有什么不同;
// Well， 不论是哪个，效果都是一致的；
func (dog Dog) Rename1() {
	dog.Name = "Rename1"
}

func (dog *Dog) Rename2() {
	dog.Name = "Rename2"
}

type Cat struct {
	Name string
	Age int
	Breed string
	Color string
	Weight float32
	Size
}
