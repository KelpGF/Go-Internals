# Advancing in Go Runtime Arch

## Synchronization Problems

Go routines share the same address space, so they can access the same variables. This can lead to synchronization problems when not using channels, because the order of execution is not guaranteed and any go routine can access a variable at any time.

### Race Conditions

Data races occur when two or more goroutines access the same variable concurrently and at least one of the accesses is a write.

For instance, some go routines access a variable to make many decisions. If the variable is not protected, another go routines can edit it while the first go routine is still making decisions based on the old value. This can lead to unexpected results and bugs.

#### Mutex

To avoid it, you can use Mutex (Mutual Exclusion) to protect the variable. Mutex is a synchronization primitive that allows only one goroutine to access a variable at a time.

For instance, a go routine A locks the variable, reads it, and unlocks it. While the variable is locked, no other go routine can access it. If another go routine B tries to access the variable while it is locked, it will be blocked until the variable is unlocked.

#### Deadlocks

Mutex looks a good solution, but it can lead to deadlocks. Deadlocks occur when two or more goroutines are waiting for each other to release the lock.

This can happen when a go routine A locks a variable and waits for another go routine B to unlock another variable, while the other go routine C is waiting for the go routine A to unlock the first variable.

Or when a go routine locks a variable and does not unlock it.

## Channels

Channels are a mechanism to synchronize goroutines and communicate between them. They allow that goroutines send and receive values to and from each other with security and efficiency.

"Do not communicate by sharing memory; instead, share memory by communicating." - Rob Pike

Instead of sharing variables, you can use channels to send and receive values between goroutines. This way, you can avoid synchronizations problems such deadlocks and race conditions.

### Unbuffered Channels

Unbuffered channels have no capacity. They block the sender until the receiver is ready to receive the value. This way, the sender and receiver are synchronized.

In resume, we can send only one value at a time and the sender will be blocked until the receiver processes the value.

### Buffered Channels

Buffered channels have a capacity. They allow the sender to send values until the buffer is full. When the buffer is full, the sender is blocked until the receiver processes a value.
