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
