// Command difference prints the difference (a - b) of two integer arguments.
// It is a thin wrapper over the difference package used by the e2e test to
// exercise Sub through a real built binary. Production source — mutable.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hanahmily/tdd-sandbox/difference"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: difference <a> <b>")
		os.Exit(2)
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid operand a:", err)
		os.Exit(2)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid operand b:", err)
		os.Exit(2)
	}
	fmt.Println(difference.Sub(a, b))
}
