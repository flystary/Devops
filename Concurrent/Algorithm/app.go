package main

import (
	"fmt"
	"math"
	"time"
)

func sequence() {
	const n = 40
	starttime := time.Now()
	fibN := fib(n)
	endtime := time.Now()
	costTime := endtime.Sub(starttime)
	fmt.Println(costTime)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func multiplication() {
	start := time.Now()
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
	//end := time.Now()
	tc := time.Since(start)
	//tc := end.Sub(start)
	fmt.Printf("耗时是%v", tc)
}

func getMaximumCommonDivisor(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}

func isPalindrome(a string) bool {
	var i, j int
	var b bool
	for i = 0; i < len(a)/2-1; i++ {
		j = len(a) - 1
		if a[i] != a[j] {
			b = false
		}
		b = true
		j--
	}
	return b
}
func isDaffodilNum(num int) bool {
	a := num / 100
	b := (num / 10) % 10
	c := num % 10
	result := a*a*a + b*b*b + c*c*c
	if num == result {
		return true
	}
	return false
}

func isPrime(i int) bool {
	for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func wage() {
	var n, salary, sum, a int
	var err error

	fmt.Print("请属于你的工龄：")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("请输入你的基本工资：")
	_, err = fmt.Scanln(&salary)
	if err != nil {
		fmt.Println(err)
	}
	if n >= 0 && n < 1 {
		a = 200
	} else if n >= 1 && n < 3 {
		a = 500
	} else if n >= 3 && n < 5 {
		a = 1000
	} else if n >= 5 && n < 10 {
		a = 2500
	} else if n >= 10 && n < 15 {
		a = 5000
	}
	sum = salary + a
	fmt.Printf("您目前工作了%d年，基本工资为%d元,应涨工资%d元,涨后工资%d元", n, salary, a, sum)
	return
}

/*
func mapDbWrite()  {

	//创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	//map插入对应的key - value对,各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是",countryCapitalMap[country])
	}

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	}else {
		fmt.Println("American 的首都不存在")
	}
	//用户名：密码@tcp(地址:3306)/数据库名
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.2.132:3306)/test")
	if err != nil {
		fmt.Println(err)
	}

	//往数据插入数据
	for k,v := range countryCapitalMap {
	result, err := db.Exec("INSERT INTO countryCapital(country,capital)VALUES (?,?)", k, v)
		//result, err := db.Prepare("INSERT INTO countryCapital(country,capital)VALUES ($1,$2)")
		if err != nil {
			fmt.Println(result,err)
		}
		//result, err = result.Exec(k,v)
	}



}
*/
