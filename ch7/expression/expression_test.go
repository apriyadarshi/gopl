package expression

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		//{"sqt(A / pi)", Env{"B": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3) + pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x,3) + pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		{"min(x,y,z)", Env{"x": 3, "y": 1, "z": 0}, "0"},
	}

	var prevExpr string
	for _, test := range tests {
		//Print expr only when it changes
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr) //Parse function populates the expression object's fields
		if err != nil {
			t.Error(err)
			continue
		}

		//Error in check
		vars := make(map[Var]bool)
		if err := expr.Check(vars); err != nil {
			fmt.Printf("error found in check: %s", err)
			continue
		}
		for v := range vars {
			if _, ok := test.env[v]; !ok {
				fmt.Printf("Var %s is undefined", v)
				continue
			}
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %s = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}

func TestPrint(t *testing.T) {
	env := make(map[Var]float64)
	env["A"] = float64(87616)
	env["pi"] = math.Pi
	expD := Expr(call{fn: "sqrt", args: []Expr{Expr(binary{op: '/', x: Expr(Var("A")), y: Expr(Var("pi"))})}})

	expP, errP := Parse(fmt.Sprintf("%s", expD))
	if errP != nil {
		fmt.Printf("Error while parsing %s", errP)
	}

	if expD.Eval(env) != expP.Eval(env) {
		t.Errorf("not reparsable: ")
	}

}
