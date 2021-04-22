package main;
import "fmt";
import "strconv";

func hello(s string) {
	fmt.Println(s);
}
func add(s string , b string) string {
	return s + " " + b;
}
func total(num ...int) int {
	count := 0;
	for _, v := range num {
		fmt.Println("?????" + strconv.Itoa(v));
		count += v;
	}
	return count;
}
func main () {
	var test = make([]int, 10, 10);
	test[0] = 22;
	hello("good");
	fmt.Println(add("good","morning"));
	fmt.Println(total(1,2,3,4,5));
	fmt.Println(test);
	fmt.Println(total(test...));
}