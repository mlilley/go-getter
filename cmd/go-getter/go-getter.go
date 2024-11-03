package main

import (
	"fmt"

	"github.com/mlilley/go-getter/engine"
)

func main() {
	err := engine.Download("url goes here")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
