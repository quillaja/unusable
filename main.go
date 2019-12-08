package main

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var lineParseRe *regexp.Regexp

func init() {
	lineParseRe = regexp.MustCompile(`^(\S+)\s*((?:\S*)|(?:\".+\"))\s*(?:\s#.*)?$`)
}

func main() {
	const prog = `
push	-5
push	4	# comment
add

push A

# a comment
#another comment
println	I
print	N
read	"Is this 9? "
`

	fmt.Printf("program:\n%s\n\n", prog)

	s := bufio.NewScanner(strings.NewReader(prog))
	lineNumber := 0
	stack := newStack()

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		lineNumber++
		if len(line) == 0 || isComment(line) {
			continue
		}

		cmd, arg, err := splitCmdArg(line)
		if err != nil {
			fmt.Printf("error on line %d: %s\n", lineNumber, err)
			break
		}

		fmt.Printf("cmd: %s\targ: %s\n", cmd, arg)

		switch cmd {
		case "push":
			n, err := strconv.Atoi(arg)
			if err != nil {
				n = int(arg[0])
			}
			stack.push(int64(n))
		case "pop":
			stack.pop()
		case "dup":
			t := stack.pop()
			stack.push(t)
			stack.push(t)

		case "add":
			a, b := stack.pop(), stack.pop()
			stack.push(a + b)
		case "sub":
		case "mul":
		case "div":
		case "mod":
		case "pow":

		case "eq":
		case "neq":
		case "gt":
		case "gte":
		case "lt":
		case "lte":
		case "not":

		case "print":
		case "println":
		case "read":

		default:
			fmt.Printf("error on line %d: not a command: %s\n", lineNumber, cmd)
			break
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
	// parts := strings.SplitN(line, " ", 2)
	// cmd = parts[0]
	// if len(parts) > 1 {
	// 	rest = parts[1]
	// }
	// return
}

func isComment(line string) bool {
	return len(line) > 0 && line[0] == '#'
}

type stack []int64

func newStack() *stack {
	s := stack(make([]int64, 0, 64))
	return &s
}

func (s *stack) push(i int64) {
	fmt.Printf("\tpushed %d\n", i)
	*s = append(*s, i)
}

func (s *stack) pop() int64 {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	fmt.Printf("\tpopped %d\n", top)
	return top
}
