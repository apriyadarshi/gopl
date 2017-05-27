package ch4

//Implementing stack using slices. Right now only distributed functions.
//Later can be associated with one struct and interfaces to create a proper data structure
// Assumption : Int stack
// TBD- Generic stack

//pushes an elem on stack and returns appended stack. Caller should set the stack to its original value.
func push(stack []int, v int) []int {
	return append(stack, v)
}

//Gets top element of stack
func top(stack []int) int {
	return stack[len(stack)-1]
}

//Pops/removes and returns the top element of stack
/*
func pop(stack []int) int {
	top := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return top
}
Important: The above won't set stack as the new value as in function a copy of reference to the
slice named stack is passed*/

//Returns the stack with one less item. Caller should set the value to its original value
func pop(stack []int) ([]int, int) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

//Remove int at index i; order = yes if order needs to be preserved
func remove(stack []int, i int, order bool) []int {
	if order {
		copy(stack[i:], stack[i+1:])
	} else {
		stack[i] = stack[len(stack)-1]
	}
	return stack[:len(stack)-1]
}
