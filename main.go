package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	lineParseRe *regexp.Regexp
)

func init() {
	rand.Seed(time.Now().UnixNano())
	lineParseRe = regexp.MustCompile(`^(\S+)\s*((?:\S*)|(?:\".+\"))\s*(?:\s#.*)?$`)
}

func main() {

	flag.Parse()
	filepath := flag.Arg(0)

	if filepath == "" {
		fmt.Println("no program file specified")
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	lineNumber := 0
	stack := newStack()

	defer func() {
		perr := recover()
		if perr, ok := perr.(error); ok {
			fmt.Printf("error on line %d: %s\n", lineNumber, perr)
			os.Exit(1)
		}
	}()

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		lineNumber++
		if len(line) == 0 || isComment(line) {
			continue
		}

		cmd, arg, err := splitCmdArg(line)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("cmd: %s\targ: %s\n", cmd, arg)

		switch cmd {
		case "push":
			stack.push(parseInt(arg))
		case "pop":
			stack.pop()
		case "dup":
			a := stack.pop()
			stack.push(a)
			stack.push(a)

		case "add":
			a, b := stack.pop2()
			stack.push(b + a)
		case "sub":
			a, b := stack.pop2()
			stack.push(b - a)
		case "mul":
			a, b := stack.pop2()
			stack.push(b * a)
		case "div":
			a, b := stack.pop2()
			stack.push(b / a)
		case "mod":
			a, b := stack.pop2()
			stack.push(b % a)
		case "pow":
			a, b := stack.pop2()
			stack.push(int64(math.Pow(float64(b), float64(a))))

		case "eq":
			a, b := stack.pop2()
			stack.pushBool(b == a)
		case "neq":
			a, b := stack.pop2()
			stack.pushBool(b != a)

		case "gt":
			a, b := stack.pop2()
			stack.pushBool(b > a)

		case "gte":
			a, b := stack.pop2()
			stack.pushBool(b >= a)

		case "lt":
			a, b := stack.pop2()
			stack.pushBool(b < a)

		case "lte":
			a, b := stack.pop2()
			stack.pushBool(b <= a)

		case "not":
			a := stack.pop()
			if a == 0 {
				stack.push(rand.Int63())
			} else {
				stack.push(0)
			}

		case "print":
			if arg == "C" {
				fmt.Print(string(stack.pop()))
			} else {
				fmt.Print(stack.pop())
			}
		case "println":
			if arg == "C" {
				fmt.Println(string(stack.pop()))
			} else {
				fmt.Println(stack.pop())
			}
		case "read":
			fmt.Print(strings.Trim(arg, `"`))
			var in string
			fmt.Scanln(&in)
			stack.push(parseInt(in))

		default:
			panic(fmt.Errorf("not a command: %s", cmd))
		}
	}

	fmt.Println("\n stack at the end")
	fmt.Println(stack)
}

func splitCmdArg(line string) (cmd, arg string, err error) {
	matches := lineParseRe.FindStringSubmatch(line)
	if matches == nil || len(matches) < 2 {
		return "", "", errors.New("syntax error")
	}
	return matches[1], matches[2], nil
}

func isComment(line string) bool {
	return len(line) > 0 && line[0] == '#'
}

func parseInt(arg string) int64 {
	n, err := strconv.Atoi(arg)
	if err != nil {
		n = int(arg[0])
	}
	return int64(n)
}

type stack []int64

func newStack() *stack {
	s := stack(make([]int64, 0, 64))
	return &s
}

func (s *stack) length() int64 {
	return int64(len(*s))
}

func (s *stack) push(i int64) {
	// fmt.Printf("\tpushed %d\n", i)
	*s = append(*s, i)
}

func (s *stack) pushBool(b bool) {
	if b {
		s.push(1)
		return
	}
	s.push(0)
}

func (s *stack) pop() int64 {
	if s.length() < 1 {
		panic(fmt.Errorf("stack underflow"))
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	// fmt.Printf("\tpopped %d\n", top)
	return top
}

func (s *stack) pop2() (int64, int64) {
	return s.pop(), s.pop()
}
