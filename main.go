package main

func main() {
	p := new(peg)
	p.Buffer = `required optional`
	p.Init()
	p.Parse()
}
