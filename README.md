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

## Not Currently Supported

- [ ] Background templates (grid, lines, etc)
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

## Contributing

1. If you want to make any other modification or refactor: please create an
   issue and talk to me prior to making your PR. This way we can talk about the
   feature, approach, and my own ability to commit to reviewing and merging code.

2. I don't guarantee that I will keep this repo up to date, or that I will respond
   in any sort of timely fashion! Your best bet for any change is to keep PRs small
   and focused on the minimum changeset to add your feature.

3. Of course, you are welcome to fork, modify, and distribute this code with your
   changes in accordance with the LICENSE.

## Acknowledgements

- https://plasma.ninja/blog/devices/remarkable/binary/format/2017/12/26/reMarkable-lines-file-format.html
- https://remarkablewiki.com/tech/filesystem#lines_file_format
- https://github.com/juruen/rmapi