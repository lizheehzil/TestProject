package main

import (
	"fmt"
	"github.com/imroc/biu"
	"io/ioutil"
	"math"
)

func main() {
	file, err := ioutil.ReadFile("file/test.txt")
	if err != nil {

	}
	fmt.Printf("type is %T\n", file)
	println(len(file))
	println(biu.BytesToBinaryString(file))
	fmt.Printf("%x\n", file)
	println(file)
	println(string(file))
	println(math.Pow(2, 8))

}
