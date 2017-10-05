package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Triztian/dep-no-root/subpkg"
)

func main() {
	a, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("A", a)
	fmt.Println("B", b)
	fmt.Println(subpkg.Add(int(a), int(b)))
}
