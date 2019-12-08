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
arg types:
`I` - integer
`C` - UTF8 character
`S` - UTF8 string
`[]` - indicates the arg is optional

| keyword | args | description |
| --- | --- | --- |
| # | | a comment line. Inline comments are not allowed. |
| | | |
| exit | | immediately exit program |
| push | I,C | pushes the arg onto the top of the stack |
| pop | | pops and discards the top of the stack |
| dup | | duplicates the top of the stack |
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
| print | [I,C] | prints the top of the stack to `stdout`. Optional arg formats as an integer or character. |
| println | [I,C] | prints the top of the stack to `stdout` followed by a newline. Optional arg formats as an integer or character. |
| read | [S] | reads a single value from `stdin` and puts it on the top of the stack. Optional arg is a prompt. |
