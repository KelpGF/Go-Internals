package main

// func add(a, b int) int {
// 	sum := a + b // sum is allocated on the stack
// 	return sum   // sum is deallocated when the function returns
// }

// type User struct {
// 	Name string
// }

// func NewUser(name string) *User {
// 	// name leaks to heap because it is used in pointer receiver
// 	u := User{Name: name} // u is allocated on the heap because it is a pointer receiver
// 	return &u
// }

// func StoreInMap() map[string]*int {
// 	m := make(map[string]*int) // m can survive the function
// 	x := 100                   // x is allocated on the stack
// 	m["x"] = &x                // x is stored in the map as a pointer, so it escapes to the heap
// 	return m                   // m escapes to the heap only because it is returned

// 	// when a data isn't returned, no one can access it, so it doesn't escape to the heap
// }

// func StartGoroutine() {
// 	i := 2

// 	go func() {
// 		// println(i) // i doesn't escape to the heap: can inline

// 		fmt.Println(i) // i escapes to the heap
// 	}()

// 	go func(i int) {
// 		// println(i) // i doesn't escape to the heap

// 		fmt.Println(i) // i escapes to the heap
// 	}(i)
// }

func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

func main() {
	recursive(10)
	// println(add(1, 2))     // can inline ---> when compile, it will be replaced by 1+2
	// fmt.Println(add(1, 2)) // add result escapes to heap because the fmt.Println needs to access it
}
