# unusable
the **UNU**sable **S**t**A**ck **B**ased **L**anguag**E**

A toy stack based language written in Go.

# stack
64 bit signed integers
'infinite' depth

# scripts
single file
read line by line. each line is a statement, comment, or blank.

a statement is:
`KEYWORD [arg]`

# keywords

| keyword | args | description |
| --- | --- | --- |
| # | | a comment line. Inline comments are allowed. |
| | | |
| exit | | immediately exit program |
| cond | statement | executes statement if the top of the stack is non-zero. |
|  | | |
| def | name | begins proceedure definition |
| end | name | ends proceedure definition |
| call | name | calls named proceedure |
| | | |
| push | integer or character | pushes the arg onto the top of the stack |
| pop | | discards the top of the stack |
| dup | | duplicates the top of the stack |
| len | | pushes current stack size onto stack |
| rot | | rotates the stack. 2nd-from-top is depth (>=0), top is times. Positive 'times' rotates right, negative 'times' rotates left.
| | | |
| add | | adds the top 2 elements of the stack and puts the result on the stack |
| sub | | subtracts the top element from the 2nd-from-top element and puts the result on the stack |
| mul | | multiplies the top 2 elements of the stack and puts the result on the stack |
| div | | divides the 2nd-from-top element by the top element and puts the result on the stack |
| mod | | modulus the 2nd-from-top element by the top element and puts the result on the stack |
| pow | | raises 2nd-from-top element by the top element and puts the result on the stack |
| | | |
| eq | | compares the top 2 elements. puts `1` on the stack if equal, `0` if not equal |
| neq | | compares the top 2 elements. `1` if not equal, `0` if equal |
| gt | | compares the 2nd-from-top element with top element. |
| gte | | compares the 2nd-from-top element with top element. |
| lt | | compares the 2nd-from-top element with top element. |
| lte | | compares the 2nd-from-top element with top element. |
| not | | inverts the 'truth' of the top element. Non-zero (true) becomes `0` (false). `0` (false) becomes a random non-zero integer (true). |
| | | |
| print | [literal I or C] | pops and prints the top of the stack to `stdout`. Optional arg formats as an integer or character. |
| println | [literal I or C] | pops and prints the top of the stack to `stdout` followed by a newline. Optional arg formats as an integer or character. |
| read | [double-quoted string] | reads a single value from `stdin` and puts it on the top of the stack. Optional arg is a prompt. |

# example

```
# this program defines a recursive proceedure that
# counts down from a 100 to 1

def countdown
    dup
    println
    push 1
    sub

    dup
    cond call countdown # recurse if top != 0
end countdown

read "enter a positive number: "
call countdown
```