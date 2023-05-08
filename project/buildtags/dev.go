//go:build dev
// +build dev

package main

import "fmt"

const version = "DEV"

func main() {
	fmt.Printf("running %s version", version)
}
