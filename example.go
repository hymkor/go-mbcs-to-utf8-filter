//+build ignore

package main

import (
	"fmt"
	"os"

	"mbcs"
)

func main() {
	filter := mbcs.NewFilter(os.Stdin)
	for filter.Scan() {
		fmt.Println(filter.Text())
	}
	if err := filter.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
