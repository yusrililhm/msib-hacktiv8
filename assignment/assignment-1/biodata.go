package main

import (
	"assignment-1/util"
	"os"
)

func main() {
	args := os.Args
	util.ValidateArgs(args)
}
