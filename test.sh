#!/bin/bash
assert() {
  expected="$1"
  input="$2"

  ./chibicc-golang "$input" > tmp.s || exit
  gcc -static -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$input => $expected expected, but got $actual"
    exit 1
  fi
}

assert 0 0
assert 42 42
assert 10 '5+11-6'
assert 11 '4 + 10 - 3   + 123456  -   123456'

echo OK
