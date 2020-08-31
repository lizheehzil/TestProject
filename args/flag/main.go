package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		user string
		pwd  string
		host string
		port int
	)

	flag.StringVar(&user, "u", "lizhe", "This is a name")
	flag.StringVar(&pwd, "pwd", "lizhe", "This is a pwd")
	flag.StringVar(&host, "h", "127.0.0.1", "This is a host")
	flag.IntVar(&port, "p", 80, "This is a port")

	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,prot=%v", user, pwd, host, port)
}
