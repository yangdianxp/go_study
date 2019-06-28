package main

import "fmt"

func main()  {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
{
func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers) * 2))

	copy(numbers1, numbers)
	printSlice(numbers1)
}

{


	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println("numbers[:3] ==", numbers[:3])
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	number2 := numbers[:2]
	printSlice(number2)

	number3 := numbers[2:5]
	printSlice(number3)

}

{
	var numbers []int
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切处是空的")
	}
}

{
func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

}

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

{
type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

	rect := Rect{1, 2, 3, 4}
	fmt.Println(rect.Area())
}

{
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func changeBook(book *Books) {
	book.title = "book1_change"
}

	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.book_id = 1
	changeBook(&book1)
	fmt.Println(book1)
}

{
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBook( book *Books)  {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(&Book1)
	printBook(&Book2)
}

{
func printBook( book Books)  {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
}

{
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

{
type Books struct {
	title string
	author string
	subject string
	book_id int
}

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

{
	var a int = 10
	var ip *int
	fmt.Printf("变量的地址: %x\n", &a)
	fmt.Println("变量的地址: ", &a)
	ip = &a
	fmt.Println("ip 变量存储的指针地址:", ip)
	fmt.Println("ip 变量存储的指针地址的值:", *ip)
	fmt.Println("ip 变量存储的指针地址的地址:", &ip)
	var ptr *int
	if (ptr != nil) {
		if (ip != nil) {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip不是空指针")
		}else {
			fmt.Println("ptr不是空指针")
			fmt.Println("ip是空指针")
		}
	} else {
		if(ip != nil){
			fmt.Println("ptr是空指针")
			fmt.Println("ip不是空指针")
		}else{
			fmt.Println("ptr是空指针")
			fmt.Println("ip是空指针")
		}
	}

{
	var a int = 1
	var ptr1 *int = &a
	var ptr2 **int = &ptr1
	var ptr3 ***int = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a:", a)
	fmt.Println("ptr1", ptr1)
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("*ptr1", *ptr1)
	fmt.Println("**ptr2", **ptr2)
	fmt.Println("**(*ptr3)", ***ptr3) // 也可以写作：***ptr3
}

{
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a
	pptr = &ptr
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

{
const MAX int = 3

	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

{
const MAX int = 3

	a := []int{10,100,200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

{
	var ptr *int
	fmt.Printf("ptr 的值为 : %x %d\n", ptr, *ptr)
}

{
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量存储的指针地址: %x\n", ip)
	fmt.Printf("ip 变量的值: %d\n", *ip)
}

{
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)
}

{
	a := 1690
	b := 1700
	c := a * b
	fmt.Println(c)
	fmt.Println(float64(c) / 1000000)
}

{
	a := 1.69
	b := 1.7
	c := a * b
	fmt.Println(c)
}

{
	var balance = []int {1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage( balance, 5)

	fmt.Printf("平均值为: %f ", avg)

func getAverage(arr []int, size int) float32  {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)

	return avg
}

}

{
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6}, {4,8}}
	var i,j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println(a)
}

{
	var n [10]int
	var i,j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

{
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance, balance1)
}

{
type Circle struct {
	radius float64
}

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
}

{
	add_func := add(1, 2)
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(0, 0))
	fmt.Println(add_func(2, 2))

func add(x1, x2 int) func(x3 int, x4 int)(int, int, int)  {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1+x2, x3+x4
	}
}
}

{
	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

func add(x1, x2 int) func()(int, int)  {
	i := 0
	return func() (int, int) {
		i++
		return i, x1+x2
	}
}
}

{
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

{
type cb func(int) int

	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x: %d\n", x)
		return x
	})

func testCallBack(x int, f cb)  {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x: %d\n", x)
	return x
}
}


{
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}

{
func swap(x *int, y *int)  {
	*x, *y = *y, *x
}

	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)
}

{
func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swap(a, b)

	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)
}

{
func swap(x, y string) (string, string)  {
	return y, x
}

	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

{
func max(num1, num2 int) int  {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
}

{
	for true {
		fmt.Printf("这是无限循环。\n")
	}
}

{
	var a int = 10
	LOOP: for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

{
	var a int = 10
	for a < 20 {
		if a == 15 {
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

{
	var a int = 10
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			break
		}
	}
}

{
	for m:=1; m<10; m++ {
		for n:=1; n<=m; n++ {
			fmt.Printf("%dx%d=%d ", n,m,m*n)
		}
		fmt.Println()
	}
}

{
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
			if (i%j==0) {
				break;
			}
		}
		if j > (i/j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}

{
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

{
func Chann(ch chan int, stopCh chan bool)  {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Recvice", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

{
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, "to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

{
	var a int = 2
	var b int = 3
	var c int = 4
	switch {
	case a > 2:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case a == 2:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case b < 3:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case c == 4:
		fmt.Println("4、case 条件语句为 true")
	case c > 4:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

{
func test(int) float64  {
	return 2.2
}

	var x interface{}
	x = "aaa"

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

{
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70: grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C" :
		fmt.Printf("良好")
	case grade == "D" :
		fmt.Printf("及格")
	case grade == "F" :
		fmt.Printf("不及格")
	default:
		fmt.Printf("差")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

{
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 的值为 : %d\n", a)
	fmt.Printf("b 的值为 : %d\n", b)
}

{
	var age int = 23
	if age == 25 {
		fmt.Println("true")
	} else if age < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}
}

{
	var count int
	var flag bool
	count = 1
	for count < 100 {
		count++
		flag = true
		for tmp:=2; tmp<count; tmp++ {
			if count%tmp == 0 {
				flag = false
				break
			}
		}

		if flag == true {
			fmt.Println(count, "素数")
		} else {
			continue
		}
	}
}

{
	var a int = 100

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

{
	var a int = 10

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}


{
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e );

	e = ((a + b) * c) / d
	fmt.Printf("((a + b) * c) / d 的值为 : %d\n", e );

	e = (a + b) * (c / d)
	fmt.Printf("(a + b) * (c / d) 的值为 : %d\n", e );

	e = a + (b * c) / d
	fmt.Printf("a + (b * c) / d 的值为 : %d\n", e );
}

{
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a )
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b )
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c )
	fmt.Printf("第 4 行 - ptr 变量类型为 = %T\n", ptr )

	ptr = &a
	fmt.Printf("a 的值为 %d\n", a )
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

{
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - = 运算符实例， c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例， c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例， c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例， c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例， c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6 行 - <<= 运算符实例， c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例， c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例， c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例， c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例， c 值为 = %d\n", c)
}

{
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

{
	var a bool = true
	var b bool = false
	if ( a && b ) {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if ( a || b ) {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	a = false
	b = true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if ( !(a && b) ){
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

{
var a int = 21
var b int = 10

if ( a == b ) {
fmt.Printf("第一行 - a 等于 b\n")
} else {
fmt.Printf("第一行 - a 不等于 b\n")
}
if ( a < b ) {
fmt.Printf("第二行 - a 小于 b\n")
} else {
fmt.Printf("第二行 - a 不小于 b\n")
}
if ( a > b ) {
fmt.Printf("第三行 - a 大于 b\n")
} else {
fmt.Printf("第三行 - a 不大于 b\n")
}

a = 5
b = 20
if ( a <= b ) {
fmt.Printf("第四行 - a 小于等于 b\n")
}
if ( b >= a ) {
fmt.Printf("第五行 - b 大于等于 a\n")
}
}

{
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}

{
	const (
	i = 1<<iota
	j = 3<<iota
	k
	l
)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

{
const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

{
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)
	println(a, b, c)
}

{
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为 ： %d", area)
	println()
	println(a, b, c)
}

{
_, numb, strs := numbers()
	fmt.Println(numb, strs)
func numbers()(int, int, string)  {
	a, b, c := 1, 2, "str"
	return a, b, c
}
}

{
	var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

}

{
	f := "Runoob"
	fmt.Println(f)
}

{
	var intVal int
	intVal, intVal1 := 1, 2
	fmt.Println(intVal, intVal1)
}

{
var d = true
	fmt.Println(d)
}

{
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}

{
	var a = "RUNOOB"
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)
}

{
var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
}

fmt.Println("Google" + "Runoob")

fmt.Println("Hello World!")


*/
