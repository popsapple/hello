package main;
import (
	"fmt"
	"time"
);
func main() {

	// 나중에 자로형에 표기할 때 <-chan int 꺼내기 전용 채널
	// 나중에 자로형에 표기할 때 chan<- int 받기 전용 채널
	//done := make(chan bool, 2); // 두번째가 버퍼 크기
	//count := 4;
	// ??? 이게 돌리다보니 버퍼의 최대 용량이 2개인것 뿐이지 
	// 예제처럼 두개 다 차고 나서야 빼는 행위를 한다는게 아님...
	// 스택오버플로우 형님들 답변으로는 대기열 크기 이기 때문에 처리속도가 빠르면 스택이 안 쌓이고 그냥 나간다고 함...
	channel1 := make(chan string);
	channel2 := make(chan string);

	// 그리고 내 고루틴이 왜 자꾸 동작 안하는지 확인해봤는데
	// 스택오버플로우 형님들이 프로세스가 바로 끝나버리므로 고루틴이 기다릴 시간이 없어서 그냥 종료되는거라 함
	// 그래서 time.Sleep 넣어줌 ㅡㅡ
	go func() {
		channel1 <- "hello";
	}()
    time.Sleep(2 * time.Second)
	go func() {
		for {
			select {
				case v:= <-channel1:
					fmt.Println("channel1", v);
					return
				case s:= <-channel2:
					fmt.Println("channel2", s);
					return
				default:
					fmt.Println("default");
					return
			}
		}
	}()
    time.Sleep(5 * time.Second)
}