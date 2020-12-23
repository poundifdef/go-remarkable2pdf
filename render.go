package rm2pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/poundifdef/go-remarkable2pdf/models"
)

func renderLinesFile(pdf *gofpdf.Fpdf, rm *models.Rm) error {

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

			// Change line thickness
			// TODO: Perhaps the line width should be proportional
			// 		 to the actual value of this enum
			switch line.BrushSize {
			case models.Large:
				pdf.SetLineWidth(0.03)
			case models.Medium:
				pdf.SetLineWidth(0.02)
			case models.Small:
				pdf.SetLineWidth(0.01)
			}

			// Special logic for different brush types
			switch line.BrushType {
			case models.Eraser:
				// TODO: pay attention to line thickness here
				pdf.SetDrawColor(255, 255, 255)
			case models.HighlighterV5:
				// Highlight the text in transparent yellow
				pdf.SetLineWidth(0.1)
				pdf.SetDrawColor(255, 255, 0)
				pdf.SetAlpha(0.2, "Normal")
			default:
				pdf.SetAlpha(1, "Normal")
			}

			// Requred to keep lines smooth
			pdf.SetLineCapStyle("round")
			pdf.SetLineJoinStyle("round")

			// Use FPDF's functionality to draw a line between
			// consecutive points
			for k, point := range line.Points {
				xpdf := float64((xpmax) * (point.X / xrmax))
				ypdf := float64((ypmax) * (point.Y / yrmax))
				if k == 0 {
					pdf.MoveTo(xpdf, ypdf)
				} else {
					pdf.LineTo(xpdf, ypdf)
				}
			}
			pdf.DrawPath("D")
		}

		pdf.EndLayer()
	}

	return nil
}
