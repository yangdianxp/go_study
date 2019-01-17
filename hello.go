package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

//import "fmt"

//func main() {
//	// 这里我们定义了一个可以存储整数类型的带缓冲通道
//	// 缓冲区大小为2
//	ch := make(chan int, 2)

//	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
//	// 而不用立刻需要去同步读取数据
//	ch <- 1
//	ch <- 2

//	// 获取这两个数据
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}

//import "fmt"

//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // 把 sum 发送到通道 c
//}

//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}

//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	x, y := <-c, <-c // 从通道 c 中接收

//	fmt.Println(x, y, x+y)
//}

//import (
//	"fmt"
//	"time"
//)

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

//func main() {
//	go say("world")
//	say("hello")
//}

//import (
//	"fmt"
//)

//// 定义一个 DivideError 结构
//type DivideError struct {
//	dividee int
//	divider int
//}

//// 实现 `error` 接口
//func (de *DivideError) Error() string {
//	strFormat := `
//    Cannot proceed, the divider is zero.
//    dividee: %d
//    divider: 0
//`
//	return fmt.Sprintf(strFormat, de.dividee)
//}

//// 定义 `int` 类型除法运算的函数
//func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
//	if varDivider == 0 {
//		dData := DivideError{
//			dividee: varDividee,
//			divider: varDivider,
//		}
//		errorMsg = dData.Error()
//		return
//	} else {
//		return varDividee / varDivider, ""
//	}

//}

//func main() {

//	// 正常情况
//	if result, errorMsg := Divide(100, 10); errorMsg == "" {
//		fmt.Println("100/10 = ", result)
//	}
//	// 当被除数为零的时候会返回错误信息
//	if _, errorMsg := Divide(100, 0); errorMsg != "" {
//		fmt.Println("errorMsg is: ", errorMsg)
//	}

//}

//import (
//	"fmt"
//)

//type Phone interface {
//	call()
//}

//type NokiaPhone struct {
//}

//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}

//type IPhone struct {
//}

//func (iPhone IPhone) call() {
//	fmt.Println("I am iPhone, I can call you!")
//}

//func main() {
//	var phone Phone

//	phone = new(NokiaPhone)
//	phone.call()

//	phone = new(IPhone)
//	phone.call()

//}

//import "fmt"

//func Factorial(n uint64) (result uint64) {
//	if n > 0 {
//		result = n * Factorial(n-1)
//		return result
//	}
//	return 1
//}

//func main() {
//	var i int = 15
//	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
//}

//import "fmt"

//func main() {
//	var countryCapitalMap map[string]string /*创建集合 */
//	countryCapitalMap = make(map[string]string)

//	/* map插入key - value对,各个国家对应的首都 */
//	countryCapitalMap["France"] = "Paris"
//	countryCapitalMap["Italy"] = "罗马"
//	countryCapitalMap["Japan"] = "东京"
//	countryCapitalMap["India "] = "新德里"

//	/*使用键输出地图值 */
//	for country := range countryCapitalMap {
//		fmt.Println(country, "首都是", countryCapitalMap[country])
//	}

//	/*查看元素在集合中是否存在 */
//	captial, ok := countryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
//	/*fmt.Println(captial) */
//	/*fmt.Println(ok) */
//	if ok {
//		fmt.Println("美国的首都是", captial)
//	} else {
//		fmt.Println("美国的首都不存在")
//	}
//}

//import "fmt"

//func main() {
//	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
//	nums := []int{2, 3, 4}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:", sum)
//	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
//	for i, num := range nums {
//		if num == 3 {
//			fmt.Println("index:", i)
//		}
//	}
//	//range也可以用在map的键值对上。
//	kvs := map[string]string{"a": "apple", "b": "banana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
//	for i, c := range "go" {
//		fmt.Println(i, c)
//	}
//}

//import "fmt"

//func main() {
//	/* 创建切片 */
//	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
//	printSlice(numbers)

//	/* 打印原始切片 */
//	fmt.Println("numbers ==", numbers)

//	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
//	fmt.Println("numbers[1:4] ==", numbers[1:4])

//	/* 默认下限为 0*/
//	fmt.Println("numbers[:3] ==", numbers[:3])

//	/* 默认上限为 len(s)*/
//	fmt.Println("numbers[4:] ==", numbers[4:])

//	numbers1 := make([]int, 0, 5)
//	printSlice(numbers1)

//	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
//	number2 := numbers[:2]
//	printSlice(number2)

//	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
//	number3 := numbers[2:5]
//	printSlice(number3)

//}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//import "fmt"

//func main() {
//	var numbers []int

//	printSlice(numbers)

//	if numbers == nil {
//		fmt.Printf("切片是空的")
//	}
//}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//import "fmt"

//func main() {
//	var numbers = make([]int, 3, 5)

//	printSlice(numbers)
//}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//import "fmt"

//func getSequence() func() int {
//	i := 0
//	fmt.Println(&i)
//	return func() int {
//		i += 1
//		//fmt.Println(&i)
//		return i
//	}
//}

//func set_func() {
//	i := 0
//	fmt.Println(&i)
//}

//func main() {
//	/* nextNumber 为一个函数，函数 i 为 0 */
//	nextNumber := getSequence()
//	set_func()
//	set_func()

//	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())

//	/* 创建新的函数 nextNumber1，并查看结果 */
//	nextNumber1 := getSequence()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//}

//import "fmt"

//func swap(x, y string) (string, string) {
//	return y, x
//}

//func main() {
//	a, b := swap("Mahesh", "Kumar")
//	fmt.Println(a, b)
//}

//import "fmt"

//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//	var ret int

//	/* 调用函数并返回最大值 */
//	ret = max(a, b)

//	fmt.Printf("最大值是 : %d\n", ret)
//}

///* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 定义局部变量 */
//	var result int

//	if num1 > num2 {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}

//import "fmt"

//func main() {
//	for true {
//		fmt.Printf("这是无限循环。\n")
//	}
//}

//import "fmt"

//func main() {
//	var a int = 4
//	var b int32
//	var c float32
//	var ptr *int

//	/* 运算符实例 */
//	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
//	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
//	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

//	/*  & 和 * 运算符实例 */
//	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
//	fmt.Printf("a 的值为  %d\n", a)
//	fmt.Printf("*ptr 为 %d\n", *ptr)
//}

//import "fmt"

//func main() {

//	var a int = 21
//	var b int = 10
//	var c int

//	c = a + b
//	fmt.Printf("第一行 - c 的值为 %d\n", c)
//	c = a - b
//	fmt.Printf("第二行 - c 的值为 %d\n", c)
//	c = a * b
//	fmt.Printf("第三行 - c 的值为 %d\n", c)
//	c = a / b
//	fmt.Printf("第四行 - c 的值为 %d\n", c)
//	c = a % b
//	fmt.Printf("第五行 - c 的值为 %d\n", c)
//	a++
//	fmt.Printf("第六行 - a 的值为 %d\n", a)
//	a = 21 // 为了方便测试，a 这里重新赋值为 21
//	a--
//	fmt.Printf("第七行 - a 的值为 %d\n", a)
//}

//import "fmt"

//const (
//	i = 1 << iota
//	j = 3 << iota
//	k
//	l
//)

//func main() {
//	fmt.Println("i=", i)
//	fmt.Println("j=", j)
//	fmt.Println("k=", k)
//	fmt.Println("l=", l)
//}

//import "fmt"

//func main() {
//	const (
//		a = iota //0
//		b        //1
//		c        //2
//		d = "ha" //独立值，iota += 1
//		e        //"ha"   iota += 1
//		f = 100  //iota +=1
//		g        //100  iota +=1
//		h = iota //7,恢复计数
//		i        //8
//	)
//	fmt.Println(a, b, c, d, e, f, g, h, i)
//}

//import "unsafe"

//const (
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)

//func main() {
//	println(a, b, c)
//}

//func main() {
//	const LENGTH int = 10
//	const WIDTH int = 5
//	var area int
//	const a, b, c = 1, false, "str" //多重赋值

//	area = LENGTH * WIDTH
//	fmt.Printf("面积为 : %d", area)
//	println()
//	println(a, b, c)
//}

//var x, y int
//var (  // 这种因式分解关键字的写法一般用于声明全局变量
//    a int
//    b bool
//)

//var c, d int = 1, 2
//var e, f = 123, "hello"

////这种不带声明格式的只能在函数体中出现
////g, h := 123, "hello"

//func main(){
//    g, h := 123, "hello"
//    println(x, y, a, b, c, d, e, f, g, h)
//}

/*
var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

func main(){
    println(a, b, c)
}
*/

//import "fmt"

//func main(){  // 错误，{ 不能在单独的行上
//    fmt.Println("Hello, World!")
//}
