package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if (len(os.Args) > 1) && (len(os.Args[1]) > 0) {
		stringWordBuff := strings.Split(os.Args[1], " ")
		fmt.Println(len(stringWordBuff))
	} else {
		fmt.Println(0)
	}
}
