package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	mydata := []byte("Hello GoLang")

	err := ioutil.WriteFile("temp.txt", mydata, 0777)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fmt.Println(string(data))

	f, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("New write file\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}
