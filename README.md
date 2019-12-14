# unusable
the **UNU**sable **S**t**A**ck **B**ased **L**anguag**E**

A toy stack based language written in Go.

# stack
64 bit signed integers
'infinite' depth

# scripts
A single file read line by line. Each line is a statement, comment, or blank.

A statement is:
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

# examples

```
# this program defines a recursive proceedure that
# counts down from a user-entered integer to 1

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

```
# This program calculates and prints the first N Fibonacci
# numbers where N is a user-entered value between 1 and 92.
# (Fib(93) overflows signed 64-bit integers.)

# rotates top 3 elements right 2.
# used to move the 'bottom' element up.
def rot2
    push 3
    push 2
    rot
end rot2

# rotates top 3 elements right 1.
# mostly used to move the counter below the fib nums
def rot1
    push 3
    push 1
    rot
end rot1

# swaps the top 2 elements of the stack
def swap
    push 2
    push 1
    rot
end swap

# ensures that the user param N is in [1,92].
def check_bounds
    dup
    push 1
    lt
    cond exit

    dup
    push 92
    gt
    cond exit
end check_bounds

# the recursive fibonacci proceedure.
def fib
    # print counter (debug)
    #dup
    #print
    #push 32    # ascii space
    #print C

    call rot1   # put counter below args

    # print the fib number
    # want 2nd from top
    call swap
    dup
    println
    call swap

    dup
    call rot2   # move a copy of high fib number down to safety
    add         # calculate the next fib number

    call rot2   # bring counter to top
    push 1
    sub         # decrease counter

    dup
    cond call fib # recurse if counter != 0
end fib

# initialize
push 1
push 1

read "Which Fibonacci? range [1, 92] "
call check_bounds
call fib
```