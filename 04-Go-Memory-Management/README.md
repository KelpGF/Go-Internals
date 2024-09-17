# Go Memory Management

The Go language use as base the TCMalloc, a memory allocator developed by Google. This allocator is optimized for multi-threaded applications, so it is very efficient for Go programs.

As time goes by, the Go team has been improving the memory management and now the own Go Runtime is in charge of managing the memory. Go memory allocator is called `mallocgc`.

## Mallocgc

The `mallocgc` is a memory allocator that is used by the Go Runtime to manage the memory. It is a garbage collector that is responsible for allocating and deallocating memory in the heap.

It separates the objects in three groups:

- **Tiny**: Objects that are smaller than 16 bytes.
- **Small**: Objects that are between 16 and 32 bytes.
- **Large**: Objects that are larger than 32 bytes.

It avoids always requesting memory from the OS, so it uses a memory pool to store the objects that are allocated in the heap.

![malloc-gc](img/malloc-gc.png)

When needs memory:
Goroutine -> Processor -> mcaches -> mcentral -> mheap -> OS

### Spans

Spans are memory blocks that are used to store the objects and can be tiny, small, or large. The spans are managed by the Mcentral.

### Mheap

Mheap is the memory pool that is used to store the objects that are allocated in the heap. It is responsible for requesting memory from the OS.

### Mcentral

Mcentral takes care of the spans that are used to store objects of the same size. It is responsible for managing the spans, being not necessary to request a new span from the OS every time that is necessary to allocate memory.

### Mcache

Mcache is a local cache used by P's. First, the P tries to allocate memory from the Mcache, and if it is not possible, the P requests a span from the Mcentral.

## Garbage Collector

Garbage Collector is a mechanism that is responsible for freeing the memory that is not being used by the program. This mechanism is used to avoid memory leaks and to improve the performance of the program.

How the Go Garbage Collector works:

- **non-generational**: Consider all objects in the heap as the same generation, without distinguishing between young and old objects.
- **concurrent**: Run concurrently with the application, so it does not stop the execution of the program.
- **tricolor**: Uses a tricolor algorithm to determine which objects are reachable and which are not.

### Reachable Objects

The garbage collector considers an object as reachable if it is possible to access it from the root of the program.

- **Roots:** Roots are entry points to access the reachable objects. Includes global variables, stack variables, and CPU registers.
- **Referenced Objects:** Objects that are accessible from the roots. The garbage collector marks these objects as reachable.
  - Example: If a global variable references an object A, and the object A references an object B, the objects A and B are considered reachable.
  - If a object C that is not referenced by any root, it is considered unreachable.

Then, the garbage collector marks the objects that are reachable and frees the memory of the objects that are not reachable.

### How GO Garbage Collector Works

1. SWT (Stop The World): The garbage collector stops the execution of the program.
    1. Mark Setup: The garbage collector prepares the marking phase. Raises the Write Barrier.
    2. Write Barrier: The Write Barrier is a mechanism that is used to intercept the writes in the heap. It is used to mark the objects that are being modified.
2. Marking Work: The garbage collector starts marking the reachable objects.
    1. Uses 25% of the CPU to mark the objects.
    2. Mark Assist: If the marking phase takes too long, the garbage collector asks the goroutines to help with the marking.
    3. In this phase, the program is already running again. The marking occurs concurrently with the program.
3. Mark Termination: The garbage collector finishes the marking phase.
    1. SWT again, because new objects can be created during the marking phase.
    2. Finalize the marking phase.
    3. Turn off the Write Barrier.
4. Sweeping: The garbage collector frees the memory of the objects that are not reachable.

#### Marking Phase

On this step, the garbage collector marks the objects that are reachable with a tricolor algorithm. The algorithm uses three colors to determine the state of the objects:

- **White**: Objects that are not marked yet.
- **Grey**: Objects that are marked but not scanned yet.
- **Black**: Objects that are marked and scanned.

1. Mark each root as grey.
2. For each grey object, mark it as black and mark its references as grey.
3. Repeat the process until there are no more grey objects.

#### GOGC Environment Variable

The GOGC environment variable is used to set the size of heap that triggers the garbage collector. The default value is 100, which means that the garbage collector is triggered when the heap size reaches 100%.

Example: If we not set the GOGC environment variable and the heap size on the last garbage collection was 4mb, the garbage collector will be triggered when the heap size reaches 8mb (4mb+100%).

If we set the GOGC environment variable less than 100, the garbage collector will be triggered more frequently.

- Less memory usage.
- More CPU usage (each collection uses 25% of the CPU).
- More STW (Stop The World).

If we set the GOGC environment variable greater than 100, the garbage collector will be triggered less frequently.

- More memory usage.
- Less CPU usage.
- Less STW.

#### GC Trace

```bash
gc 1 @0.019s 0%: 0.014+0.56+0.01 ms clock, 0.029+0/0.55/0+0.021 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
```

- **Clock Time**: Time spent by the garbage collector on the point of view of a external visitor. That includes aspects of execution and waiting time. Time since starts until the end of the garbage collector cycle.
- **CPU Time**: Time spent by the garbage collector on the point of view of the CPU. That includes only the time that the CPU is working. Time since starts until the end of the garbage collector cycle, excluding the time that the CPU is waiting.

- **gc N**: Number of cycle of garbage collector.
  - If the number is 1, it is the first cycle. If the number is 200, it is the 200th cycle.
- **@0.019s**: Time since the program started until the garbage collector cycle starts.
- **0%**: Percentage of the CPU usage time spent by the garbage collector.
- **0.014+0.56+0.01 ms clock**: Time spent by the garbage collector.
  - **0.014**: Time spent in the setup phase.
  - **0.56**: Time spent in the concurrent phase.
  - **0.01**: Time spent from the beginning of marking completion to the end of sweeping.
- **0.029+0/0.55/0+0.021 ms cpu**: CPU time spent by the garbage collector.
  - **0.029**: CPU time spent in the SWT_SWEEP_TERMINATION phase.
    - **+0**: Additional CPU time spent in the marking phase.
  - **0.55**: CPU time spent in the MARK_AND_SWEEP phase.
  - **0+0.021**: CPU time spent in the SWT_MARK_TERMINATION phase.
    - **0+**: Additional CPU time spent SWT before this phase.
- **4->4->1 MB**: Heap size before the garbage collector cycle, heap size after the garbage collector cycle starts, and the heap size after the garbage collector cycle ends.
- **5 MB goal**: The goal of the heap size for the next garbage collector cycle.
- **8 P**: Number of processors used by the Go Scheduler.

#### Running example

```bash
GODEBUG=gctrace=1 GOGC=300 go run main.go m
```

#### Memory Limit

We can set a memory limit for the Go program using the `runtime/debug` package.

```go
package main

import (
  "runtime/debug"
)

func main() {
  debug.SetMaxStack(1000000)
}
```

When we run the same program with different memory limits, we can see that the garbage collector is triggered more frequently when the memory limit is lower.

Its happens because the program reaches the memory limit faster, so the garbage collector is triggered more frequently.

```bash
GODEBUG=gctrace=1 GOGC=300 go run main.go m-limit
```
