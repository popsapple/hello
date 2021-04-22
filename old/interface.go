package main;
import (
	"fmt"
);
type Prase string;
type Numbering int;
func (i Prase) Print() { 
	// 리시버를 통해서 연결 근데 여기서 리시버는 사용자 정의 타입이어야 한다.. ㅠㅠ 
	fmt.Println(i);
}
func (i Numbering) Print() { 
	// 리시버를 통해서 연결 근데 여기서 리시버는 사용자 정의 타입이어야 한다.. ㅠㅠ 
	fmt.Println(i);
}
type Hello interface {
	Print()
}
func main() {
	/*
		인터페이스의 주체와 데이터 구조의 주체가 다르도록 해야 하는 부분이 헷갈린다...
	*/
	var i Prase = "hello";
	var j Numbering = 15;
	var h Hello; // 인터페이스 설정
	h2 := Hello(j); // 인터페이스 설정 하면서 값 세팅
	h = i; // 데이터타입 설정? 및 값 세팅
	h.Print();
	h2.Print();
}