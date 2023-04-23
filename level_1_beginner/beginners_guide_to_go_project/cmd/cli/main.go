package main

import "fmt"

func main() {
	// We are creating out entry point in the cmd/cli/main.go because this allows us to expand
	// easily in the future. for example if we wanted to create a server entry pointer, we could
	// then put server/main.go in the root on the same level as cmd. TOTAL PERSONAL PREFERENCE
	fmt.Println("Awesome CLI v0.0.1")
}
