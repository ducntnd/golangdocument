package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64  {
	return math.Pi * math.Pow(c.radius,2)
}

//type Student struct {
//	name string
//	grades []int
//	age int
//}
//
//func (s *Student) setAge(age int)  {
//	s.age = age
//}

func test(x,y int) (int,int) {
	fmt.Println("test")
	return x+y,x-y
}


func main()  {
	//scanner := bufio.NewScanner(os.Stdin)
	//input,_ := strconv.ParseInt(scanner.Text(),10,64)
	//fmt.Println("You typed: ",input)
	//
	//number := 26000000000.0
	//var name = fmt.Sprintf("%T %v",number,number)
	//name="duc"
	//fmt.Printf("%9.2f", name)


	//num3 := 7
	//num1 = num1+num3
	//num3 = num1 - num3
	//num1 -= num3


	//var num1 int= 4
	//var num2 float64 = 3.2
	// answer:= float64(num1) / num2
	// boole:=(4>3.2)||false
	//fmt.Println(answer,boole)


	//var arr [5]int64
	//for i := 0; i < len(arr) ; i++{
	//	fmt.Println("Nhap so")
	//	scanner.Scan()
	//	arr[i], _ = strconv.ParseInt(scanner.Text(),10,64)
	//	fmt.Println(arr[i])
	//}
	//for i,element := range arr{
	//	fmt.Println(i,element)
	//}
	//fmt.Println(arr[1:3])


	//var mp map[int]string = map[int]string{
	//	1:"apple",
	//	2:"banana",
	//	3:"orange",
	//}
	//delete(mp,2)
	//val,ok := mp[2]
	//fmt.Println(mp,ok,val)
	//
	//returnFunc("duc")()


	//map vs slice: mutable


	//x := 3
	//y := &x
	//*y = 5
	//fmt.Println(x)


	//s1 := Student{"Duc",[]int{9,10,11,12},20}
	//s1.setAge(9)
	//fmt.Println(s1.age)


	c1 := circle{4.5}
	r1 := rect{5,7}
	shapes := []shape{&c1,&r1}
	for _,shape := range shapes{
		fmt.Println(shape.area())
	}

	fmt.Println(fibonaci(6))
}

func fibonaci(n int) int {
	if n<=2 {
		return 1
	}
	return fibonaci(n-1) + fibonaci(n-2)
}

func returnFunc(x string) func()  {
	return func() {
		fmt.Println(x)
	}
}
