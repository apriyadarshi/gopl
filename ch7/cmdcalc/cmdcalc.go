package main

import (
	"fmt"
	"gopl/ch7/expression"
	"strconv"
)

func main() {
	for {
		var strExpr string
		fmt.Println("enter expression: ")
		fmt.Scanln(&strExpr)
		fmt.Println(strExpr)
		expr, errP := expression.Parse(strExpr)
		if errP != nil {
			fmt.Printf("parse: couldn't parse expression: %s\n", errP)
			continue
		}

		vars := make(map[expression.Var]bool)
		errC := expr.Check(vars)
		if errC != nil {
			fmt.Printf("check: %s", errC)
			continue
		}

		env := make(map[expression.Var]float64)
		for k := range vars {
			var s string
			for {
				fmt.Printf("enter value for variable %s:\n", k)
				fmt.Scanln(&s)
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					fmt.Printf("entered value %s can't be parsed as float\n", s)
					continue
				} else {
					env[expression.Var(k)] = f
					break
				}
			}
		}
		fmt.Printf("\nAns: %f\n", expr.Eval(env))
		fmt.Println("---------------------------------------\n")
	}
}
