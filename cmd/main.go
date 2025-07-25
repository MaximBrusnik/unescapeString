package main

import (
	"fmt"
	"log"
	"unescapestring"
)

func main() {
	s, err := unescapestring.UnescapeString("a\\12b3")
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Println(s)
}
