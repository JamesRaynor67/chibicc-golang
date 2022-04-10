package main

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	tokenHead := Tokenize([]byte("1 + 234 - 5"))

	cur := tokenHead
	if cur.val != 1 || cur.kind != TK_NUM || string(cur.ByteSlice()) != "1" {
		t.Error("Error! Got", cur.val, int(cur.kind), cur.ByteSlice())
	}

	cur = cur.next
	if cur.kind != TK_PUNCT || string(cur.ByteSlice()) != "+" {
		t.Error("Error! Got", int(cur.kind), cur.ByteSlice())
	}

	cur = cur.next
	if cur.val != 234 || cur.kind != TK_NUM || string(cur.ByteSlice()) != "234" {
		t.Error("Error! Got", cur.val, int(cur.kind), cur.ByteSlice())
	}

	cur = cur.next
	if cur.kind != TK_PUNCT || string(cur.ByteSlice()) != "-" {
		t.Error("Error! Got", int(cur.kind), cur.ByteSlice())
	}

	cur = cur.next
	if cur.val != 5 || cur.kind != TK_NUM || string(cur.ByteSlice()) != "5" {
		t.Error("Error! Got", cur.val, int(cur.kind), cur.ByteSlice())
	}

	cur = cur.next
	if cur.kind != TK_EOF || string(cur.ByteSlice()) != "" {
		t.Error("Error! Got", int(cur.kind), cur.ByteSlice())
	}
}
