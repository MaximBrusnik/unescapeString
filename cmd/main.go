package main

import (
	"fmt"
	"log"
	"unescapeString"
)

func main() {
	s, err := unescapeString.UnescapeString("a\\12b3")
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Println(s)
}
