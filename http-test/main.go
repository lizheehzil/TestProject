package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s := fmt.Sprintln("你好世界! -- Time:%s", time.Now().String())
		_, _ = fmt.Fprintln(writer, s)
		log.Printf("%v\n", s)

	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
