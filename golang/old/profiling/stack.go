package main

type Obj struct {
	Val int
}

func f3() {
	obj := f2()
	_ = obj
}

func f2() *Obj {
	obj := f1()
	_ = *obj
	//return obj
	return nil
}

////go:noinline
func f1() *Obj {
	obj := Obj{Val: 1}
	return &obj
}

//func main() {
//}
