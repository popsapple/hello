package main;
import (
	"fmt"
);

type Ractangle struct {
	width int
	height int
}

// 생성자 패턴
func NewRactangle(width, height int) *Ractangle {
	return &Ractangle{width, height}; // 인스턴스 생성 후 포인터 리턴
}

// 메서드 개념을 넣어 클래스처럼 쓸 수 있도록 하자
// funt (리시버) 함수명(매개변수) 형태
func (ract *Ractangle) area() int {
	// 여기서 ract를 리시버라고 함
	return ract.width * ract.height; // 인스턴스 생성 후 포인터 리턴
}


func main() {
	var testRact Ractangle;
	var testRact1 *Ractangle;
	testRact1 = new(Ractangle);
	testRact2 := new(Ractangle);
	
	testRact3 := Ractangle{10, 20};
	testRact4 := Ractangle{width: 30, height: 25};

	fmt.Println(testRact);
	fmt.Println(testRact1); // 포인터래서 찍으면 앞에 &나옴
	fmt.Println(testRact1.width);
	fmt.Println(testRact2.width);
	fmt.Println(testRact3.width);
	fmt.Println(testRact4.width);
	testRact1.width = 10;
	testRact1.height = 10;
	fmt.Println(testRact1.area());
}