package main;
import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
);
/*
	고루틴
	여러 함수를 동시해 실행하는 기능

	<동시성 - Concurrency>
	동시에 작업 하는 것 같이 보이는 것 (싱글코어)

	<병렬성 - Parallelism>
	실제로 동시에 여러 작업을 처리하는 것 (멀티코어)
		
*/
func hello(n int) {
	r := rand.Intn(100);
	time.Sleep(time.Duration(r));
	fmt.Println(n);
}
func main() {
	// 멀티코어 활용하기
	fmt.Println("start");
	runtime.GOMAXPROCS(runtime.NumCPU()); // CPU갯수 구한 다음 설정
	fmt.Println(runtime.GOMAXPROCS(0)); // 0 이면 설정값 출력
	defer fmt.Println("end");
	for i := 0; i < 100; i++ {
		go hello(i);
	}
}