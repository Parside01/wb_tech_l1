package main

import "fmt"

func DivideTemperature(temp []float32) map[int][]float32 {
	res := make(map[int][]float32)
	for _, t := range temp {
		key := int(t) / 10 * 10
		res[key] = append(res[key], t)
	}
	return res
}

func main() {
	fmt.Println(DivideTemperature([]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}))
}
