package main

import (
	"bufio"
	"os"
)

func main() {
	filepath := "file/filedemo1/demo1.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		println(err.Error())
		return
	}

	str := "hello file\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}
