package main

import (
	"errors"
	"strconv"
)

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func StringToInt64(s string, base int) (val int64, offset int, err error) {
	if len(s) == 0 {
		return -1, -1, errors.New("unable to parse empty string into int64")
	}

	scanIndex := 0
	if s[0] == '-' || s[0] == '+' {
		scanIndex++
	}
	for scanIndex < len(s) {
		if isDigit(s[scanIndex]) {
			scanIndex++
			continue
		}
		break
	}
	val, err = strconv.ParseInt(s[:scanIndex], base, 64)

	if err != nil {
		return -1, -1, err
	}
	return val, scanIndex, nil
}
