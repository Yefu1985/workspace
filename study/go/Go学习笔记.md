<p style='text-align: left;'>

# Go学习笔记

## 前言
**这不是一篇只关于Go的学习笔记**

**这不是一篇只关于Go的学习笔记**

**这不是一篇只关于Go的学习笔记**

本是ME的PhD，鬼使神差来了湾区，跟写代码一毛钱关系都没有，但是看着湾区的房价和消费，嗯，没得办法了，为五斗米折腰，按照湾区常规剧本，转行做码农了。

转行只做了一件事情，就是刷题，最少算是把数据结构和算法搞得七七八八了，当然这些东西对于转行工作来说用处实在是有限。自转行以来，一直都是在广告组的backend岗位搬砖，从0
开始（或许说从负数开始更合适），学习成长，一路上有少许心得和体会，都希望在这篇笔记中有所记录。IT行业的知识和技能更新迭代太快，学习资料太多，对于像我这样的新手而言，往往不知所措，面对一大堆的工具和术语，脑子里一片混沌，我TMD
该从哪个下手呢？随便挑一个开始闷头学，发现看着看着就看不懂了，不知道在说什么。无论是各种视频也好，还是各种document也好，都是在用专业名词来解释专业名词，额，臣妾还是看不懂啊。举个例子体会一下, 下面是对Restful api
的wikipedia定义：
* Representational state transfer (REST) is a software architectural style that defines a set of constraints to be used for creating Web services. 
* 表现层状态转换（英语：Representational State Transfer，缩写：REST）是Roy Thomas Fielding博士于2000年在他的博士论文中提出来的一种万维网软件架构风格，目的是便于不同软件/程序在网络（例如互联网）中互相传递信息。表现层状态转换是根基于超文本传输协议（HTTP）之上而确定的一组约束和属性，是一种设计提供万维网络服务的软件构建风格。

不管是中文还是英文，真心觉得这样的定义对于小白来说，真的是太费劲了，看了半天还是不知道这到底是个啥，究竟能干啥。

网上各种各样的辅导班，五花八门的学习资料，真的是琳琅满目，反而不知道该如何下手了。我该学啥呢，哪些资料是适合新手的呢，哪些资料是适合老鸟的呢，其实并没有太多头绪，最后只会发现，嗯，时间投入了挺多，可是好像学习效果挺一般，那些不会的，还是不会。而且，和其它学科相比，IT行业最大的特点是要动手去做，所有的技能都不是学会的，最终目的都是为了最后的实践，都应该是做会的。

在这篇学习笔记中，会详细记录我在每一个学习知识点上用到的学习资料（视频也好，document
也好，网页也好），这些会是用来学习的基石，但是不要强求和较真，最关键的还是落实到做上面去。看不懂也没关系，很正常，不要觉得绝望，先略过去，做点事情，做的过程中发现问题了，在解决问题的时候再回头去学习，理解往往更深刻。

OK
, 在阅读这篇笔记的时候，希望你至少刷过一些题，最少对数据结构和算法有个基本认识，至少囫囵吞枣的因为面试而准备过一些系统设计，知道里面的一些高大上的名词术语。这篇笔记我也不知道自己会写多久，尽力而为，争取每一个写出来的知识点都能尽力理解。

目前这篇笔记包含了以下内容：
* 学习工具
* Go的基本认识
* 数据结构
* Go内置package
* 

## Part 0: 工具篇
工欲善其事，必先利其器！

一句话，如果不差钱，就上Mac，这是和以后的工作环境最接近的平台了。

如果不喜欢Mac，那就用windows就好了，不用太高大上的机器。我就是用的T560，8G内存，独立显卡都没。有时候在学习的过程中会将Windows换成Ubuntu来尝试Linux
环境，或者也可以选择安装虚拟机来进行学习，这个影响不大。只是，有时候在Linux下面安装工具和设置环境的时候那个费劲啊，哎，想着还是有个Mac就好了。

如果你长期使用Mac，并且经常使用terminal的话，那问题不大，可以跳过这个part。否则的话，可以稍微了解一下Linux的内容，比如说用户权限，比如说一些基本的terminal的操作，比如说一些基本的设置，一些基础的shell
教程（看看就可以了，如果用的多的话，再去深入学习）。 这些是我以前看过的资料，可以根据自身情况来快速过一遍。
* [Linux基础教程(youtube视频)](https://www.youtube.com/watch?v=4hUemuE_eU4&list=PLmOn9nNkQxJFYoQN9mBcr18fU1UAuWps0)
* [MacOS Terminal常用命令](https://juejin.im/post/5b03e16c6fb9a07a9b364a90)
* [Shell基础教程](https://www.runoob.com/linux/linux-shell.html)

多余的真的不用看了，除非你工作中要大量运用，否则看完这些，足够应付你平时的工作需求了。

千万记住，熟能生巧，学了就一定要用。举个最最最简单的例子，以前用鼠标右键做的复制粘贴文件的动作不要做了，去用terminal
完成，一些简单的文本编辑，也可以尝试用Vim来完成，如果不会了，就去Google，这样才算真的是掌握了。

## Part I: Go的初体验
如果你学过一点点的面向对象的编程语言，比如说C++, JAVA
之类的，那这个部分应该会很快。不需要在第一遍看的时候就全部记下来，还是那句话，这个不是看会的，一定是做多了做会的。就像开车一样，初体验嘛，拿第一滴血就好了，你也不可能一上来就是老司机，对吧。看这个的目的只是为了有一个最初的认识，对Go
的一些Syntax有基本的第一感受。推荐使用下面两个学习资料。
1. The official [Go Tutorial](https://tour.golang.org/): 如果你有编程经验的话，这个真的挺简单，挺快的，我只学了一下午就学完了。
2. The official [Go Documentation](https://golang.org/pkg): 官方文档，亲，一定要保存在收藏夹里面。

其实所有的官方文档给我的感觉一直都不是太好，总觉得很晦涩难懂，当然最最主要的原因是因为自己太挫。但是，读document
确实是你必须掌握的一个技能，我的经验和建议是：在最初学习的时候，看不懂很正常，不要沮丧，不要怀疑自己的智商，尝试去Google搜索一下别的网页怎么解释的，应该能找到让自己看懂的网页，看懂了之后，再回头去看这部分对应的document
，看看在document里面是怎么去解释和描述的，体会document作者对这部分的写法，慢慢看多了，就会发现，自己去Google的频率越来越低了，慢慢地就会发现，嗯，Document
也是在说人话。只是，这个过程会有点漫长，但是不要急，坚持做就好了。

其实，泛泛地看完了第一个，就完全可以上手做一些东西开始深入学习了。 Let's Go!

## Part II: 数据结构
如前所述，希望你在看这篇笔记的时候是刷过一些题的，那么，一些基本的数据结构和算法应该不在话下吧。如果这部分还欠缺的话，有两个学习资料可以推荐：
* [九章的算法基础班](https://www.jiuzhang.com/): 特此声明，我没拿辅导班的回扣，辅导班也没有那么傻逼会给我这样的弱鸡回扣，这是我上过的为数不多的辅导班，对于刷题和算法而言，觉得性价比极高。
* Data Structure for Python/Java: 这是一本电子书，网上随便都能找到，根据自己需求，用Python或Java都可以，里面会详细介绍所有的数据结构，比如说刷题中经常用到的`stack`, `queue`
, `heap`, `linked list`, `hashmap`, etc. 

OK, 以上资料，如果你刷过题或者正在刷题的话，应该都不会陌生的。还是那句话，抓紧看完就可以上手了。废话不多说，干就完了。

### BuiltIn Data Structures
按照我个人的理解，Go里面的数据类型(data type)大概就是这么个结构：
* 基本类型：可以说是basic，也可以理解是primitive，简单说，这个类型里面存的就是真正的data,又可以分为以下几个分支
    * 数字类numbers: 有`uint`, `int`; 这两个是默认的，还有`int(8, 16, 32, 64)`, `uint(8, 16, 32, 64)`; 有`float32`, `float64`; 还有复数`complex`
    类型，我没用过，所以也没在意过；
    * 字节`byte`和`rune`：**Go里面没有`char`类型**， `rune`本质上是一个`int32`;
    * 字符串`string`：这个和JAVA还真不太一样，在Go里面竟然是基本类型;
    * 布尔型`bool`： 正常的true和false就是了; 
    * 数组`array`：里面的数据要是一致的类型，定义之后，大小固定了，不能修改, 每个元素称为element;
    
* 引用类型: 就是reference type, 就是说这里面存的不是真实的data，而是地址，指向真实的data而已
    * 切片`slice`： 这个很像Python里面的List，就是截取了原始的array里面的一组数据作为切片
    * 字典`map`：这就是常规的key-value的那种数据结构;
    * 结构体`struct`: 这个是C语言体系的概念吧，不过跟JAVA里面的`class`也很像;只是`struct`里面不能有方法，只能通过`func`和`receiver`的方式来给`struct`添加方法;
    * 指针`pointer`：这个是C语言体系才有的概念，回头慢慢聊;
    * 通道`channel`： 这是处理并发才需要用到的概念，回头如果工作需要再去学习;

以上就是Golang里面的数据类型，里面有无数个坑可以跳。嗯，我尽可能的把我跳过的坑写出来。

* 1号坑：[Go里面的字符，字符串，编码等一堆乱麻](https://github.com/Yefu1985/workspace/blob/e65b71d4a1ffaf2363c3cdce88e2ec9d0c1cdc85/study/go/data_structure/built_in/hole_1.go#L17)
* 2号坑：[数组`array`和切片`slice`](https://github.com/Yefu1985/workspace/blob/80741d708177b1b3f4b66349cadc180f97cf6fcb/study/go/data_structure/built_in/hole_2.go#L5)
* 3号坑：[指针`pointer`]()
* 4号坑：[结构体`struct`]()
* 5号坑：[接口`interface`]()

### Other Common Data Structures

## Part III Basic BuiltIn Packages
### ioutils
### time & date
### data format: json, csv, avro, protobuf
### logger
### database:sql
### regex
### testing
### modules

## Part IV Other Utilities
### Package and Project Management
### Cloud Components: AWS, GCP
### Container:Docker
### Dashboard:Grafana



</p>
