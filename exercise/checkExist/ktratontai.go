package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("testfile.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	buf1 := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			buf1 = append(buf1, buf...)

		}
	}

	make := "5"

	fmt.Println(strings.Contains(string(buf1), make))

}
