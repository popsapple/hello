package main;
import (
	"fmt"
	"time"
	"runtime"
	"sync"
);
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU());
	var data = []int{};
	var mutex = new(sync.Mutex);

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock(); // 다른 고루틴 접근 금지
			data = append(data, 1);
			mutex.Unlock(); // 다른 고루틴 접근금지 풀기

			runtime.Gosched(); // 다른 고루틴테게 CPU 양보
		}
	}();
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock();
			data = append(data, 1);
			mutex.Unlock();

			runtime.Gosched();
		}
	}();

	// mutex 를 걸지 않으면 두 고루틴이 같은 data를 바라보기 때문에 값이 이상하게 증가함
	time.Sleep(time.Second * 2);
	fmt.Println(len(data));
}