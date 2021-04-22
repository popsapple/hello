package main;
import (
	"fmt"
);

/*
	덕타이핑 :: 실제 타입이 아닌, 구현된 메서드로만 타입을 판단하는 방식..
	예를들어 사람이 꽥 하고 "울든" 오리가 안녕 하고 "울든"... 울기만 하면 같다고 보겠다.
*/
func (i Duck) Howl() {
	fmt.Println("꽥");
}
func (i Person) Howl() {
	fmt.Println("사람이 꽥");
}
type Duck struct {}
type Person struct {}
type Sound interface {
	Howl();
}
func AtNight(i Sound) { // 여기가 핵심... 메서드 기준으로 판단한다는 뜻
	i.Howl();
}
func main() {
	/*
		인터페이스의 주체와 데이터 구조의 주체가 다르도록 해야 하는 부분이 헷갈린다...
	*/
	var duck Duck;
	var ps Person;
	i := Sound(duck);
	j := Sound(ps);
	AtNight(i);
	AtNight(j);
	// 인터페이스 체크
	if v, ok := interface{}(ps).(Sound); ok {
		fmt.Println(v, ok);
	}
	// 빈 인터페이스 - any타입이나 마찬가지임
	var ifce1 interface{} = 1;
	var ifce2 interface{} = "Hello~";
	var num1 int = ifce1.(int); // type assert (타입 단언) 개발자가 직접 자료형을 지정
	var num2 string = ifce2.(string); // type assert (타입 단언) 개발자가 직접 자료형을 지정
}