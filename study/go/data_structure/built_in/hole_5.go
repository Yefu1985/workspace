package built_in

import "fmt"

/*
这个坑是关于Go里面的interface。
关于面向对象编程中的这些概念，感觉下面这些概念都是需要掌握的，在这里只能大概提一下了，依照JAVA的原则来理解的:
1. 类: class
2. 实例化: initialization
3. 构造方法: constructor
4. 重写: override
5. 重载： overload
6. 抽象类 abstract class
7. 接口 interface
8. 多态 polymorphism
9. 类的继承 extend
10.接口的implement

这些概念网上一搜一大堆，有时候解释的不是很好，如果是初学者的话，有时候会在云里雾里的。虽然Go并不是那么复杂的OOD，但是这些概念还是挺重要的，这里尝试用最浅显的语言来简单翻译一下吧:
1. 首先是类的概念，在OOD中，一切都是object和class。 Class是object的抽象化，object是class的具体化，就这么简单。我们以dog为例子，Dog就是一个类，
那么这个类就有一些自己的变量，有一些自己的方法，比如说每个Dog都有年龄，名字，品种， 性别这些变量，也都会Eat, 都会Sleep，等等这些方法。
那么这个Dog的class就是一个狗狗的蓝图，它并没有具体值哪一条狗狗，只是抽象出了狗狗这个概念而已。OK，那么，具体到每一条狗狗，那就是一个具体的实例了object了，
根据这个蓝图生成的实例就叫做实例化initialization.
2. 在这个狗狗被实例化的过程中，会涉及到一个构造方法的概念，也就是constructor。比如说，你的constructor可以以Age, Name, Breed, Gender作为参数，
只要这几个参数设定了，就算是实例化了一个具体的狗狗了。
3. 那么，构造方法可以是多种多样的, 可以选择用不同的参数来实现实例化。再说大一点，就是说在一个class里面，如果有一些方法的方法名相同，但是传给这个方法的参数不同，
这就叫做方法的重载,overload;
4. OK, 现在有了一个叫做Dog的类别; 我们想在这个基础上再细分一下，比如说有一种狗狗叫做拉普拉多，我们想给它做一个单独的class，那么这个class的名字就叫做labrado；
我们很容易发现，即使是单独出了这个labrado的class，它还是有很多的特性跟Dog是一致的，比如说，它还是会有名字，年龄，性别，还是会Eat, Sleep等等。
我们不想在定义这个Labrado的class的时候又把这些东西重复写一遍，那么我们可以做的是让这个Labrado作为Dog的一个子类(subclass). 这个在逻辑上也很合理。
也就是说subClass会是上一级class的一个小的分支。这在OOD里面就叫做类的继承，在JAVA中的写法就是class Labrado extends Dog
5. 毫无意外，Labrado这个class会继承Dog里面的所有方法，包括Sleep，但是，如果我们觉得Labrado的Sleep和Dog里面的Sleep方法不太一致，我们想重新定义Labrado
的Sleep的方法，那么我们可以在Labrado class里面重写一下新的Sleep方法，这个过程就叫重写override, 那么Labrado的实例就会执行这个新写的Sleep方法
6. 有一种很特殊的类，叫做抽象类，抽象类里面只有方法的概述，并没有方法的实体，比如说，创建一个抽象类叫做abstract class animal； 那么在这个类中，
除了一些变量之外，我们还可以定义几个抽象方法，比如说Eat， Drink， Sleep， 等等。我们知道所有的animal都会有这几个方法，但是，具体每个animal会怎么实行呢？
我们不知道，也不需要在animal这个level就搞清楚。而是可以等到某个实现了这个abstract class的实例object上去真正实现这个方法，比如说Dog这个类，可以从
animal这个抽象类继承，然后在Dog里面实现Eat, Drink, Sleep的具体方式。抽象类并不能被实例化，只可以被其它的类来继承。
7.接口也是一个抽象的概念，经常会和抽象类搞混。接口不能有constructor，也不能被继承，它里面都是一些抽象的方法, 它和抽象类最大的区别在于: 抽象类是抽象"是什么"，
而接口是抽象“干什么”。 比如说，我可以定义一个接口interface live(), 里面只有四个方法,吃喝拉撒，对吧。那么对于不同的class的实例来说，比如说，对于猫猫，狗狗，
甚至human，都可以实现这四个方法，只是实现的手段不一样。而这些class，有时候不不方便用subclass的方式来继承和重写方法，所以就有了这个接口的概念了。
8. 在之前的例子中，如果实例化了一个Labrado叫做Nick，因为这个Labrado是从Dog里面extend过来的，所以，这个Nick既是Labrado，同时也是Dog。就说明它同时
有了两个形态，这就是多态Polymorphism. 但是Nick在Labrado里面，还想有别的Dog里没有的func，那么怎么办呢？这里就可以借助于interface了，让Labrado去
implement Interface就好了，然后就可以再具体写自己想要的func了。

当然，真实中的OOD比我上面说的这些复杂多了，还有好多没有提到的点，如果有兴趣的话，强烈推荐看一看JAVA的OOD，理解会比较深刻。OOD的写法不仅仅要对编程很熟悉，
而且也需要对业务很熟悉，才能尽可能地优化整个OOD的机制。
好在Go里面对于OOD没有这么复杂的机制，大体上还是简单了很多，所以官方也是说Go是支持面向对象的，但是不是纯粹的面向对象。

*/

type DogClass struct { // 这里我用了DogClass，而不是用Dog，是因为在这个built_in package里面，在hole4里面，我用了Dog struct，在同一个package内部，不可以有同名的struct
	Name string
	Gender string
	Age int
}

func NewDog(name string, gender string, age int) *DogClass {
	return &(DogClass{name, gender, age})
}

type DogClassSuper struct {
	Name string
	Gender string
	Age int
	*Labrado
}

type Labrado struct {
	Color string
	Height float32
	Weight float32
}

func NewDogSuper(name string, gender string, age int, labrado *Labrado) *DogClassSuper {
	return &(DogClassSuper{name, gender, age, labrado})
}

func NewLabrado(color string, height float32, weight float32) *Labrado {
	return &(Labrado{color, height, weight})
}

func (dogClassSuper *DogClassSuper) Bark() {
	fmt.Println("Wang1, Wang1, Wang1")
}

func (dogClassSuper *Labrado) Bark() {
	fmt.Println("Wang2, Wang2, Wang2")
}

type Live interface {
	Eat(food string) int
	Drink()
}

func (dogClassSuper *DogClassSuper) Eat(food string) int {
	fmt.Println("I am having food: "+food)
	return 10
}

func (dogClassSuper *DogClassSuper) Drink() {
	fmt.Println("I am enjoying drinking.")
}

type Live2 interface {
	Eat(food string) int
	Drink()
}

func Hole5Example1() {
	fmt.Println("Welcome to hole #5.")
	// 首先，还是要创建一个struct，和class差不多对等吧，还是用Dog来创建类
	dog1 := DogClass{"nick", "male", 2}
	_ = dog1
	// 上面是实例化一个Dog的方法，Go并没有JAVA中那样的constructor，所以，可以自己写一个；
	dog2 := *(NewDog("nick", "male", 2)) //我在NewDog这个构造函数中，返回的是DogClass的指针，所以，这里要对指针取值才可以得到dog2的实例；
	fmt.Println(dog2)

	// 注意，Go是不支持重载overload的，也就是说不能有同名的func，即使它们的传入参数不一样也不可以;

	// 类的继承是通过struct的嵌套完成的，比如说在DogClass里面再嵌套一个Labrado的struct
	// 注意，我用了DogClassSuper来区分DogClass，所以，我也写了两个constructor函数，分别是对应Labrado和DogClassSuper的。
	// 在Labrado的构造函数NewLabrado中，我返回的还是指针，所以，在DogClassSuper中，我对Labrado也要用指针
	labrado1 := NewLabrado("white", 10.0, 20.0)
	dogSuper1 := NewDogSuper("Jack", "female", 3, labrado1)
	fmt.Println(*dogSuper1) // {Jack female 3 0xc0000044e0} 这才是打印结果，因为dogSuper1里面的Labrado是一个指针，所以就只会打印出它所指向的地址而已;

	// 现在，我想给类加一个方法；这个比较简单，在hole_4里面已经讲过了，就是用func加上receiver的方法就可以实现了;
	// 比如说，我们给DogClassSuper加一个方法叫做Bark()吧；用值或者是用reference作为receiver都可以，我用了reference
	// 毫无疑问，这样就是类的方法。那么，我们看看这个方法可不可以在子类Labrado中重写,答案是可以的
	dogSuper1.Bark() // 这个会输出Wang1, Wang1, Wang1
	labrado1.Bark()  // 这个会输出Wang2, Wang2, Wang2

	// 下面是我们的重头戏，Interface的用法。在JAVA中，implement某一个interface是要显性说明的，但是在Go里面不需要。
	// Interface在Go里面是一种type, 它里面定义了一些func，也就是这个interface的行为；如果某一个类型实现了这个interface里面所有的方法，
	//那么就说明这个类型实现了这个interface。
	// 我们还是以interface Live()为例吧，假设里面有Eat()和Drink()这两个方法，其中Eat需要有输入，比如说bone, 返回一个int，就是需要的时间seconds；
	// Drink不需要输入参数，也没有输出参数

	// 和实现struct里面的func很类似，这个interface里面的方法实现，也是用receiver来实现的。
	dogSuper1.Eat("bone")
	dogSuper1.Drink()
	// 所以在Go中，实现struct里面的方法(也就相当于类方法) 和实现接口interface里面的方法其实是一样的，都可以认为是用receiver去实现了一个func，仅此而已

}