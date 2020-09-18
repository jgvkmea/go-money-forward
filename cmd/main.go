package main

import (
	"fmt"
	"os"

	"github.com/jgvkmea/go-money-forward/service"
)

func main() {
	// service.Server()
	err := service.UpdateBankData()
	fmt.Fprintf(os.Stderr, "err: %v", err)
}
