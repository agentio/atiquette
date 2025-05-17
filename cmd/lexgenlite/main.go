package main

import (
	"context"
	"os"
)

func main() {
	if err := Cmd(context.Background()).Execute(); err != nil {
		os.Exit(1)
	}
}
