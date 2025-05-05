// AddFooter adds page numbers and hyperlink
func (g *PDFGenerator) AddFooter() {
	g.pdf.SetY(-15)
	g.pdf.SetFont("Arial", "I", 8)

	// Add hyperlink
	g.pdf.SetTextColor(0, 0, 255)
	g.pdf.Cell(0, 10, "https://chiphub.com")

	// Add page number
	g.pdf.SetTextColor(0, 0, 0)
	g.pdf.Cell(0, 10, fmt.Sprintf("Page %d", g.pdf.PageNo()))
}

// AddHeader adds a header with logo and title
func (g *PDFGenerator) AddHeader(title string) {
	g.pdf.SetY(10)

	// Add title if provided
	if title != "" {
		g.pdf.SetX(10)
		g.pdf.SetFont("Arial", "B", 16)
		g.pdf.Cell(0, 10, title)
		g.pdf.Ln(20)
	}
} 