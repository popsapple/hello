package main

import "fmt"
import "unsafe"

func main() {
	var num1 int8 = 1;
	var num2 int16 = 1;
	var num3 int32 = 1;
	var hello string = "hello";

	fmt.Println(unsafe.Sizeof(num1));
	fmt.Println(unsafe.Sizeof(num2));
	fmt.Println(unsafe.Sizeof(num3));
	fmt.Println(len(hello));
	fmt.Printf("%c\n", hello[1]);
	fmt.Printf("%c\n", hello[2]);
}