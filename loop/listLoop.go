package loop

import (
	"fmt"
)

func newLoop() int {
	fmt.Println("Loop started.")
	return 42
	// for{
	// 	fmt.Println("This is a loop running in the background.")
	// 	// Simulate some work
	// 	for i := 0; i < 100000000; i++ {
	// 		_ = i * i // Just a dummy operation to simulate work
	// 	}
	// }
}
