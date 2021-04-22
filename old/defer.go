package main;
import "fmt";
import _ "strconv";

func print() {
	fmt.Println("Hello");
}
func world() {
	fmt.Println("world");
}
func main() {
	defer world();
	print();
	print();
	print();
}