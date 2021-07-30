package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	f, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)

}
