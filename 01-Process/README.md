# Process

A process is an instance of a program that is being executed. It contains the program code and its current activity. Depending on the operating system (OS), a process may be made up of multiple threads of execution that execute instructions concurrently.

## Components of a Process

- **Addressing Space**: The range of addresses that is allocated to the process for the execution of the program.
- **Context**: Group of data that the OS saves to manager a process. It includes the process state, program counter, CPU registers, CPU scheduling information, memory-management information, accounting information, and instruction pointer (IP) to the next instruction to be executed.
- **CPU Registers**:
  - Registers are small, fast storage locations within the CPU that are used to store data temporarily to run the program.
  - Performs arithmetic and logical operations.
  - Address Register
    - Storage of memory addresses.
- **Heap**: Memory that is dynamically allocated during the process runtime. It increase and decrease in size as the program runs.
- **Stack**: Memory that is used to store local variables and function call information. It is automatically allocated and deallocated. LIFO (Last In First Out) structure.
- **Status Registers / Flags**: Contains information about the state of the process.
  - Flags Zero (Z): Set if the result of an operation is zero.
  - Flags Sign (S) or Negative (N): Set if the result of an operation is negative.
  - Flags Overflow: Set if the result of an operation is too large to fit in the register.

## Process Lifecycle

A process goes through various states during its lifecycle. The states are as follows:

- **Creation**: The process is created by the OS.
  - fork() for Unix/Linux
  - CreateProcess() for Windows.
- **Ready**: The process is waiting to be assigned to a processor.
- **Execution**: The process is running.
- **Waiting/Blocked**: The process is waiting for some event to occur. Commonly, it is waiting for I/O operations to complete.
- **Termination**: The process has finished execution or has forced to stop.
  - Exit: Finished with success.
  - Killed: Forced to stop.

## Creating a new Process

- **UNIX/Linux**: The `fork()` system call is used to create a new process.
  - The new process is a child and an exact copy of the existing process.
  - Each fork() call returns 0 for the child process and the process ID for the parent process.
  - A new process has its own address space, stack, and heap.
  - Each process are independent and have their own memory space.
  - The parent process receives the child process ID (positive integer).
  - The child process receives 0 that indicates it is the child process.

## Managing Processes

### Scheduler

The scheduler is responsible for selecting the next process to run.
It can switch between processes to provide concurrent execution.
There are many scheduling algorithms to maximize CPU utilization and throughput.

**What it can do:**

- Select processes from the ready queue and choose one to run.
- Allocate CPU time to the selected process. From ready to running state.
- Deallocate CPU time. From running to ready state.

#### Types of Schedulers

- **Cooperative Scheduler**: The process gives up control voluntarily.
  - Processes that are running has the control of when to give up the CPU.
- **Preemptive Scheduler**: The OS decides when to give control to another process.
  - The OS can interrupt the process and give control to another process.

Cooperative processes can monopolize the CPU and not give up control.

Preemptive processes has many context switching that means lower performance.

## Threads

Threads are the smallest unit of execution that can be scheduled by the OS. It is a lightweight process that can be managed independently by the OS and are contained within a process.

How threads are in the same process, they can share the process memory space, address space, and resources. It means that threads access and modify the same data, variables, and resources.

Within a process, threads can run concurrently working on different tasks. It can improve the performance of the application.

### Race Condition

Race condition happens when 2 or more threads are changing the same variable at the same time.

It can lead to unpredictable results because, for instance, we are calculating the value of a variable and another thread changes it before we finish the calculation. The result will be different from the expected.

### Deadlock

Deadlock happens when 2 or more threads are waiting for each other to release a resource.

For instance, Thread A is waiting for Thread B to release a resource, and Thread B is waiting for Thread A to release a resource. Both threads are waiting for each other to release the resource, and the program is stuck.

### Threads x Memory

Threads needs a lower memory space than processes because they share the same memory space.

And each thread has its own Stack and registers. Aprox. 2MB of memory (linux).

## Parallelism vs Concurrency

### Parallelism

Parallelism is when multiple tasks are running at the same time. It requires multiple CPUs or cores to execute the tasks simultaneously.

For instance, a computer with 4 cores with 4 collaborative schedulers can run 4 tasks at the same time.

### Concurrency

Concurrency is when multiple tasks are making progress, but not necessarily running at the same time. It can be achieved by interleaving tasks.

For instance, a computer with 1 core can run multiple tasks by switching between them (preemptive scheduler).
