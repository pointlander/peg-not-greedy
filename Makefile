.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

all: greedy.go greed

greed: greedy.go
	go build -o greed

