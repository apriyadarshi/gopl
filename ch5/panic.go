package ch5

import "fmt"

func cheekyReturn() (s string) {
	defer func() string {
		if p := recover(); p != nil {
			s = fmt.Sprintf("%v", p)
		}
	}()

	panic("Hi")
}
