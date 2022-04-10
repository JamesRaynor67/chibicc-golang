GOCMD=go
GOBUILD=$(GOCMD) build

chibicc-golang: main.go
	$(GOBUILD) -o chibicc-golang main.go strUtil.go tokenizer.go

test: chibicc-golang
	bash ./test.sh

clean:
	rm -f chibicc-golang tmp*

.PHONY: test clean
