package main
import "fmt"
// 한 블럭 내의 iota 는 0에서부터 증가함
const (
	// 첫번째를 _로 패스할 수 없다
	Sunday = iota
	_
	Tuesday
	Thursday
)

func main () {
	fmt.Println(Thursday);
}