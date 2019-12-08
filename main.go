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
	// lineParseRe = regexp.MustCompile(`^(\S+)\s*((?:\S*)|(?:\".+\"))\s*(?:\s#.*)?$`)
	lineParseRe = regexp.MustCompile(`((?:#.*)|(?:".+")|(?:\S+))`)
}

func main() {

	flag.Parse()
	filepath := flag.Arg(0)

	if filepath == "" {
		fmt.Println("no program file specified")
		os.Exit(1)
	}

	program, err := readLines(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state := &state{
		executing:   true,
		proceedures: make(map[string]proc),
		stack:       newStack(),
		program:     program,
	}

	defer func() {
		perr := recover()
		if perr, ok := perr.(error); ok {
			fmt.Printf("error on line %d: %s\n", state.lineNumber, perr)
			os.Exit(1)
		}
	}()

	start := time.Now()
	execute(state, state.program)
	took := time.Since(start)

	fmt.Println("\nstate at the end")
	fmt.Println(state)
	fmt.Println("took:", took)
}

func execute(state *state, lines []string) {
	for _, line := range lines {
		state.lineNumber++
		// fmt.Printf("%d: %s\n", state.lineNumber, line)
		if len(line) == 0 || isComment(line) {
			continue
		}
		cmd, args, err := splitCmdArgs(line)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("cmd: %s\targs: %v\n", cmd, args)
		interpret(state, cmd, args)
	}
}

func interpret(state *state, cmd string, args []string) {
	stack := state.stack

	switch cmd {
	case "def":
		if len(args) == 0 {
			panic(fmt.Errorf("def requires a proceedure name"))
		}
		name := args[0]
		if p, exists := state.proceedures[name]; exists {
			panic(fmt.Errorf("proceedure %s already defined on line %d", name, p.start))
		}

		state.proceedures[name] = proc{start: state.lineNumber}
		state.executing = false
		return

	case "end":
		if len(args) == 0 {
			panic(fmt.Errorf("def requires a proceedure name"))
		}
		name := args[0]
		if _, exists := state.proceedures[name]; !exists {
			panic(fmt.Errorf("no proceedure named %s defined", name))
		}

		p := state.proceedures[name]
		p.end = state.lineNumber
		state.proceedures[name] = p
		state.executing = true
		return
	}

	if !state.executing {
		return
	}

	switch cmd {
	case "push":
		if len(args) < 1 {
			panic(fmt.Errorf("push requires 1 argument"))
		}
		stack.push(parseInt(args[0]))
	case "pop":
		stack.pop()
	case "dup":
		a := stack.pop()
		stack.push(a)
		stack.push(a)
	case "len":
		stack.push(stack.length())
	case "rot":
		a, b := stack.pop2()
		stack.rotate(b, a)

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
		if len(args) > 0 && args[0] == "C" {
			fmt.Print(string(stack.pop()))
		} else {
			fmt.Print(stack.pop())
		}
	case "println":
		if len(args) > 0 && args[0] == "C" {
			fmt.Println(string(stack.pop()))
		} else {
			fmt.Println(stack.pop())
		}
	case "read":
		if len(args) > 0 {
			fmt.Print(strings.Trim(args[0], `"`))
		}
		var in string
		fmt.Scanln(&in)
		stack.push(parseInt(in))

	case "cond":
		if len(args) == 0 {
			panic(fmt.Errorf("cond requires a statement"))
		}
		a := stack.pop()
		if a != 0 {
			interpret(state, args[0], args[1:])
		}

	case "call":
		if len(args) == 0 {
			panic(fmt.Errorf("call requires a proceedure name"))
		}
		name := args[0]
		p, exists := state.proceedures[name]
		if !exists {
			panic(fmt.Errorf("no proceedure named %s defined", name))
		}

		prevLineNumber := state.lineNumber
		state.lineNumber = p.start
		execute(state, state.program[p.start:p.end-1])
		state.lineNumber = prevLineNumber

	case "exit":
		os.Exit(0)

	default:
		panic(fmt.Errorf("not a command: %s", cmd))
	}
}

func readLines(filepath string) (program []string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		program = append(program, line)
	}
	err = s.Err()
	return
}

func splitCmdArgs(line string) (cmd string, args []string, err error) {
	matches := lineParseRe.FindAllString(line, -1)
	length := len(matches)
	if matches == nil || length < 1 {
		return "", nil, errors.New("syntax error")
	}
	if isComment(matches[length-1]) {
		matches = matches[:length-1]
	}
	return matches[0], matches[1:], nil
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

func (s *stack) rotate(depth, times int64) {
	switch {
	case depth < 0:
		panic(fmt.Errorf("rotation depth must be non-negative"))
	case depth == 0:
		return
	}

	index := s.length() - depth
	if index < 0 {
		panic(fmt.Errorf("depth %d is larger than stack size %d", depth, s.length()))
	}

	forward := true
	if times < 0 {
		times = -times
		forward = false
	}

	for ; times > 0; times-- {
		if forward {
			temp := (*s)[s.length()-1]
			copy((*s)[index+1:], (*s)[index:index+depth-1])
			(*s)[index] = temp
		} else {
			temp := (*s)[index]
			copy((*s)[index:index+depth-1], (*s)[index+1:])
			(*s)[s.length()-1] = temp
		}
	}
}

type state struct {
	executing   bool
	lineNumber  int
	stack       *stack
	proceedures map[string]proc
	program     []string
}

func (s *state) String() string {
	return fmt.Sprintf("stack: %v\nprocs: %v", *s.stack, s.proceedures)
}

type proc struct {
	start, end int
}
