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
		layerName := fmt.Sprintf("Layer %d", i+1)
		pdfLayer := pdf.AddLayer(layerName, true)
		pdf.BeginLayer(pdfLayer)

		for _, line := range layer.Lines {

			pdf.SetLineCapStyle("round")
			pdf.SetLineJoinStyle("round")

			switch line.BrushColor {
			case models.Black:
				pdf.SetDrawColor(0, 0, 0)
			case models.Grey:
				pdf.SetDrawColor(160, 160, 160)
			case models.White:
				pdf.SetDrawColor(255, 255, 255)
			}

			switch line.BrushSize {
			case models.Large:
				pdf.SetLineWidth(0.03)
			case models.Medium:
				pdf.SetLineWidth(0.02)
			case models.Small:
				pdf.SetLineWidth(0.01)
			}

			switch line.BrushType {
			case models.Eraser:
				pdf.SetDrawColor(255, 255, 255)
			case models.HighlighterV5:
				pdf.SetLineWidth(0.1)
				pdf.SetDrawColor(255, 255, 0)
				pdf.SetAlpha(0.2, "Normal")
			default:
				pdf.SetAlpha(1, "Normal")
			}

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
