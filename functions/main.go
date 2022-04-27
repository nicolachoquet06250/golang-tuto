package functions

import "fmt"

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

/*func rectProps(length, width float64) (float64, float64) {
	area = length * width
	perimeter = (length + width) * 2
	return area, perimeter
}*/

func Main() {
	fmt.Println("calculateBill(200, 2) =", calculateBill(200, 2))

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}
