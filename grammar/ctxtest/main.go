package main

import (
	"context"
	"fmt"
)

type stringKey string

func setKet() {

	key := stringKey("psetReq")
	ctx := context.WithValue(context.Background(), key, "req")
	getKey(ctx)
}
func getKey(ctx context.Context) string {
	var key = stringKey("psetReq")
	req := ctx.Value(key).(string)
	fmt.Println(req)
	return req
}

func main() {
	setKet()
}
