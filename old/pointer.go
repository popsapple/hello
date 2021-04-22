package main;
import (
	"fmt"
);
func main() {
	var testPtr *int = new(int); // new 함수로 공간 할당
	*testPtr = 1; // 역참조 해서 직접 접근
	fmt.Println(*testPtr); // 역참조 해서 직접 접근

	var testItem = 2;
	var testPtr2 *int = &testItem; // 레퍼런스(참조)로 주소를 넣어준다.
	fmt.Println(testPtr2); // 주소
	fmt.Println(*testPtr2); // 역참조 해서 직접 접근

	/*
		C와 차이점
		int *testPtr -> c에서는 변수명에다가 붙여서 쓴다.
		C는 포인터 연산 등을 통해 동적접근이 가능하지만 go는 그런거 없다.
	*/

	// 포인터를 쓰지 않으면 참조가 아닌 값복사로 들어가기 때문에 n이 바뀌지 않음.
	var n = 1;
	hello := func(n *int) { 
		*n = 2;
	}	
	hello(&n) // 주소 전달.
	fmt.Println(n);
}