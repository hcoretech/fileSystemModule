package loop

import (
	"fmt"
)

func NewLoop() int {
	fmt.Println("Loop started.")
	return 42
}

func CheckLoop() {
	data := 50
	for {
		if data > 0 {
			fmt.Println("Loop is running...")
			data--
		}
	}

}
