package main

import "fmt"

func zero(xPtr *int, str *string) {
	*xPtr = 0    //ค่าที่เปลี่ยนเมื่อใช้ Pointer
	*str = "Pug" //ค่าที่เปลี่ยนเมื่อใช้ Pointer
}

func main() {
	x := 5       //ค่าเริ่มต้น
	t := "satit" //ค่าเริ่มต้น
	zero(&x, &t)
	fmt.Println(x, t)
}
