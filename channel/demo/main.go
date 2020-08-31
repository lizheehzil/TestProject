package main

import "time"

var num = 80000000
var cpunum = 2

func writeData(inChan chan int) {
	for i := 0; i < num; i++ {
		inChan <- i
		//fmt.Printf("写入一个数据是%v", i)
		//println()
		//time.Sleep(time.Second)
	}

	close(inChan)
}

func readData(inChan chan int, exitChan chan bool) {

	for true {

		_, ok := <-inChan
		//time.Sleep(time.Second)
		if !ok {
			break
		}

		//fmt.Printf("读到一个数据是%v", i)
		//println()
	}
	exitChan <- true
}

func main() {

	inChan := make(chan int, num)
	exitChan := make(chan bool, 4)

	go writeData(inChan)
	for i := 0; i < cpunum; i++ {

		go readData(inChan, exitChan)
	}
	start := time.Now().Unix()
	for i := 0; i < cpunum; i++ {
		<-exitChan
	}
	close(exitChan)
	println(time.Now().Unix() - start)
}
