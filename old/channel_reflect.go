package main;
import (
	"fmt"
	"reflect"
);
type Data struct {
	a,b int
}
type Person struct {
	name string `tag1:"이름"` // 세미콜론 옆에 띄어쓰지말자...
	age int `tag1:"나이"`
}
func main() {
	// reflect 실행 시점에 인터페이스나 구조체등의 타입 정보를 얻거나 결정하는 기능.
	var num int = 1;
	fmt.Println(reflect.TypeOf(num));

	var s string = "Hello, world!";
	fmt.Println(reflect.TypeOf(s));
	
	var p *Person = new(Person);
	p.name = "가람";
	p.age = 12;
	fmt.Println(reflect.TypeOf(p));

	name, ok := reflect.TypeOf(*p).FieldByName("name");
	fmt.Println(name);
	fmt.Println(ok, name.Tag.Get("tag1"));

	age, ok := reflect.TypeOf(*p).FieldByName("age");
	fmt.Println(age.Tag.Get("tag1"));
}