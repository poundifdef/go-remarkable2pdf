# go-remarkable2pdf
Go library to parse and render Remarkable lines files as PDF.

# Example

```
package main

import (
	"fmt"
	"os"

	rm2pdf "github.com/poundifdef/go-remarkable2pdf"
)

func main() {
	input, _ := os.Open("4.rm")
	output, _ := os.Create("out.pdf")
	rm2pdf.RenderRmFile(input, output)
	output.Close()
}

```