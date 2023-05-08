//go:build !release
// +build !release

package main

import "fmt"

const version = "RELEASE"

func main() {
	fmt.Printf("running %s version", version)
}
