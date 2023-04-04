package main

import (
	"strconv"
)

func main() {
	ints()

}
func ints() int32 {
	var ret int32
	if t, err := strconv.Atoi("value"); err != nil { //nolint:gosec
		return ret
	} else {
		ret = int32(t)
	}
	return ret
}
