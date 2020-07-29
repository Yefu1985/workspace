package built_in

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)
/*
这是一号坑，主要是整理一下关于编码的一些知识点。
相信我们都碰到过ASCII码, UTF-8, UTF-16, UNICODE, 甚至乱码这样的东西，但是，却从来没有认真地整理过一下到底是什么鬼东西。
这里就好好把些概念过一遍：
http://cenalulu.github.io/linux/character-encoding/
上面这篇文章很好，可以比较系统的了解字符编码的概况。
只记住，Go是用UTF-8来作为中文的编码规则的，一个中文汉字需要3个byte来记录。
下面用例子来实现
 */
func Hole1Example1() {
	eChar := 'a' //Go里面没有char这个数据类型，用单引号的时候，其实是创建了一个rune，也就是一个int32
	eSize := unsafe.Sizeof(eChar)
	fmt.Printf("%q 的类型是 %T, 它的长度是%v个byte.\n", eChar, eChar, eSize) // 'a' 的类型是 int32, 它的长度是4个byte.

	cChar := '我' //同样，单引号搞的还是rune，还是一样的int32;
	cSize := unsafe.Sizeof(cChar)
	fmt.Printf("%q 的类型是 %T, 它的长度是%v个byte.\n", cChar, cChar, cSize) // '我' 的类型是 int32, 它的长度是4个byte.

	eString := "a" // 这里用的是string这个类型了，用了双引号了，长度就会是字符串的长度，也就是一个字符a，因为是英文字母，在UTF-8中，
				   // 一个英文字母用的是一个byte，所以，它的size就是一个byte，它的len()也是1，len返回的是byte的个数;
    eStrLen := len(eString)
    fmt.Printf("%q 的类型是 %T, 它有%d个byte.\n", eString, eString, eStrLen) //"a" 的类型是 string, 它有1个byte.

	cString := "爱" // 这里用的是string这个类型了，用了双引号了，长度就会是字符串的长度，也就是一个字符爱，因为是中文字符，在UTF-8中，
	// 一个中文字符用的是三个byte，所以，它的len()就是3;
	cStrLen := len(cString)
	fmt.Printf("%q 的类型是 %T, 它有%d个byte.\n", cString, cString, cStrLen) //"爱" 的类型是 string, 它有3个byte.

	//ok, 我们可以看一下这个汉字"爱"的真正的三个byte是啥
	fmt.Printf("%q的真正的byte是这样的: %0x \n", cString, cString) //"爱"的真正的byte是这样的: e788b1， 这是16进制，一共就是3个byte

	//那问题来了，如果我要算一下“我爱你”这句话中有多少个字符呢，那就岂不是要返回9个了？好像不太对了啊，Go里面怎么处理呢？
	//有个专门的function可以干这个事情utf8.RuneCountInString(), 这个会将中文汉字一个一个算出来;
	//下面两个例子就是纯中文和中英文混合的字符串的长度计算;
	cLongString := "我爱你"
	cLongStringSize := utf8.RuneCountInString(cLongString)
	fmt.Printf("%q 的长度为 %d.\n", cLongString, cLongStringSize) //"我爱你" 的长度为 3.

	mLongString := "我爱你Go"
	mLongStringSize := utf8.RuneCountInString(mLongString)
	fmt.Printf("%q 的长度为 %d.\n", mLongString, mLongStringSize) //"我爱你Go" 的长度为 5.
}
