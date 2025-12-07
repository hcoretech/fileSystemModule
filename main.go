package main

import (
	"fmt"
	"os"
	// "github.com/hcoretech/fileSystemModule/loop"
)

func main() {

	go func() {
		fmt.Println("This is a goroutine running in the background.")

	}()

	fmt.Println("Starting the program...")
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current working directory:", currentDir)

	// Get the absolute path of the current directory

}
