package main

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}

func (e *HTTPError) Error() string {
	return "some error"
}

type CustomError struct {
	*HTTPError
	ownField int
}

func (e *CustomError) As(c interface{}) bool {
	fmt.Printf("CALLED: %T; v: %v;\n", c, c)
	v, ok := c.(**HTTPError)
	*v = e.HTTPError
	fmt.Printf("res: %v; %T\n", (*v).Code, v)
	return ok
}

func main() {
	fmt.Printf("Start\n")
	
	err1 := &HTTPError{Code: 222}
	err2 := &CustomError{HTTPError: &HTTPError{Code: 111}}
	var httpErr *HTTPError
	var httpErr2 *HTTPError
	//var customError *CustomError

	if errors.As(err1, &httpErr) {
		fmt.Printf("1. As: %v\n", errors.As(err1, &httpErr))
		fmt.Printf("res: %v\n", httpErr.Code)
		fmt.Printf("------------\n")
	}
	if errors.As(err2, &httpErr2) {
		fmt.Printf("2. As\n")
		fmt.Printf("res: %v\n", httpErr2.Code)
		fmt.Printf("------------\n")
	}

	//fmt.Printf("As: %v\n", errors.As(err2, &httpErr))
	//fmt.Printf("As: %v\n", errors.As(err2, &customError))

	swit(err1)
	swit(err2)
}

func swit(err error) {
	fmt.Printf("err: %v\n", err)
	
	switch h := err.(type) {
	case *HTTPError:
		fmt.Printf("httpErrror\n")
	case *CustomError:
		fmt.Printf("custom: %v\n", h.Code)
	}
}