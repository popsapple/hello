package main;
import "fmt";
import "strconv";

func main () {
	a, b := 5,3;
	testfun := func (a, b int) int {
		return a + b;
	}
	fmt.Println("AAAA :: " + strconv.Itoa(testfun(a,b)));
}