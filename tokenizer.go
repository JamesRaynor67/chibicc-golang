package main

import (
	"bytes"
	"log"
)

type TokenKind int32

const (
	TK_PUNCT TokenKind = iota
	TK_NUM   TokenKind = iota
	TK_EOF   TokenKind = iota
)

var TOKEN_TYPE TokenKind

type Token struct {
	kind      TokenKind
	next      *Token
	val       int64
	wholeLine []byte // Where token embedded in
	offset    int    // The 0-based start index of the token
	length    int    // The length of the token
}

func (t Token) ByteSlice() []byte {
	return t.wholeLine[t.offset : t.offset+t.length]
}

func (t Token) Equals(s []byte) bool {
	return bytes.Equal(t.ByteSlice(), s)
}

func (t Token) Skips(s []byte) *Token {
	if !t.Equals(s) {
		log.Panicf("Expects token string %s but got %s", t.ByteSlice(), s)
	}
	return t.next
}

func (t Token) GetInt64Val() int64 {
	if t.kind != TK_NUM {
		log.Panicf("Expects token of type TK_NUM but got %d", t.kind)
	}
	return t.val
}

func NewToken(kind TokenKind, wholeLine []byte, offset, length int) *Token {
	p := new(Token)
	p.kind = kind
	p.wholeLine = wholeLine
	p.offset = offset
	p.length = length
	return p
}

func Tokenize(s []byte) *Token {
	dummyTokenHead := new(Token)
	var cur *Token = dummyTokenHead

	scanIndex := 0
	for scanIndex < len(s) {
		// Skip spaces
		if s[scanIndex] == ' ' || s[scanIndex] == '\t' {
			scanIndex++
			continue
		}

		// Get int token
		// Assume only positive int without sign in input
		if IsDigit(s[scanIndex]) {
			val, offset, err := StringToInt64(s[scanIndex:], 10)
			if err != nil {
				log.Panic(err)
			}

			cur.next = NewToken(TK_NUM, s, scanIndex, offset)
			cur.next.val = val
			cur = cur.next
			scanIndex += offset
			continue
		}

		// Get Punctuator
		if s[scanIndex] == '-' || s[scanIndex] == '+' {
			cur.next = NewToken(TK_PUNCT, s, scanIndex, 1)
			cur = cur.next
			scanIndex++
			continue
		}

		log.Panicf("Invalid token: %s", s[scanIndex:])
	}

	cur.next = NewToken(TK_EOF, []byte(""), 0, 0)
	return dummyTokenHead.next
}
