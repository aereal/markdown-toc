package cli

import (
	"flag"
	"fmt"
	"io"
	"os"

	markdowntoc "github.com/aereal/markdown-toc"
)

func NewCLI() *CLI {
	return &CLI{out: os.Stdout, err: os.Stderr}
}

type CLI struct {
	out, err io.Writer
}

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("markdown-toc", flag.ContinueOnError)
	flags.SetOutput(cli.err)

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) < 1 {
		fmt.Fprintln(cli.err, "Usage: markdown-toc FILE.md")
		return 1
	}

	target := parsedArgs[0]
	f, err := os.Open(target)
	if err != nil {
		fmt.Fprintln(cli.err, err)
		return 1
	}
	defer f.Close()

	markdowntoc.InjectToc(f, cli.out)

	return 0
}
