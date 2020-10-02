package main

import "fmt"

type A struct {
	Anum1 int
	Anum2 int
}

type B struct {
	A
	Bnum1 int
	Bnum2 int
}

func main() {
	//var t1 int
	//var t2 interface{}
	//t3 := new(int)
	//t4 := make([]int, 3, 4)
	//
	//fmt.Printf("t1 type is %T\n", t1)
	//fmt.Printf("t2 type is %T\n", t2)
	//fmt.Printf("t3 type is %T\n", t3)
	//fmt.Printf("t4 type is %T\n", t4)
	//fmt.Print(t4, "\n")
	//fmt.Println(len(t4))
	b := new(B)
	b.Anum1 = 1
	b.Anum2 = 2
	b.Bnum1 = 3
	b.Bnum2 = 4

	c := B{*new(A),
		1,
		3}

	//fmt.Printf("type is %T",b)
	fmt.Println(c.Anum2)
	fmt.Println(c.Bnum2)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", new(B))


}
