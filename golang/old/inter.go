package main

import (
	"fmt"
)

type CustomBuffer struct{
	fmt.Stringer
}

func (cb CustomBuffer) String() string {
	return "custom buffer"
}

type NamedBuffer CustomBuffer

//func (nb NamedBuffer) String() string {
//	return "named buffer"
//}

func cast(i interface{}) {
	if obj, ok := i.(fmt.Stringer); ok {
		fmt.Printf("casted: %T: %T\n", i, obj)
		//fmt.Printf("Casted: %s\n", obj.String())
	} else {
		fmt.Printf("err cast: %T\n", i)
	}
}

func main() {
	fmt.Println(CustomBuffer{}.String())
	fmt.Printf("T: %T\n", NamedBuffer{})

	cast(CustomBuffer{})
	cast(NamedBuffer{})
}
