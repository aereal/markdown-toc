package main

import (
	"os"

	"github.com/aereal/markdown-toc/internal/cli"
)

func main() {
	app := cli.NewCLI()
	os.Exit(app.Run(os.Args))
}
