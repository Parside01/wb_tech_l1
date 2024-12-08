package main

import "fmt"

// Все применяют указатели, а у нас будет арифметика)
func Swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

// Ну это просто подмена
func SwapClassic(a, b *int) {
	a, b = b, a
}

func main() {
	a, b := 10, 5
	Swap(&a, &b)
	fmt.Println(a, b)
}
