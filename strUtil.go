package main

import (
	"errors"
	"strconv"
)

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func StringToInt64(s []byte, base int) (val int64, offset int, err error) {
	if len(s) == 0 {
		return -1, -1, errors.New("unable to parse empty string into int64")
	}

	scanIndex := 0
	if s[0] == '-' || s[0] == '+' {
		scanIndex++
	}
	for scanIndex < len(s) {
		if IsDigit(s[scanIndex]) {
			scanIndex++
			continue
		}
		break
	}
	val, err = strconv.ParseInt(string(s[:scanIndex]), base, 64)

	if err != nil {
		return -1, -1, err
	}
	return val, scanIndex, nil
}
