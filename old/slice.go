package main;
import "fmt";


func main () {
	var s []int; // 슬라이스 선언
	s = make([]int, 5); // 초기화
	a := []int{1,2,3,4,5,6};
	b := make([]int, 5 , 10); // 길이가 5이고 용량이 10 용량은 전체 메모리 할당량, 길이는 인덱스 접근 가능 길이
	fmt.Println(s);
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(len(b));
	fmt.Println(cap(b));
	fmt.Println(b[0:3]);

	// 슬라이스는 배열과 다르게 참조형
	c := b;
	c[0] = 1;
	fmt.Println(c);

	// b 를 c에 복사
	f := make([]int, 5);
	copy(f, b);
	f[0] = 3;
	fmt.Println(f);
	fmt.Println(b);
}