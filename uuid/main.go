package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func wString(str uuid.UUID) {
	fmt.Print(str)
}

func main() {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4:%s\n", u1)

	u2, err := uuid.FromString("37db4899-f479-4fa2-943b-d90db5ea01e9")
	if err != nil {
		fmt.Printf("Something gone wrong %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s\n", u2)
	fmt.Printf("u2 type is %T\n", u2)

	fmt.Print(u2, "\n")

	wString(u2)

}
