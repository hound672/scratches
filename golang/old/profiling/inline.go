package main

//func main() {
//	// Declare variable of type int with a value of 10.
//	count := 10
//	sl := make([]int, 0, 10)
//
//	// Display the "value of" and "address of" count.
//	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
//
//	// Pass the "address of" count.
//	res := increment(&count)
//	sl = append(sl, *res)
//
//	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
//}
//
//////go:noinline
//func increment(inc *int) *int{
//
//	// Increment the "value of" count that the "pointer points to". (dereferencing)
//	*inc++
//	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
//	i := 10
//	return &i
//}
