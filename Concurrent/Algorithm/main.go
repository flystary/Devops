package main

import "fmt"

func main() {
	//sequence() //计算斐波那契数的算法

	//multiplication() //九九乘法表

	//var a , b = 24, 10
	//num := getMaximumCommonDivisor(a, b)
	//fmt.Println("a,b的最大公约数是：", num)
	//fmt.Println("a,b的最小公倍数是：", a*b/num)

	//a := "1000213320001"   //回文判断数字
	//isPalindrome(a)
	//fmt.Println(isPalindrome(a))

	//for i := 1; i < 1000; i++ {   //水仙花数
	//	if isDaffodilNum(i) {
	//		fmt.Println("属于水仙花数有: ", i)
	//	}
	//}

	//var a, b time.Duration //时间重置
	//a = time.Second
	//b = a * 3
	//fmt.Println(b)
	//a = time.Duration(500) * time.Millisecond
	//b = a * 1
	//fmt.Println(b)

	var str []int
	j := 0
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			str = append(str, i)
			j++
		}
	}
	fmt.Println(str)
}
