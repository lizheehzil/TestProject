package main

import (
	"encoding/json"
	"fmt"
)
import "TestProject/json/person"

func main() {

	person1 := person.Person{
		Name: "name1",
		Age:  12,
		Des:  "test information",
	}

	data, err := json.Marshal(person1)
	if err == nil {
		fmt.Println(string(data))

		person2 := person.Person{}
		json.Unmarshal(data, &person2)
		fmt.Println(person2.Name, person2.Age, person2.Des)
	}

	println()

}
