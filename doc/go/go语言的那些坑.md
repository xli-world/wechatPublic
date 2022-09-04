## 前言

本系列文章分为两篇，讲述go语言开发中常见的一些坑，帮助大家排坑解难。排名不分先后。

## 程序结构

### 短变量声明

*name := expression* 这样的声明方式为短变量声明，name的类型由expression的类型决定。

短变量声明有以下几个特性：

1. 多个变量可以以短变量声明的方式声明和初始化：

```go
i, j := 0, 1
f, err := os.Open(name)
```

2. 短变量声明不需要声明所有在左边的变量，在**同一词法块**中已经存在变量的情况下，短变量声明的行为和赋值操作一样，外层的声明将被忽略。

3. 短变量声明最少声明一个新变量，否则，代码编译将无法通过。

这里常见的问题是对**同一词法块**的理解有误，for、if、switch语句都会创建自己的词法块：

1. for循环创建了两个词法块，一个是循环体本身的显式块，以及一个隐式块，它包含了一个闭合结构，其中就有初始化语句中声明的变量，如变量i。隐式块中声明的变量的作用域包括条件、后置语句（i++），以及for语句体本身。

```go
for i := 0; i < 5; i++ {
  x := i
  fmt.Println(x)
}
// 变量i在for循环的开始创建，变量x在循环的每次迭代中创建。
```

2. if语句创建了一个词法块，包含初始化语句中声明的变量。if语句中声明的变量在之后的else-if语句及else语句中是可见的，else-if语句中声明的变量在之后的else-if语句中是可见的。

```go
if x := f(); x == 0 {
  fmt.Println(x)
} else if y := g(x); x == y {
  fmt.Println(x, y)  // 上个if语句的初始化部分中声明的变量x在这个语句中是可见的
} else {
  fmt.Println(x, y)
}
fmt.Println(x, y) // 编译错误：x、y在这里不可见
```

3. switch语句中条件对应一个块，每个case语句体对应一个块。

```go
switch 
```

代码示例如下：

```go
filename := "test"
f, err := os.Open(filename) // 声明了两个新变量f，err
_ = f
fmt.Println(err) // open test: no such file or directory

if m ,err := os.Readlink(filename);err != nil { // err为新声明的局部变量
	fmt.Println(err) // readlink test: no such file or directory
	_ = m
}
fmt.Println(err) // open test: no such file or directory

for n, err := os.Readlink(filename); err != nil; { // err为新声明的局部变量
	fmt.Println(err) // readlink test: no such file or directory
	_ = n
	break
}
fmt.Println(err) // open test: no such file or directory

p, err := "new variable", errors.New("new error") // 未声明新变量err，而是将errors.New("new error")赋值给了之前的err变量
fmt.Println(err) // new error
```

### 指针

1. 指针的值是一个变量的地址。不是所有的值都有地址，但所有的变量都有。

```go
var j = 1
var k = &j
var i = &1 // 编译错误：Cannot take the address of '1'
```

2. 指针是可比较的，两个指针当且仅当指向同一个变量或者**两者都是nil**的情况下才相等。这里常见的误区是指针相等不一定指向同一个变量，也有可能它们都是nil。

```go
type A struct {
}

func main() {
	var a *A
	var b *A
	fmt.Println(a == b) // a,b都为nil，比较的结果为true
}
```

## 基本数据

### 无符号整数

无符号整数一般用于位运算符和特定算术运算符，如实现位集时，解析二级制格式的文件，或散列和加密，极少用于表示非负值。

表示非负值时容易忘记它的非负特性：

```go
medals := []string{"gold", "silver", "bronze"}
var length uint32 = 3
for i := length - 1; i >= 0; i-- {
	fmt.Println(medals[i])
}
```

输出结果：

```
bronze
silver
gold
panic: runtime error: index out of range [4294967295] with length 3
```

### 浮点数

Go具有两种大小的浮点数float32和float64。十进制下，float32的有效数字大约是6位，float64的有效数字大约是15位，绝大多数情况下，应优先选用float64，因为除非格外小心，否则float32的运算会迅速累积误差。另外，float32能精确表示的正整数范围有限：

```go
var f float32 = 16777216 // 1<< 24
fmt.Println(f == f+1) // true
```

### NAN、+Inf、-Inf

Go中有几个特殊的整数数值：

1. +Inf：正无穷大，表示超出最大许可值的数；
2. -Inf：负无穷大，表示除以0的商；
3. NAN：not a number，表示数学上无意义的运算结果

```go
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN
```

在数学运算中，倾向于将NAN当作信号值（使用math.IsNaN()来检查），而不是直接判断具体的计算结果是否是NAN，因为和NAN的比较总不成立（除了!=，它总是与==相反）

```go
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z)  // 0 -0 +Inf -Inf NaN
fmt.Println(1/z == math.Inf(1))  // true
fmt.Println(z/z == math.NaN())  // false
fmt.Println(math.IsNaN(z/z)) // true
```

### 无类型常量

虽然常量可以是任何基本数据类型，如int或float64，但有时候定义常量的时候并不会给它设置具体的类型。编译器将这些从属类型待定的常量表示成某些值，这些值比基本类型的数字精度更高，且算数精度高于原生的机器精度。可以认为它们的精度至少达到256位。从属类型待定的常量共有6种，分别是无类型布尔、无类型整数、无类型文字符号、无类型浮点数、无类型复数、无类型字符串。

借助于退出确定从属类型，无类型常量不仅能暂时维持更高的精度，与类型已确定的常量相比，它们还能写进更多表达式而无需转换类型：

```go
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776               （超过1 << 32）
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424      （超过1 << 64）
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(YiB/ZiB) // 1024
}
```

变量声明（包括短变量声明）中，假设没有显式指定类型，无类型常量会隐式转换成该变量的默认类型：

```go
i := 0
r := `\000`
f := 0.0
c := 0i
fmt.Printf("%T %T %T %T", i, r, f, c) // int string float64 complex128
```

## 复合数据类型

### 数组

数组的几个特性：

1. 数组的长度是数组类型的一部分，所以[3]int和[4]int是两种不同的数组类型。
2. 数组的长度必须是常量表达式，也就是说，这个表达式的值在程序编译时就可以确定。
3. 如果一个数组的元素类型是可比较的，那么这个数组也是可比较的。

### slice

slice操作符s[i:j]（其中0<=i<=j<=cap(s)）可以创建一个新的slice，这个新的slice引用了序列s中从i到j-i索引位置的所有元素。

如果slice的引用超过了被引用对象的容量，即cap(s)，那么会导致panic；但是如果slice的引用只是超出了被引用对象的长度，即len(s)，这种情况下会访问到原slice中不存在的元素：

```go
var months []string = []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
summer := months[3:6]
fmt.Println(summer[:5]) // [7月 8月 9月 10月 11月]
```

## 函数

函数实参是按值传递的，所以函数接收到的是每个实参的副本，修改函数的形参变量并不会影响到调用者提供的实参。然后，如果提供的实参包含引用类型，比如**指针、slice、map、函数或者通道**，那么当函数使用形参变量时就有可能会间接的修改实参变量。

### defer

这里分享一个小技巧，defer不仅可以用来清理“战场”，还可以配合闭包使用，在函数“入口”、“出口”处设置调试行为：

```go
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...这里是一些处理...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
```

输出结果如下：

```
2022/07/10 23:18:22 enter bigSlowOperation
2022/07/10 23:18:32 exit bigSlowOperation (10.000840033s)
```

还有个常见的错误是，在for语句中使用defer，这个可能会导致无法释放资源的问题：

```go
for _, filename := range filenames {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // 注意：可能会用尽文件描述符
	// ...处理文件f...
}
```

一种解决方式是将循环体（包括defer语句）放到另一个函数里，每次循环迭代都会调用文件关闭函数：

```go
for _, filename := range filenames {
  if err := doFile(filename); err != nil {
		return err
	}
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...处理文件f...
}
```

### recovery

很多人都推荐在一些重要或者易出错的函数的入口用defer recovery的方式捕获未知异常，但是有些异常是通过这种方式捕获不到的，比如，内存耗尽使得Go运行时发生严重错误而直接终止进程。

## 一些小知识

### fallthrough

Go里面switch默认每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 可以使用fallthrough强制执行后面的case代码。

fallthrough有以下几个特性：

1. fallthrough不能用在switch的最后一个分支
2. fallthrough到下一个case块时，**不执行case匹配检查！**

代码示例如下：

```go
package main

import "fmt"

func main() {
	var i = 5
	switch {
	case i <= 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case i <= 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case i <= 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case i <= 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case i <= 3:
		fmt.Println("The integer was <= 3")
	default:
		fmt.Println("default case")
	}
}
```

输出结果如下，i <=3 不成立，但仍然输出了"The integer was <= 3"

```
The integer was <= 5
The integer was <= 6
The integer was <= 7
The integer was <= 3
```

