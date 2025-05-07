package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b
}

func factorize(n int) []int {
	if n < 0 {
		n = -n
	}
	var facs []int
	for n%2 == 0 {
		facs = append(facs, 2)
		n /= 2
	}
	for p := 3; p*p <= n; p += 2 {
		for n%p == 0 {
			facs = append(facs, p)
			n /= p
		}
	}
	if n > 1 {
		facs = append(facs, n)
	}
	return facs
}

func printFactors(facs []int, primeOnly, exponents bool) {
	if primeOnly {
		seen := make(map[int]bool)
		var uniq []int
		for _, p := range facs {
			if !seen[p] {
				seen[p] = true
				uniq = append(uniq, p)
			}
		}
		facs = uniq
	}
	if exponents {
		counts := make(map[int]int)
		for _, p := range facs {
			counts[p]++
		}
		first := true
		for p, cnt := range counts {
			if !first {
				fmt.Print(" × ")
			}
			if cnt > 1 {
				fmt.Printf("%d^%d", p, cnt)
			} else {
				fmt.Printf("%d", p)
			}
			first = false
		}
		fmt.Println()
		return
	}
	for i, p := range facs {
		if i > 0 {
			fmt.Print(" × ")
		}
		fmt.Print(p)
	}
	fmt.Println()
}

type tokenType int

const (
	NUMBER tokenType = iota
	OP
	LPAREN
	RPAREN
)

type token struct {
	typ   tokenType
	value string
}

var prec = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

func tokenize(expr string) []string {
	var toks []string
	var buf strings.Builder
	ops := "+-*/^()%"
	for _, r := range expr {
		switch {
		case r == ' ':
			continue
		case strings.ContainsRune(ops, r):
			if buf.Len() > 0 {
				toks = append(toks, buf.String())
				buf.Reset()
			}
			toks = append(toks, string(r))
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		toks = append(toks, buf.String())
	}
	return toks
}

func lexRaw(strs []string) ([]token, error) {
	var out []token
	for _, s := range strs {
		switch s {
		case "(":
			out = append(out, token{LPAREN, s})
		case ")":
			out = append(out, token{RPAREN, s})
		case "+", "-", "*", "/", "^":
			out = append(out, token{OP, s})
		default:
			if _, err := strconv.ParseFloat(s, 64); err == nil {
				out = append(out, token{NUMBER, s})
			} else {
				return nil, fmt.Errorf("invalid token %q", s)
			}
		}
	}
	return out, nil
}

func evalExpr(tokens []token) (float64, error) {
	var vals []float64
	var ops []token

	applyOp := func() error {
		if len(vals) < 2 {
			return fmt.Errorf("not enough operands")
		}
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		b := vals[len(vals)-1]
		a := vals[len(vals)-2]
		vals = vals[:len(vals)-2]

		var res float64
		switch op.value {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				return fmt.Errorf("division by zero")
			}
			res = a / b
		case "^":
			res = math.Pow(a, b)
		}
		vals = append(vals, res)
		return nil
	}

	for _, tok := range tokens {
		switch tok.typ {
		case NUMBER:
			v, _ := strconv.ParseFloat(tok.value, 64)
			vals = append(vals, v)
		case OP:
			for len(ops) > 0 && ops[len(ops)-1].typ == OP {
				top := ops[len(ops)-1].value
				if prec[top] > prec[tok.value] ||
					(prec[top] == prec[tok.value] && tok.value != "^") {
					if err := applyOp(); err != nil {
						return 0, err
					}
					continue
				}
				break
			}
			ops = append(ops, tok)
		case LPAREN:
			ops = append(ops, tok)
		case RPAREN:
			for len(ops) > 0 && ops[len(ops)-1].typ != LPAREN {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("mismatched parens")
			}
			// pop "("
			ops = ops[:len(ops)-1]
		}
	}
	for len(ops) > 0 {
		if err := applyOp(); err != nil {
			return 0, err
		}
	}
	if len(vals) != 1 {
		return 0, fmt.Errorf("invalid expr")
	}
	return vals[0], nil
}

type cmdFunc func(nums []int, flags string)

var commands = map[string]cmdFunc{
	"gcm": func(nums []int, flags string) {
		res := nums[0]
		for _, v := range nums[1:] {
			res = gcd(res, v)
		}
		if strings.Contains(flags, "f") {
			facs := factorize(res)
			fmt.Printf("gcm = %d → ", res)
			printFactors(facs, strings.Contains(flags, "p"), strings.Contains(flags, "e"))
		} else {
			fmt.Println(res)
		}
	},
	"lcm": func(nums []int, flags string) {
		res := nums[0]
		for _, v := range nums[1:] {
			res = lcm(res, v)
		}
		if strings.Contains(flags, "f") {
			facs := factorize(res)
			fmt.Printf("lcm = %d → ", res)
			printFactors(facs, strings.Contains(flags, "p"), strings.Contains(flags, "e"))
		} else {
			fmt.Println(res)
		}
	},
	"fact": func(nums []int, flags string) {
		for _, v := range nums {
			facs := factorize(v)
			fmt.Printf("%d → ", v)
			printFactors(facs, strings.Contains(flags, "p"), strings.Contains(flags, "e"))
		}
	},
}

func main() {
	fmt.Println("Advanced Math REPL (flexible commands)")
	fmt.Println("Available:", keys(commands))
	fmt.Println("Type e.g. gcm%fe 12 18 30  or  fact%p 84 90")
	fmt.Println("Or enter any expression.  exit/quit/q to leave.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		lc := strings.ToLower(line)
		if lc == "exit" || lc == "quit" || lc == "q" {
			break
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmdToken := parts[0]
		if idx := strings.Index(cmdToken, "%"); idx != -1 {
			cmdToken, _ = cmdToken[:idx], cmdToken[idx+1:]
		}

		// strip flags off for lookup
		base, flags := cmdToken, ""
		if p := strings.Index(parts[0], "%"); p >= 0 {
			base = parts[0][:p]
			flags = parts[0][p+1:]
		}

		if fn, ok := commands[base]; ok {
			// parse ints
			var nums []int
			for _, s := range parts[1:] {
				if v, err := strconv.Atoi(s); err == nil {
					nums = append(nums, v)
				}
			}
			if len(nums) == 0 {
				fmt.Println("need at least one integer")
				continue
			}
			fn(nums, flags)
			continue
		}

		// fallback: arithmetic expression
		toks := tokenize(line)
		lexed, err := lexRaw(toks)
		if err != nil {
			fmt.Println("Lex error:", err)
			continue
		}
		res, err := evalExpr(lexed)
		if err != nil {
			fmt.Println("Eval error:", err)
			continue
		}
		if math.Floor(res) == res {
			fmt.Printf("%.0f\n", res)
		} else {
			fmt.Printf("%.10f\n", res)
		}
	}
}

// helper to list available commands
func keys(m map[string]cmdFunc) []string {
	var ks []string
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
