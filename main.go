package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		stringWordBuff := strings.Split(os.Args[1], " ")
		fmt.Println(len(stringWordBuff))
	}
}
