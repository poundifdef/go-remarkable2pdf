# go-remarkable2pdf
Go library to parse and render Remarkable lines files as PDF.

## Usage

```go
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
- [x] Notebook .zip file
- [ ] PDF annotations
- [ ] ePUB
- [ ] Layer naming and visibility
- [ ] Line effects (ie, the painbrush shows "brush marks" based on the speed of your stroke)

## Example Output

**Remarkable Desktop:**

![](/static/remarkable-desktop.png)

**go-remarkable2pdf**

![](/static/go-remarkable2pdf.png)

**Original**

![](/static/original.jpg)

## Acknowledgements

- https://plasma.ninja/blog/devices/remarkable/binary/format/2017/12/26/reMarkable-lines-file-format.html
- https://remarkablewiki.com/tech/filesystem#lines_file_format
- https://github.com/juruen/rmapi