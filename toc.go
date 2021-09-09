package markdowntoc

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"
	"unicode"
)

const marker = "<!-- toc -->"

func InjectToc(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	f := func(c rune) bool {
		return !(unicode.IsSpace(c) || c == 35)
	}
	var toc []string
	var formatted []string
	tocIndex := -1
	i := 0
	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)
		if strings.HasPrefix(line, "#") {
			pos := strings.IndexFunc(line, f)
			if pos == -1 {
				pos = 1
			}
			title := line[pos:]
			escapedTitle := url.QueryEscape(title)
			toc = append(toc, fmt.Sprintf("* [%s](#%s)", title, escapedTitle))
			formatted = append(formatted, line+fmt.Sprintf(`<a name="%s"></a>`, escapedTitle))
		} else if line == marker {
			tocIndex = i
		} else {
			formatted = append(formatted, line)
		}
		i++
		if err == io.EOF {
			break
		}
	}
	if tocIndex != -1 {
		toc = append([]string{marker}, toc...)
		formatted = append(formatted[:tocIndex], append(toc, formatted[tocIndex:]...)...)
	}
	fmt.Fprintln(out, strings.Join(formatted, "\n"))
}
