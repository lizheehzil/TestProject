package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("year:%v\n", now.Year())
	fmt.Printf("month:%v\n", now.Month())
	fmt.Printf("day:%v\n", now.Day())
	fmt.Printf("hour:%v\n", now.Hour())
	fmt.Printf("UTC:%v\n", now.UTC())
	fmt.Printf("Unix timestamp:%v\n", now.Unix())
	fmt.Printf("Local:%v\n", now.Local())
	fmt.Printf("format_time:%v\n", now.Format("2006--01--02 15.04.05"))

}
