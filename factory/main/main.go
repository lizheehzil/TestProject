package main

import (
	"TestProject/factory/modle"
	"fmt"
)

func sayHello() {
	for i := 0; i < 300; i++ {

		println("hello")
	}
}

func sayNo() {
	println("No")
}

func main() {
	//stu := modle.Student{"name1", 11}
	//fmt.Printf("student name is %v,age is %v",stu.Name,stu.Age)

	stu := modle.NewStudent("name1", 11)

	fmt.Println(stu.Age, stu.Name)
	fmt.Println(stu)

}
