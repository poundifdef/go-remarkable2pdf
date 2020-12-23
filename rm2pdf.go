package rm2pdf

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"strings"

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
// input is the name of the Notebook, in zip, format, to open.
//
// output is where the PDF is written.
func RenderRmNotebook(input string, output io.Writer) error {
	reader, err := zip.OpenReader(input)
	defer reader.Close()

	if err != nil {
		return err
	}

	pdf := gofpdf.New("P", "in", "letter", "")

	for _, file := range reader.File {
		// TODO: make sure we're rendering these in order
		if strings.HasSuffix(file.Name, ".rm") {
			fd, err := file.Open()
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(fd)
			if err != nil {
				return err
			}

			rm := &models.Rm{}
			err = rm.UnmarshalBinary(b)
			if err != nil {
				return err
			}

			err = renderLinesFile(pdf, rm)

			if err != nil {
				return err
			}

		}
	}

	pdf.Output(output)

	return nil
}
