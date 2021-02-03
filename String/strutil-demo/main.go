package main

import "github.com/ozgio/strutil"

func main() {
	box, err := strutil.DrawBox("gostrutil", 20, strutil.Center)
	if err != nil {
		print(err)
	}
	println(box)
}
