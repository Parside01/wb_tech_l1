package main

import (
	"strconv"
)

// В целом там в тесте генерирются рандомные значения, так что main писать смысла немного.
func ChangeBit(num int64, pos, bit int) int64 {
	binLen := len(strconv.FormatInt(num, 2))
	mask := int64(1 << (binLen - pos))
	res := num

	if bit == 1 {
		res |= mask
	} else {
		res &= ^mask
	}
	return res
}
