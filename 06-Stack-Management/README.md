# Stack Management

## Intro

Stack is a data structure that follows the LIFO (Last In, First Out) principle, used to storage local variable and function calls.

It more simple, but faster than the heap, because it is a contiguous block of memory, and the heap is a pool of memory, then always try to use the stack.

Examples of data that is stored in the stack:

- Local variables
- Function parameters
- Addresses of the next instruction to be
- Variables with scope and life time shared with goroutines

## Stack in Go

Stack works in a segmented way, allowing the program to allocate a new stack when it is needed.

The default size of a stack in Go is 2mb. When a goroutine starts, it uses a small amount of memory, and when it needs more memory, the program allocates a new stack for it.

A stack can decrease its size when it is not being used. The goruntime has checkpoints to check if a stack is being used, and if it is not, the program can, eventually, deallocate it.

### Guard Pages

Guard pages are used to protect of accessing memory that was not allocated to the stack. When a goroutine tries to access memory that was not allocated to it, the program will throw a segmentation fault. So, the guard pages are used to prevent this.

### Stack Overflow

When a goroutine uses all the memory that was allocated to its stack, the program will throw a stack overflow error. This can happen when a goroutine has a large number of recursive calls, or when it has a large number of local variables.

## How know if something is in the stack or heap?

How decides where a variable will be allocated is the compiler.

### Escape Analysis

Escape analysis is a technique used by the compiler to decide where a variable will be allocated. If the variable can survive outside the scope where it was declared, the compiler will allocate it in the heap, otherwise, it will allocate it in the stack.

There are some cases that we know where a variable will be allocated:

- **Returning Pointers:** When you return a pointer to a variable, the compiler will allocate it in the heap, because the variable will survive outside the scope where it was declared.
- **Storage in a data structure:** When you store a variable in a map, slice, or channel, and the data structure is shared with other scopes, the compiler will allocate it in the heap.
- **Goroutines:** When you pass a local variable to a goroutine, the compiler will allocate it in the heap because the goroutine can outlive the scope where the variable was declared.

#### Checking if a variable escaped to the heap

To check if a variable escaped to the heap, you can use the `-m` flag in the `go build` command.

```bash
go build -gcflags "-m" main.go
```

#### Debugging with dlv

To debug where a variable was allocated, you can use the `dlv` debugger.

Install the `dlv` debugger:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Build the program with the `-gcflags "-N -l"` flags:

```bash
go build -gcflags "-N -l" main.go
```

run the program with the `dlv` debugger:

```bash
dlv debug
```

Set a breakpoint in the line where the variable was declared:

```bash
break main.go:10
```

Set a breakpoint in the function where the variable was declared:

```bash
break main.fn
```

Check the stack trace:

```bash
stack
```
