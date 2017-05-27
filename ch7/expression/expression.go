package expression

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	//returns value of expression in env
	Eval(env Env) float64
	//Checks for static errors in expression tree.
	//It adds successfully checked vars to map
	Check(vars map[Var]bool) error
	//Returns string representation of the tree
	String() string
}

//Maps variable names to their actual values
type Env map[Var]float64

//Identifies a variable
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

//map is used after checking the whole expression
//its checked with env whether a particular var is defiend or not
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil //No error possible
}

func (v Var) String() string {
	return string(v)
}

//A literal is a numeric constant
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil //No error possible
}

func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

//A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune //+ or -
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env)
	case '-':
		return 0.0 - u.x.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported unary operator : %q", u.op))
	}
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary operator %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%c(%s)", u.op, u.x)
}

type binary struct {
	op   rune //+,-,*,/
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported binary operator : %q", b.op))
	}
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary operator %q", b.op)
	}
	errX := b.x.Check(vars)
	errY := b.y.Check(vars)
	if errX == nil && errY == nil {
		return nil
	} else {
		return fmt.Errorf("errors: X :- %s Y:- %s", errX, errY) //Combine the two errors
	}
}

func (b binary) String() string {
	return fmt.Sprintf(" ( (%s) %c (%s) ) ", b.x, b.op, b.y)
}

//a function call expression
type call struct {
	fn   string //pow,sin,sqrt
	args []Expr //arguments to the function 1-2
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "min":
		floats := make([]float64, len(c.args), len(c.args))
		for i, exp := range c.args {
			floats[i] = exp.Eval(env)
		}
		result, err := min(floats)
		if err != nil {
			panic(fmt.Sprintf("min: ", err))
		}
		return result
	default:
		panic(fmt.Sprintf("unsupported function call : %q", c.fn))
	}
}

func (c call) Check(vars map[Var]bool) error {
	if !(c.fn == "sin" || c.fn == "pow" || c.fn == "sqrt" || c.fn == "min") {
		return fmt.Errorf("unexpected function call: %s", c.fn)
	}
	if (c.fn == "sin" || c.fn == "sqrt") && len(c.args) != 1 {
		return fmt.Errorf("unexpected no of arguments for call %s: needed 1, got: %d", c.fn, len(c.args))
	}
	if (c.fn == "pow") && len(c.args) != 2 {
		return fmt.Errorf("unexpected no of arguments for call %s: needed 2, got: %d", c.fn, len(c.args))
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	if len(c.args) == 1 {
		return fmt.Sprintf(" ( %s(%s) ) ", c.fn, c.args[0])
	} else {
		return fmt.Sprintf(" ( %s(%s, %s) ) ", c.fn, c.args[0], c.args[1])
	}
}

func min(floats []float64) (float64, error) {
	if floats == nil || len(floats) == 0 {
		return 0, fmt.Errorf("min: empty float slice")
	}
	min := floats[0]
	for _, v := range floats {
		if v < min {
			min = v
		}
	}
	return min, nil
}
