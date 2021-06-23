//+build ignore

package main

import (
	"os"

	"mbcs"
)

func main() {
	filter := mbcs.NewFilter(os.Stdin)
	for filter.Scan() {
		println(filter.Text())
	}
}
