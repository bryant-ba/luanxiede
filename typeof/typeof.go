package main

import (
	"fmt"
	"reflect"
)

func main() {
	var booknum float32 = 6
	var isbook bool = true
	bookauthor := "www"
	bookdetail := make(map[string]string)
	bookdetail["Go"] = "www"
	fmt.Println(reflect.ValueOf(booknum))
	fmt.Println(reflect.ValueOf(isbook))
	fmt.Println(reflect.ValueOf(bookauthor))
	fmt.Println(reflect.ValueOf(bookdetail))
}
