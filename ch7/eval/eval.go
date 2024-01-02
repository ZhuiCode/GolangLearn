package eval

import (
	"fmt"
	"math"
)

type Expr interface {
	Eval(env Env) float64
}

/*变量*/
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

/*浮点数常量*/
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)

}

/*有一个操作数的操作符表达式*/
type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator:%q", u.op))
}

/*有两个操作数的操作符表达式*/
type binary struct {
	op   rune
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
	}
	panic(fmt.Sprintf("unsupported binary operator:%q", b.op))
}

/*函数表达式*/
type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call:%q", c.fn))
}

type Env map[Var]float64
