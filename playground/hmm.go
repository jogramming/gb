package main

import (
	"fmt"
)

func main() {
	i := uint8(244)
	fmt.Println(i)
	i += 20
	fmt.Println(i)

	b := 244
	fmt.Println(b)
	b += 20
	fmt.Println(b)
	fmt.Println(uint8(b))

}
