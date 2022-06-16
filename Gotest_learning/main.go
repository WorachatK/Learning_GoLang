package main

import (
	"fmt"
)

////////////////////// main ///////////////////////////

func main() {
	struct_func()
}

////////////////////// if_else ///////////////////////////

func if_else() {
	var number int
	fmt.Print("กรุณาใส่คะแนน : ")
	fmt.Scanf("%d", &number)
	fmt.Println("คะแนน = ", number)

	if number == 1 {
		fmt.Println("เปิดบัญชีใหม่")
	} else if number == 2 {
		fmt.Println("ฝากขถอน")
	} else {
		fmt.Println("กรุณาใส่เลข 1 หรือ 2")
	}
}

////////////////////// switch_case ///////////////////////////

func switch_case() {
	var number int
	fmt.Print("กรุณาใส่คะแนน : ")
	fmt.Scanf("%d", &number)
	fmt.Println("คะแนน = ", number)

	switch number {
	case 1:
		fmt.Println("เปิดบัญชีใหม่")
	case 2:
		fmt.Println("เปิดบัญชีใหม่")
	default:
		fmt.Println("กรุณาใส่เลข 1 หรือ 2")
	}
}

////////////////////// Array ///////////////////////////

func array_func() {
	//var numbers [3]int = [3]int{100,200,300}
	numbers := [...]int{100, 200, 300}
	size := len(numbers) //size of array
	fmt.Println(numbers)
	fmt.Println("size of array = ", size)
}

////////////////////// Slices ///////////////////////////

func slices_func() {
	numbers := []int{100, 200, 300} //เพิ่มและลดขนาดได้
	numbers = append(numbers, 400)
	fmt.Println(numbers)
	fmt.Println(numbers[1:3]) //เริ่มจาก index 1-3
}

////////////////////// Maps ///////////////////////////

func maps_func() {
	country := map[string]string{"TH": "Thailand", "JP": "Japan"}
	country["EN"] = "English"

	fmt.Println(country)
	value, check := country["JP"]

	if check {
		fmt.Println(value)
	} else {
		fmt.Println("No data")
	}
}

////////////////////// For loop ///////////////////////////

func for_loop() {
	for count := 1; count <= 10; count++ {
		if count == 5 {
			// break //หยุดโปรแกรม
			continue //ข้ามรอบนั้นไปรอบต่อไป
		}
		fmt.Println(count)
	}
	fmt.Println("End Program")
}

////////////////////// For Rang ///////////////////////////

func loop_rang() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, value := range numbers {
		fmt.Println("value = ", value)
	}
}

////////////////////// Struct ///////////////////////////

type Product struct {
	name string
	price float64
	category string
	discount int
}

func struct_func(){
	Product1 := Product{name: "ปากกา",price: 10.5,category: "เครื่องเขียน",discount: 10}
	fmt.Println(Product1)
}
