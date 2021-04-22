package main;
import (
	"fmt"
);

type Person struct {
	name string
	age string
}
func (p *Person) greeting() {
	fmt.Println("hello~");
}
type Student struct {
	p Person
	school string
	grade int
}
type Students struct {
	Person
	school string
	grade int
}

func main() {
	var VSP Student;
	var VSPS Students;
	VSP.p.greeting(); // has-a 관계
	VSPS.greeting(); // is-a 관계
}