package main

import (
	"os"
)

func main() {
	cli := &CLI{out: os.Stdout, err: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
