# unusable
the **UNU**sable **S**t**A**ck **B**ased **L**anguag**E**

A toy stack based language written in Go.

# stack
64 bit signed integers
'infinite' depth

# scripts
single file
read line by line

line is:
`KEYWORD [arg]`

# keywords

| keyword | args | description |
| --- | --- | --- |
| # | | a comment line. Inline comments are allowed. |
| | | |
| exit | | immediately exit program |
| cond | statement | |
|  | | |
| def | name | |
| end | name | |
| call | name | |
| | | |
| push | integer or character | pushes the arg onto the top of the stack |
| pop | | discards the top of the stack |
| dup | | duplicates the top of the stack |
| len | | pushes current stack size onto stack |
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
| gt | | compares the 2nd-from-top element with top element.|
| gte | | |
| lt | | |
| lte | | |
| not | | inverts the 'truth' of the top element. Non-zero (true) becomes `0` (false). `0` (false) becomes a random non-zero integer (true). |
| | | |
| print | [literal I or C] | pops and prints the top of the stack to `stdout`. Optional arg formats as an integer or character. |
| println | [literal I or C] | pops and prints the top of the stack to `stdout` followed by a newline. Optional arg formats as an integer or character. |
| read | [double-quoted string] | reads a single value from `stdin` and puts it on the top of the stack. Optional arg is a prompt. |
