package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println(os.Args[1:])

	f, err := os.Create("/home/timotewb/development/go/1.19/api/app01.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
