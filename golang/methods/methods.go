package main

type Inter interface {
	Method()
}

// receiver is pointer
type Pointer struct {}
func (p *Pointer) Method() {}

// receiver is value
type Value struct {}
func (v Value) Method() {}

func f(i Inter) {i.Method()}

func main() {
	pP := &Pointer{}
	f(pP) // correct

	pV := Pointer{}
	f(pV) // wrong

	vP := &Value{}
	f(vP) // correct

	vV := Value{}
	f(vV) // correct
}
