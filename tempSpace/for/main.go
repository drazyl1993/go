/*
  for loop example
  1.go 语句结束不需要分号和python一样，
  2.代码块用{}规范，python用空格
  3.++自加运算。
*/

// package 似乎必须大头，就和c的头文件声明类似，package 的名字不一定要和路径一致
package main

func main() {
	var c int
	for i := 0; i < 10; i++ {
		c++
	}
	//注意go的for除了没有（），以及变量赋值语法不同，其他都和java一样
	println(c)
}
