package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// clearScreen clears the terminal screen.
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// TokenType represents the type of a token.
type TokenType int

const (
	INTEGER TokenType = iota
	OPERATOR
	LEFT_PAREN
	RIGHT_PAREN
	FUNCTION
)

// Token represents a token in the lexer.
type Token struct {
	Type  TokenType
	Value string
}

// Lexer struct represents the lexer.
type Lexer struct{}

// lex performs lexical analysis on the given tokens.
func (l *Lexer) lex(tokens []string) ([]Token, error) {
	var lexedTokens []Token
	operatorType := map[string]TokenType{
		"+": OPERATOR,
		"-": OPERATOR,
		"*": OPERATOR,
		"/": OPERATOR,
		"^": OPERATOR,
		"%": OPERATOR,
	}

	for _, token := range tokens {
		switch token {
		case "(":
			lexedTokens = append(lexedTokens, Token{Type: LEFT_PAREN, Value: token})
		case ")":
			lexedTokens = append(lexedTokens, Token{Type: RIGHT_PAREN, Value: token})
		default:
			if opType, ok := operatorType[token]; ok {
				lexedTokens = append(lexedTokens, Token{Type: opType, Value: token})
			} else if _, err := strconv.ParseFloat(token, 64); err == nil {
				lexedTokens = append(lexedTokens, Token{Type: INTEGER, Value: token})
			} else {
				return nil, fmt.Errorf("Invalid token: %s", token)
			}
		}
	}

	return lexedTokens, nil
}

// evaluate performs evaluation of operators and numbers.
func evaluate(operators *Stack, numbers *Stack) error {
	op := operators.Pop().(Token)

	value2 := numbers.Pop().(float64)
	value1 := numbers.Pop().(float64)

	switch op.Value {
	case "+":
		numbers.Push(value1 + value2)
	case "-":
		numbers.Push(value1 - value2)
	case "*":
		numbers.Push(value1 * value2)
	case "/":
		if value2 == 0 {
			return fmt.Errorf("Division by zero")
		}
		numbers.Push(value1 / value2)
	case "^":
		numbers.Push(math.Pow(value1, value2))
	case "%":
		numbers.Push(math.Mod(value1, value2))
	}

	return nil
}

// parse parses the tokens and performs the calculation.
func parse(tokens []Token) (float64, error) {
	var numbers Stack
	var operators Stack
	precedence := map[string]int{
		"+": 1, "-": 1, "%": 1, "^": 3, "*": 2, "/": 2,
	}

	for _, token := range tokens {
		switch token.Type {
		case INTEGER:
			numbers.Push(stringToF64(token.Value))
		case OPERATOR:
			for !operators.IsEmpty() && operators.Top().(Token).Type == OPERATOR &&
				precedence[operators.Top().(Token).Value] >= precedence[token.Value] {
				if err := evaluate(&operators, &numbers); err != nil {
					return 0, err
				}
			}
			operators.Push(token)
		case LEFT_PAREN:
			operators.Push(token)
		case RIGHT_PAREN:
			for operators.Top().(Token).Type != LEFT_PAREN {
				if err := evaluate(&operators, &numbers); err != nil {
					return 0, err
				}
			}
			operators.Pop()
		}
	}

	for !operators.IsEmpty() {
		if err := evaluate(&operators, &numbers); err != nil {
			return 0, err
		}
	}

	return numbers.Top().(float64), nil
}

// Stack represents a simple stack data structure.
type Stack []interface{}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// Pop removes and returns the element from the top of the stack.
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

// Top returns the element from the top of the stack without removing it.
func (s *Stack) Top() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// stringToF64 parses a string to a float64.
func stringToF64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Failed to parse integer: %s\n", s)
		return 0.0
	}
	return f
}

// tokenize breaks the expression into tokens.
func tokenize(expression string) []string {
	var tokens []string
	ops := "+-*/()^%"
	var builder strings.Builder

	for _, c := range expression {
		if c == ' ' {
			continue
		} else if strings.ContainsRune(ops, c) {
			if builder.Len() > 0 {
				tokens = append(tokens, builder.String())
				builder.Reset()
			}

			tokens = append(tokens, string(c))
		} else {
			builder.WriteRune(c)
		}
	}

	if builder.Len() > 0 {
		tokens = append(tokens, builder.String())
	}

	return tokens
}

func main() {
	clearScreen()

	fmt.Println("Math Expression Evaluator")
	fmt.Println("Enter quit or exit or q to close the program")

	var expression string
	lexer := Lexer{}

	for {
		fmt.Print(">> ")
		fmt.Scanln(&expression)

		checkInput := strings.ToLower(expression)

		if checkInput == "exit" || checkInput == "quit" || checkInput == "q" {
			break
		}

		tokens := tokenize(expression)

		lexedTokens, err := lexer.lex(tokens)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := parse(lexedTokens)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if result != math.Floor(result) {
			fmt.Printf("%.10f\n", result)
		} else {
			fmt.Println(result)
		}
	}
}
