package main;
import "fmt";
import _ "io/ioutil";


func main () {
	var a [5]int;
	a[2] = 7;
	c := [2]int{1,2};
	d := [...]int{1,2,3,4};
	// 그냥 배열이 바로 복사됨
	f := d;
	f[2] = 7;
	fmt.Println(a);
	fmt.Println(c);
	fmt.Println(d);
	fmt.Println(f);
}