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
    // Render a single .rm lines file as PDF
	input, _ := os.Open("4.rm")
	output, _ := os.Create("out.pdf")
	rm2pdf.RenderRmFile(input, output)
	output.Close()

    // Render a full Notebook file as PDF
	output, _ := os.Create("out.pdf")
	rm2pdf.RenderRmNotebook("Notebook.zip", output)
}

```