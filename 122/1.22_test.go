package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestMultiplyBigNumbers(t *testing.T) {
	a, _ := big.NewInt(0).SetString("1234567890", 10)
	b, _ := big.NewInt(0).SetString("9876543210", 10)
	expect, _ := big.NewInt(0).SetString("12193263111263526900", 10)
	got := MultiplyBigNumbers(a, b)
	assert.Equal(t, expect, got)
}

func TestDivideBigNumbers(t *testing.T) {
	a, _ := big.NewInt(0).SetString("12193263111263526900", 10)
	b, _ := big.NewInt(0).SetString("9876543210", 10)
	expect, _ := big.NewInt(0).SetString("1234567890", 10)
	got := DivideBigNumbers(a, b)
	assert.Equal(t, expect, got)
}

func TestAddBigNumbers(t *testing.T) {
	a, _ := big.NewInt(0).SetString("1234567890", 10)
	b, _ := big.NewInt(0).SetString("9876543210", 10)
	expect, _ := big.NewInt(0).SetString("11111111100", 10)
	got := AddBigNumbers(a, b)
	assert.Equal(t, expect, got)
}

func TestSubBigNumbers(t *testing.T) {
	a, _ := big.NewInt(0).SetString("9876543210", 10)
	b, _ := big.NewInt(0).SetString("1234567890", 10)
	expect, _ := big.NewInt(0).SetString("8641975320", 10)
	got := SubBigNumbers(a, b)
	assert.Equal(t, expect, got)
}
