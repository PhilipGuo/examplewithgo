// helloworld
package main

import (
	"fmt"
	"time"
)

const pi float32 = 3.14

func main() {
	// 1 hello world
	fmt.Println("Hello World!")

	// 2 values
	fmt.Println("go" + "lang")
	fmt.Println("7/3", 7/3)
	fmt.Println(true && false)
	fmt.Println(!false)

	// 3 const
	const a int = 5
	fmt.Println(a / 2)
	fmt.Println(pi / 2)
	const c bool = true
	fmt.Println(c)

	// 4 for loop
	for i := 0; i < 3; {
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println("loop")
		break
	}

	var d int = 9
	for d > 0 {
		fmt.Print(d)
		d--
	}

	//5 if - else
	if 1 == 0 {
		fmt.Print("1==0")
	} else if 1 == 2 {
		fmt.Print("1 == 2")
	} else {
		fmt.Print("1 !=0 and 1 != 2")
	}

	// 6 switch
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Print("this is weekend")
	case time.Monday:
	case time.Tuesday:
		fmt.Print("tuesday")
	default:
		fmt.Print("default")
	}

	switch {
	case 1 > 0:
		fmt.Print("1 > 0")
	default:
		fmt.Print("de")
	}

	whatiam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Print("bool")
		case string:
			fmt.Print("string")
		default:
			fmt.Print("type dont know", t)
		}
	}

	whatiam("aa")
	whatiam(true)

	// 7 array
	var arr [3]int
	var brr [3]int
	fmt.Println(brr)
	fmt.Println(arr)
	arr[1] = 2
	fmt.Println(arr)
	fmt.Println(brr)

	crr := [2]int{1, 2}
	fmt.Println(crr)

	drr := crr
	fmt.Println(drr)
	fmt.Print(len(drr))

	var drrr [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(drrr[i][j])
		}
	}

	// 8 slice
	var s []int // 声明
	fmt.Print(s)
	s = make([]int, 3) // 初始化
	fmt.Print(s)

	s[1] = 1 // 赋值
	s[2] = 2
	fmt.Print(s)

	s = append(s, 4) // append
	fmt.Print(s)

	s1 := s[1:3] // s[1] s[2]
	fmt.Print(s1)

	s2 := s[:4] // [0:4)
	fmt.Print(s2)

	s3 := s[:6] // 如果是 s[:7] 则会异常，因为s初始化为长度3，在append 4 时，对应的数组的长度会增加一倍为6， 所以s[:7]会out of range的异常
	fmt.Println(s)
	fmt.Println(s3)

	s4 := []string{"i", "j"}
	fmt.Print(s4)

	twod := make([][]int, 3)
	fmt.Print(twod)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		tmp := make([]int, innerlen)
		twod[i] = tmp
		for j := 0; j < innerlen; j++ {
			twod[i][j] = i + j
		}
	}

	fmt.Println(twod)

	// 9 map
	m := make(map[int]int)
	fmt.Println(m)

	m[1] = 1
	fmt.Print(m)

	delete(m, 1)
	fmt.Print(m)

	v, ok := m[1]
	fmt.Println(v, ok) // 0, false

	// 10 range
	kvs := map[int]string{1: "a", 2: "b"}
	for k, v := range kvs { // 并不按顺序
		fmt.Println(k, v)
	}

	for i, c := range "golang" {
		fmt.Print(string(i) + " :" + string(c))
	}

	// 11 function
	sum := Sum(1, 4)
	fmt.Println(sum)

	// 12 多返回值
	mk, mk2 := multi()
	fmt.Println(mk, mk2)

	// 13 变参函数
	fmt.Println(Sum2(1, 2, 3))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum2(nums...)) // 不能用 Sum2(nums)

	// 14 闭包
	seq := initSeq()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	newnext := initSeq()
	fmt.Println(newnext())

	// 15 递归调用
	fmt.Println(fact(7))

	// 16 指针
	val := 1
	fmt.Println("initial:", val)
	zeroval(val)
	fmt.Println("zeroval:", val)
	zeroptr(&val)
	fmt.Println("zeroptr:", val)

	// 17 结构体(定义，初始化，赋值，传递)
	p := person{"bob", 20}
	fmt.Println(p) // {bob 20}
	p1 := person{name: "bob", age: 20}
	fmt.Println(p1)
	p2 := &person{"philip", 28}
	fmt.Println(p2)  // &{philip 28}
	fmt.Println(*p2) //{philip 28}
	p2.age = 29
	fmt.Println(*p2)
	p3 := person{}
	fmt.Println(p3)
	p3.age = 30
	fmt.Println(p3)

	// 18 方法
	fmt.Println(p.getname())
	p.setname("guo") // 不会改变
	fmt.Println(p.getname())
	p.setnameptr("guo")
	fmt.Println(p.getname())

	// 19 接口

}
func (p *person) setnameptr(name string) {
	p.name = name
}
func (p person) setname(name string) {
	p.name = name
}
func (p person) getname() string {
	return p.name
}

type person struct {
	name string
	age  int
}

func zeroval(val int) {
	val = 0
}

func zeroptr(val *int) {
	*val = 0
	fmt.Println(&val)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func initSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func Sum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func multi() (int, bool) {
	return 1, true
}
func Sum(a, b int) int {
	return a + b
}
