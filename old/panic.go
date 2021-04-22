package main;
import (
	"fmt"
);
func f() {
	defer func() {
		s := recover();
		fmt.Println(s); // 패닉시의 내용을 recover 가 담아줌... 이거 없으면 로그에 Error !! 안뜸
	}()
	panic("Error !!"); // panic 함수로 에러 메세지 설정, 패닉 발생 후 멈춰야 하지만 recover을 함
}
func main() {
	f();
	fmt.Println("Hello, world!"); // 패닉이 발생했지만 계속 실핼됭
}