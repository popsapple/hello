package main;
import "fmt";


func main () {
	var s map[string]int; // 맵 선언
	s = make(map[string]int); // 메모리 할당..
	a := map[string]int{"hello": 10, "item": 10}; // 초기화
	s["fruits"] = 10;
	fmt.Println(s);
	fmt.Println(a);
	// k, v := range array 이 형태임을 기억하자
	for key, value := range a {
		fmt.Println(key, value);
	}
	delete(a, "hello");
	fmt.Println(a);
}