package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestChangeBit(t *testing.T) {
	casesCount := gofakeit.IntRange(1, 1)
	tt := make([]struct {
		PosBitChange int
		ChangeTo     int
		Num          int64
		ExpectNum    int64
	}, casesCount)
	for i := range tt {
		tt[i].Num = int64(gofakeit.IntRange(10, 100000))
		binStr := strconv.FormatInt(tt[i].Num, 2)
		tt[i].PosBitChange = gofakeit.IntRange(1, len(binStr))
		tt[i].ChangeTo = gofakeit.IntRange(0, 1)

		changeBin := strings.Join([]string{binStr[:tt[i].PosBitChange-1], strconv.Itoa(tt[i].ChangeTo), binStr[tt[i].PosBitChange:]}, "")
		tt[i].ExpectNum, _ = strconv.ParseInt(changeBin, 2, 64)
	}

	for _, tc := range tt {
		assert.Equal(t, tc.ExpectNum, ChangeBit(tc.Num, tc.PosBitChange, tc.ChangeTo))
	}
}
