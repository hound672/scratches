package main

import (
	"fmt"
	"log"
	"strings"
)

type Specification interface {
	Get() string
	Result() string
}

// EqSpec

type EqSpec struct {
	Val int
	Specification
}

func (e *EqSpec) Get() string {
	return fmt.Sprintf("[EQ: %d]", e.Val)
}

func (e *EqSpec) Result() string {
	res := e.Get()
	return fmt.Sprintf("RESULT: %s", res)
}

func NewEqSpec(val int) Specification {
	return &EqSpec{Val: val}
}

// AndSpec

type AndSpec struct {
	specs []Specification
	Specification
}

func (and *AndSpec) Get() string {
	vals := make([]string, 0, len(and.specs))
	for _, spec := range and.specs {
		vals = append(vals, spec.Get())
	}

	result := strings.Join(vals, " AND ")

	return fmt.Sprintf("(%s)", result)
}

func (and *AndSpec) Result() string {
	result := and.Get()
	return fmt.Sprintf("RESULT: %s", result)
}

func NewAndSpec(specs ...Specification) Specification {
	return &AndSpec{
		specs: specs,
	}
}

//

func main() {
	log.Printf("Start spec")

	// spec := NewEqSpec(10)
	spec := NewAndSpec(NewEqSpec(10))

	result := spec.Result()

	log.Printf("result: %v", result)
}
