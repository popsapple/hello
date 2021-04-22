package main;
import (
	"fmt"
	"os"
);

func ReadHello() {
	text, err := os.Open("hello.txt");
	defer text.Close();

	if (err != nil) {
		fmt.Println(err);
		return; // 종료직전 defer 호출
	}

	buf := make([]byte, 100);
	if _, err = text.Read(buf); err != nil {
		fmt.Println(err);
		return; // 종료직전 defer 호출 
	}
	fmt.Println(string(buf));
}

func main() {
	ReadHello();
}