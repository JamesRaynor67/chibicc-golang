package main

import "testing"

func TestStringToInt64(t *testing.T) {
	s := "123"
	val, offset, err := StringToInt64(s, 10)
	if val != 123 || offset != 3 || err != nil {
		t.Error("Error to parse ", s, ". Expect 123, 3 but got ", val, offset)
	}

	s = "1+12"
	val, offset, err = StringToInt64(s, 10)
	if val != 1 || offset != 1 || err != nil {
		t.Error("Error to parse ", s, ". Expect 1, 1 but got ", val, offset)
	}
	val, offset, err = StringToInt64(s[2:], 10)
	if val != 12 || offset != 2 || err != nil {
		t.Error("Error to parse ", s, ". Expect 12, 2 but got ", val, offset)
	}

	s = "-1-2"
	val, offset, err = StringToInt64(s, 10)
	if val != -1 || offset != 2 || err != nil {
		t.Error("Error to parse ", s, ". Expect -1, 2 but got ", val, offset)
	}
	val, offset, err = StringToInt64(s[3:], 10)
	if val != 2 || offset != 1 || err != nil {
		t.Error("Error to parse ", s, ". Expect 2, 1 but got ", val, offset)
	}

	s = " -1"
	val, offset, err = StringToInt64(s, 10)
	if val != -1 || offset != -1 || err == nil {
		t.Error("Expecting to get error due to leading space but no error is returned: ", val, offset, err)
	}
}
