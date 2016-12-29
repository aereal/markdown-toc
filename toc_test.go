package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestInjectToc(t *testing.T) {
	input := strings.NewReader(`<!-- toc -->
<!-- /toc -->

# h1

h1

# 姉
`)

	expected := `<!-- toc -->
* [h1](#h1)
* [姉](#%E5%A7%89)
<!-- /toc -->

# h1<a name="h1"></a>

h1

# 姉<a name="%E5%A7%89"></a>
`
	output := new(bytes.Buffer)
	InjectToc(input, output)
	if got := output.String(); strings.TrimSpace(got) != strings.TrimSpace(expected) {
		t.Fatalf("\nExpected:\n%s\n----\nGot:\n%s----\n", expected, got)
	}
}
