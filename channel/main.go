package main

import "fmt"

func main() {

	//var cha chan map[string]string
	cha := make(chan map[string]string, 10)
	map1 := make(map[string]string)
	map1["city1"] = "北京"
	map1["city2"] = "成都"
	map1["city3"] = "广州"

	map2 := make(map[string]string)
	map2["hero1"] = "宋江"
	map2["hero2"] = "武松"

	cha <- map1
	cha <- map2

	outmap1 := <-cha
	outmap2 := <-cha

	println(outmap1["city1"])
	println(outmap1["city2"])
	println(outmap1["city3"])
	println()
	println(outmap2["hero1"])
	println(outmap2["hero2"])

	println("==============================")
	for s, s2 := range outmap1 {
		fmt.Printf("%v=%v", s, s2)
		println()
	}
	println()
	for s, s2 := range outmap2 {
		fmt.Printf("%v=%v", s, s2)
		println()
	}
}
