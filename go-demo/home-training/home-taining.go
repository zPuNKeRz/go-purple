package main

import "fmt"

func main() {
	fmt.Println(getProductInfo("Mouse", 25.0, 10))

}

func getProductInfo(name string, price float64, count int) string {
	result := fmt.Sprintf("Товар: %s, Цена: %.2f, Количество: %d", name, price, count)

	return result
}
