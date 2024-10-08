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

#### When to use

- Synchronization: when we need ensure that the sender operation is directly connected to the receiver operation.
- Handshake: when we need to make sure that the receiver is ready to continue the work.
- Timed Events: when we need to wait for a data in a channel with a timeout.
- Coordinated finish: when we need to finish when all goroutines finish their work.

### Buffered Channels

Buffered channels have a capacity. They allow the sender to send values until the buffer is full. When the buffer is full, the sender is blocked until the receiver processes a value.

#### When to use buffer

- High producer and high consumer: when the producer is faster than the consumer, the producer can send values to the buffer and continue its work.
- Pipelines of goroutines: when you have a pipeline of goroutines, not use buffered channels can lead to a blocking way or deadlocks.
- Asynchronous tasks: when you want to send values to a channel and continue the work without waiting for the receiver, such logs or metrics.
- Multiple producers for the same channel: when you have multiple goroutines sending values to the same channel, the buffer can help to avoid blocking the sender.
- I/O operations: for example, reading files or network operations.

#### How large should be the buffer?

It depends on the application. You should consider the amount of data that the producer can send and the amount of data that the consumer can process.

However, we have some deadlines that can help you to decide the buffer size:

- Producer and Consumer Rate
  - **Variable Rate:** If the producer and consumer have variable rates, you should consider a bigger buffer to avoid blocking.
  - **Fixed Rate:** If the producer and consumer have fixed rates, you can use a buffer with the same size of the producer rate.

- Latency, Throughput and Performance
  - **Low Latency:** If you need low latency, you should consider a smaller buffer.
  - **High Throughput:** If you need high throughput, you should consider a bigger buffer.

- Memory
  - **Memory:** If you have memory constraints, you should consider a smaller buffer.

### Internal Implementation

```go
type hchan struct {
  qcount   uint           // total amount of data in the queue
  dataqsiz uint           // size of the circular queue
  buf      unsafe.Pointer // points to an array of dataqsiz elements
  elemsize uint16         // size of each element in the queue
  closed   uint32         // 0 or 1
  sendx    uint           // send index: where to put the next data element sent to the channel
  recvx    uint           // receive index: where to get the next data element in the channel
  recvq    waitq          // list of recv waiters (go routines) to receive data
  sendq    waitq          // list of send waiters (go routines) to send data
  lock     mutex          // lock to protect the queue
}
```

### How it works

#### Send

![send-to-channel](img/send-to-channel.png)

#### Receive

![receive-from-channel](img/receive-from-channel.png)

#### Close

![close-channel](img/close-channel.png)
