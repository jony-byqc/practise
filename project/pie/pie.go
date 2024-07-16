package main

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
)

func main() {
	Rows := []string{"apple", "banana", "cherry"}
	dbs := pie.Map(Rows, func(t string) string {

		return t
	})
	fmt.Println(dbs)

}
