package main

import "math/big"

func MultiplyBigNumbers(a, b *big.Int) *big.Int {
	return a.Mul(a, b)
}

func DivideBigNumbers(a, b *big.Int) *big.Int {
	return a.Div(a, b)
}

func AddBigNumbers(a, b *big.Int) *big.Int {
	return a.Add(a, b)
}

func SubBigNumbers(a, b *big.Int) *big.Int {
	return a.Sub(a, b)
}
