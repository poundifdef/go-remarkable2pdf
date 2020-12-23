package rm2pdf

import (
	"io"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	"github.com/poundifdef/go-remarkable2pdf/models"
)

// RenderRmFile converts a single .rm lines file and renders
// it as a single-page PDF.
//
// input is a reader which points to the raw .rm file
//
// output is where the PDF is written.
func RenderRmFile(input io.ReadCloser, output io.Writer) error {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	rm := &models.Rm{}
	err = rm.UnmarshalBinary(b)
	if err != nil {
		return err
	}

	pdf := gofpdf.New("P", "in", "letter", "")
	err = renderLinesFile(pdf, rm)
	if err != nil {
		return err
	}
	pdf.Output(output)

	return nil
}

// RenderRmNotebook converts an entire Remarkable Notebook
// and renders it as a multiple-page PDF.
//
// input is a reader which points to the raw Notebook file, which
// is a zip file format
//
// output is where the PDF is written.
func RenderRmNotebook(input io.ReadCloser, output io.Writer) string {
	return "hello"
}
