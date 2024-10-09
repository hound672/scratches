package main

type Obj struct {
	index int
}

func sumRange(objects []Obj) int {
	ret := 0
	for _, v := range objects {
		ret += v.index
	}
	return ret
}

func sumLoop(objects []Obj) int {
	ret := 0
	for i := 0; i < len(objects); i++ {
		ret += objects[i].index
	}
	return ret
}


