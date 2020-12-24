# go-remarkable2pdf
Go library to parse and render Remarkable lines files as PDF.

## Usage

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
	output.Close()
}

```

## Features Supported

- [x] Line thickness
- [x] Eraser
- [x] Highlighter
- [x] Multiple layers
- [x] Multiple pages
- [x] .rm lines file
- [x] Notebook zip file
- [ ] PDF annotations
- [ ] ePUB
- [ ] Layer naming and visibility
- [ ] Brush strokes (ie, the painbrush shows "brush marks" based on the speed of your stroke)

## Example Output

**Remarkable Desktop:**

![Remarkable Desktop](static/remarkable-desktop.png | width=100)

**go-remarkable2pdf**

![go-remarkable2pdf](static/go-remarkable2pdf.png | width=100)

**Original**

![Original](static/original.jpg | width=100)