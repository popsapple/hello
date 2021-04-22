package main;
import (
	"fmt"
	"time"
);
func main() {
	isdone := make(chan bool); // int형 채널 생성 반드시 make로 메모리 할당해줘야함
	count := 3
	go func() {
		for i := 0; i < count; i++ {
			isdone <- true
			fmt.Println("고루틴 : ", i);
			time.Sleep(1 * time.Second);
		}
	}();
	for i := 0; i < count; i++ {
		<- isdone
		fmt.Println("메인함수 : ", i);
	}
}