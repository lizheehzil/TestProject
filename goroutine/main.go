package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func sayHello() {

	for i := 0; i < 15; i++ {
		fmt.Println("hello!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func main() {
	go sayHello()
	println(runtime.GOMAXPROCS(0))

	time.Sleep(time.Second * 5)
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println("golang!" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	//println(runtime.NumCPU())
}
