package built_in

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
这个文件里面包含了所有Go的基本数据类型的一些常规操作.
 */

func CommonUsages() {
	fmt.Println("Welcome to the common usages of basic data types in Go.")

	// ************** Numbers ********************** //
	// Numbers的很多数学操作都是调用了math package，里面的绝大多数的func的输入参数都是float64类型的
	// 极值常数情况
	// 所有的min和max其实都是 (-)1>>(n-1)-1; n是位数；
	fmt.Printf("最大的MaxInt8 is: %d. \n", math.MaxInt8)
	fmt.Printf("最大的MaxInt16 is: %d. \n", math.MaxInt16)
	fmt.Printf("最大的MaxInt32 is: %d. \n", math.MaxInt32)
	fmt.Printf("最大的MaxInt64 is: %d. \n", math.MaxInt64)
	fmt.Printf("最小的MinInt8 is: %d. \n", math.MinInt8)
	fmt.Printf("最小的MinInt16 is: %d. \n", math.MinInt16)
	fmt.Printf("最小的MinInt32 is: %d. \n", math.MinInt32)
	fmt.Printf("最小的MinInt64 is: %d. \n", math.MinInt64)
	fmt.Printf("最大的MaxFloat32 is: %f. \n", math.MaxFloat32)
	fmt.Printf("最大的MaxFloat64 is: %f. \n", math.MaxFloat64)
	fmt.Printf("最小的SmallestNonzeroFloat32 is: %e. \n", math.SmallestNonzeroFloat32)
	fmt.Printf("最小的SmallestNonzeroFloat64 is: %e. \n", math.SmallestNonzeroFloat64)
	fmt.Printf("最大的正无穷是 Inf: %v. \n", math.Inf(1))
	fmt.Printf("最大的正无穷是 Inf: %v. \n", math.Inf(-1))

	// 随机数， 随机数都在math/rand这个package;
	// 这里有个坑，就是rand在设置seed的时候，默认的就是1，不会变的，所以，如果不修改seed的话，就会是输出一样的随机数了；
	// 如果真要随机的话，就要修改掉seed, 可以考虑用纳秒时间来作为seed
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("随机生成一个从0到MaxInt32的随机整数,用rand.Int31(): %v.\n", rand.Int31())
	a, b := 1, 100
	fmt.Printf("随机生成一个从[%v到%v)的随机整数,用rand.Intn(b-a)+a: %v.\n", a, b, rand.Intn(b-a)+a)

	// 进制转换, 用的是strconv package，而不是math里面的
	fmt.Printf("十进制换成二进制要用strconv.FomratInt(int64(b), base(2)): %v -> %v. \n", b, strconv.FormatInt(int64(b), 2))
	Bbinary := 0b1111
	fmt.Printf("二进制换成十进制要用strconv.FomratInt(int64(b), base(10)): %b -> %v. \n", Bbinary, strconv.FormatInt(int64(Bbinary), 10))
	fmt.Printf("二进制换成八进制要用strconv.FomratInt(int64(b), base(10)): %b -> %v. \n", Bbinary, strconv.FormatInt(int64(Bbinary), 8))

	// 控制输出格式: 其实就是fmt里面的几个定义 https://golang.org/pkg/fmt/
	// 常用的有这么几个: 1. %v; 2.%T (type); 3.%t (true or false); 4.%b base 2; 5.%c the char or unicode point; 6.%q single quoted character;

	// 简单的数学运算，pow(x,n), exp, log, 等等
	x, n := 24.0, 3.0
	fmt.Printf("%v 的 %v 次方是: %v.\n", x, n, math.Pow(x, n)) // x, n必须是float64
	fmt.Printf("%v 的平方根是: %v.\n", x, math.Sqrt(x)) // x必须是float64

	// ************** string ********************** //
	s1 := "I love you!"
	// 基本信息，长度，byte
	fmt.Printf("%q 的长度为 %d. \n", s1, len(s1))
	// 判断字母数字, 很不幸，Go本身没有Python中那种isDigit这样的func，只能自己斗智斗勇了，有两个内置的方法是可行的
	// 1.用regex, 2.用strcon.ParseInt()看是否会报错，会报错，就说明不是int
	s2 := "1234"
	fmt.Printf("Check %q 是不是digit用regex: regex.MustCompile(`^[0-9]+$`): %v.\n", s2, regexp.MustCompile(`^[0-9]+$`).MatchString(s2))
	fmt.Printf("Check %q 是不是digit用strconv.ParseInt, 去check是否有err. \n", s2)
	_, err := strconv.ParseInt(s2, 10, 64)
	fmt.Printf("%q is a digit: %v.\n", s2, err == nil)

	// 大小写切换
	fmt.Printf("%q conver to Upper case by strings.ToUpper(s): %q. \n", s1, strings.ToUpper(s1))
	fmt.Printf("%q conver to Lower case by strings.ToLower(s): %q. \n", s1, strings.ToLower(s1))

	// 比较字符串, 用strings.Compare就可以了
	fmt.Printf("比较两个字符串，用strings.Compare(s1, s2)就可以了，s1大于s2，则1，否则为-1，相等为0: compare(\"I love you\", \"1234\") 会输出 %v.\n", strings.Compare(s1, s2))

	// 查找子串
	s3 := "23"
	fmt.Printf("查看string里面是否有substring用 strings.Contains(), %q中是否有%q: %v. \n", s2, s3, strings.Contains(s2, s3))
	fmt.Printf("在字符串中找到第一个子串出现的index要用strings.Index(), %q中第一次出现%q的index在%v. \n", s2, s3, strings.Index(s2, s3))

	// 查找出现次数
	fmt.Printf("查看string里面是有多少个substring要用 strings.Count(), %q中有%v个%q. \n", s2, strings.Count(s2, s3), s3)

	// 删除trim等
	s4 := " i love you!  "
	fmt.Printf("删除左边的空格要用strings.TrimeLeft, %q 删除左边空格后为: %v. \n", s4, strings.TrimLeft(s4, " "))
	fmt.Printf("删除右边的空格要用strings.TrimRight, %q 删除右边空格后为: %v. \n", s4, strings.TrimRight(s4, " "))
	fmt.Printf("删除两边的空格要用strings.Trim, %q 删除两边空格后为: %v. \n", s4, strings.Trim(s4, " "))

	// 分段string
	// 正常情况下strings.Split可以将string分开，但是如果要用空格或者是别的特殊字符来作为sep的话，还是要用Fields来做。
	fmt.Printf("%q按照%q来分割的话，就会是正常的答案： %v, 它的长度是: %v. \n", s4, "o", strings.Split(s4, "o"), len(strings.Split(s4, "o")))
	s4Slice := strings.Fields(s4)
	fmt.Printf("%q按照空格split之后会变成：%v, 它的长度变成了%v.\n", s4, s4Slice, len(s4Slice))

	// 遍历string, 这个没什么，就是for loop就好了
	for i, v := range s4 { // 这个for loop的时候，v就会是int32, 要把它换成string才可以； 或者直接用%q也可以
		fmt.Printf("The index is %v and the value is %v. \n", i, string(v))
	}

	// ************** array ********************** //
	// array基本信息, 长度，类型, index
	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("数组arr %v的长度为%d. \n", arr, len(arr))

	// find & index element in array, 并没有现成的func做这个事情，只能是iterate了
	// 找出array中的极值, sum等; 同样没有build in func, 也是只能iterate了;
	target := 3
	minVal, maxVal, preSum := arr[0], arr[0], 0
	for i, v := range arr {
		if v == target {
			fmt.Printf("The target %d found in index %d. \n", target, i)
		}
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
		preSum += v
	}

	// array的sort，甚至是自定义的sort
	arr2 := [...]int{2, 3, 1, 2, 39, 21, 2, 9}
	fmt.Printf("Go不能对array进行排序，要换成slice才可以，然后用sort.Ints()就可以了.\n Before sort: %v.\n", arr2)
	sort.Ints(arr2[:])
	fmt.Printf("After sort the array becomes: %v. \n", arr2)
	// 自定义的sort，有几种case是经常用的: 1.按照任意key来sort; 2.按照多个key来sort; 3.按照struct里某个field来sort;
	// 先定义一个二维数组，然后让它按照第二列来sort
	arr3 := [5][5]int{
		{1, 5},
		{-3, 100},
		{2, 1},
		{4, 0},
		{-5, 4},
	}
	sort.Ints(arr3[:])

	// array的增删查改
	// array的合并

	// ************** slice ********************** //
	// slice基本信息, 长度，类型, index
	// find & index element in slice
	// 找出slice中的极值, sum等;
	// slice的sort，甚至是自定义的sort
	// slice的增删查改
	// slice的合并


	// ************** map ********************** //
	// map的增删查改
	// map的sort，by keys, by values
	// map的keys和values
	// map的合并
}
