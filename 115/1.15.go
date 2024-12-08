package main

import (
	"math/rand"
	"strings"
)

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createHugeString(length int) string {
	return randSeq(length)
}

var justString string

// Проблема в том, что мы делаем срез, который внутри указывает на строку длинной 1 << 10,
// то есть в justString мы используем только 100 символов, а внутри храним указатель на
// строку с большим количеством символов => если строка будет больше, то мы храним много ненужной памяти.

// Вот как можно исправить.
// Теперь мы клонируем нужную нам часть строки,
// а после выхода их функции указатель внутри v будет удален.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
