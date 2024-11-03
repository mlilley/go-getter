package main

import (
	"fmt"

	"github.com/mlilley/go-grabber/engine"
)

func main() {
	err := engine.Download("https://www.youtube.com/watch?v=E_vE6iDLUHA")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
