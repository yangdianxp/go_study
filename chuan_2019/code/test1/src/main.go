package main

func main() {

}

/*
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
