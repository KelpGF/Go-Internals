package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	commandParam := os.Args[1]

	if commandParam == "m-limit" {
		memoryLimit()
	}

	memoryManagement()
}

func memoryManagement() {
	for i := 0; i < 10; i++ {
		allocateMemory(10 * 1024 * 1024)
		time.Sleep(time.Second)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func memoryLimit() {
	// change the number 10 to 100 to see the difference
	debug.SetMemoryLimit(10 * 1024 * 1024) // 10MB

	for i := 0; i < 10; i++ {
		_ = allocateMemory(20 * 1024 * 1024) // allocate 20MB

		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
		fmt.Printf("Lookups = %v\n", m.Lookups)
		fmt.Printf("Mallocs = %v\n", m.Mallocs)
		fmt.Printf("Frees = %v\n", m.Frees)
		fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)
		fmt.Printf("HeapSys = %v MiB\n", m.HeapSys/1024/1024)
		fmt.Printf("HeapIdle = %v MiB\n", m.HeapIdle/1024/1024)
		fmt.Printf("HeapInuse = %v MiB\n", m.HeapInuse/1024/1024)
		fmt.Printf("HeapReleased = %v MiB\n", m.HeapReleased/1024/1024)
		fmt.Printf("HeapObjects = %v\n", m.HeapObjects)
		fmt.Printf("StackInuse = %v MiB\n", m.StackInuse/1024/1024)
		fmt.Printf("StackSys = %v MiB\n", m.StackSys/1024/1024)
		fmt.Printf("MSpanInuse = %v MiB\n", m.MSpanInuse/1024/1024)
		fmt.Printf("MSpanSys = %v MiB\n", m.MSpanSys/1024/1024)
		fmt.Printf("MCacheInuse = %v MiB\n", m.MCacheInuse/1024/1024)
		fmt.Printf("MCacheSys = %v MiB\n", m.MCacheSys/1024/1024)
		fmt.Printf("BuckHashSys = %v MiB\n", m.BuckHashSys/1024/1024)
		fmt.Printf("GCSys = %v MiB\n", m.GCSys/1024/1024)
		fmt.Printf("OtherSys = %v MiB\n", m.OtherSys/1024/1024)

		time.Sleep(time.Second)

		fmt.Println("--------------------------------------------------")
	}
}

func allocateMemory(size int) []byte {
	return make([]byte, size)
}
