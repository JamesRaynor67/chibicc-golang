package main

import (
	"errors"
	"strconv"
)

func StringToInt64(s string, base int) (val int64, offset int, err error) {
	if len(s) == 0 {
		return -1, -1, errors.New("unable to parse empty string into int64")
	}

	offset = 1
	phaseValue := int64(-1)
	err = nil
	hasValidResult := false

	if s[0] == '-' || s[0] == '+' {
		offset = 2
	}
	for offset <= len(s) {
		num, err := strconv.ParseInt(s[:offset], base, 64)
		offset++

		if err != nil {
			offset--
			break
		}

		phaseValue = num
		hasValidResult = true
	}

	if hasValidResult {
		return phaseValue, offset - 1, nil
	}
	return -1, -1, err
}
