package main

import "fmt"

func main() {
	var t1 int
	var t2 interface{}
	t3 := new(int)
	t4 := make([]int, 3, 4)

	fmt.Printf("t1 type is %T\n", t1)
	fmt.Printf("t2 type is %T\n", t2)
	fmt.Printf("t3 type is %T\n", t3)
	fmt.Printf("t4 type is %T\n", t4)
	fmt.Print(t4, "\n")
	fmt.Println(len(t4))
}
