package rm2pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/poundifdef/go-remarkable2pdf/models"
)

func renderLinesFile(pdf *gofpdf.Fpdf, rm *models.Rm) error {
	//fmt.Println(rm)

	xpmax := float32(8.5)
	ypmax := float32(11)

	xrmax := float32(1404.0)
	yrmax := float32(1872.0)

	pdf.OpenLayerPane()

	pdf.AddPage()

	for i, layer := range rm.Layers {
		// TODO: Use metadata to use correct layer names
		// TODO: Can we support layer visibility?

		layerName := fmt.Sprintf("Layer %d", i+1)
		pdfLayer := pdf.AddLayer(layerName, true)

		pdf.BeginLayer(pdfLayer)

		for _, line := range layer.Lines {
			// If this is an "erase area" brush type, then don't display anything
			// Assumes that the "area" has been erased.
			if line.BrushType == models.EraseArea {
				continue
			}

			// Set the right brush color
			switch line.BrushColor {
			case models.Black:
				pdf.SetDrawColor(0, 0, 0)
			case models.Grey:
				pdf.SetDrawColor(160, 160, 160)
			case models.White:
				pdf.SetDrawColor(255, 255, 255)
			}

			// Special logic for different brush types
			switch line.BrushType {
			case models.Eraser:
				pdf.SetDrawColor(255, 255, 255)
			case models.HighlighterV5:
				// Highlight the text in transparent yellow
				pdf.SetDrawColor(250, 250, 0)
				pdf.SetAlpha(0.2, "Normal")
			default:
				pdf.SetAlpha(1, "Normal")
			}

			// Requred to keep lines smooth
			pdf.SetLineCapStyle("round")
			pdf.SetLineJoinStyle("round")

			for k, _ := range line.Points {
				if k > 0 {
					// The smallest possible line (fineliner thin) has a Width
					// value of 2. Arbitrarily set that to 0.01" as the baseline.
					// Then, scale any segment's thickness to be relative to that.

					lineWidth := 0.01 * line.Points[k].Width / 2.0
					pdf.SetLineWidth(float64(lineWidth))

					x1 := float64((xpmax) * (line.Points[k-1].X / xrmax))
					y1 := float64((ypmax) * (line.Points[k-1].Y / yrmax))

					x2 := float64((xpmax) * (line.Points[k].X / xrmax))
					y2 := float64((ypmax) * (line.Points[k].Y / yrmax))

					pdf.Line(x1, y1, x2, y2)

				}
			}
		}

		pdf.EndLayer()
	}

	return nil
}
