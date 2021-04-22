package main;
import (
	"fmt"
	"math/rand"
	"time"
);
/*
	채널
	고루틴끼리 데이터를 주고받고 실행 흐름을 제어하는 기능 모든 타입을 채널로 사용할 수 있습니다.
*/
func hello(n int) {
	r := rand.Intn(100);
	time.Sleep(time.Duration(r));
	fmt.Println(n);
}
/*
	채널
	고루틴끼리 데이터를 주고받고 실행 흐름을 제어하는 기능 모든 타입을 채널로 사용할 수 있습니다.
*/
func sum(a int, b int, c chan int) {
	// 합을 채널로 보낸다.
	c <- a + b
}
func main() {
	c := make(chan int); // int형 채널 생성 반드시 make로 메모리 할당해줘야함
	go sum(1,2,c);
	// 이걸로 인해 채널에서 값이 들어올 때까지 대기 했다가 그 값을 넣는다.
	n := <- c 
	fmt.Println(n);
}