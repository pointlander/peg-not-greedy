package main

func main() {
	p := new(peg)
	p.Buffer = `requiredoptional`
	p.Init()
	p.Parse()
}
