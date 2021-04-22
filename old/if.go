package main;
import (
	"fmt"
	"io/ioutil"
);

// 한 블럭 내의 iota 는 0에서부터 증가함
const (
	// 첫번째를 _로 패스할 수 없다
	Sunday = iota
	_
	Tuesday
	Thursday
)

func main () {
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b);
	}
	Loop:
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if j == 2 {
					break Loop;
				}
				fmt.Println(i, j);
			}
		}
}