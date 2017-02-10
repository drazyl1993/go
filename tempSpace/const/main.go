/*
   const example
*/
package main

const x, y int = 123, 0x22
const str = "hello go"
const c = '我'

const (
	//错误实例：test,由于test上方没有常量，所以编辑的时候，保存都失败。。
	//test
	i, f = 1, 0.1314
	i1, f1
	// if,f1 和i和f 完全一样.暂时不明白为什么有这个设计
	b = false
	d
	//如果在常量中声明一个常量，不指定类型和初始化值，则，这个常量会和上一行代码的常量类型和值都保持一致，是不是就是说地址不一致？有待认证。
)

func main() {
	/*
		println(&x, x)
		println(&y, y)
		println(&str, str)
		println(&c, c)
		println(&i, i)
		println(&f, f)
		println(&b, b)
		这里发现一个问题，用&常量名，无法获得常量的内存地址
	*/
	println(x)
	println(y)
	println(str)
	println(c)
	println(i)
	println(f)
	println(b)
	println(d)
	println(i1, f1)
}
