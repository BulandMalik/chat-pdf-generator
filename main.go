package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// ChatEntry represents a single chat message
type ChatEntry struct {
	Timestamp time.Time
	User      string
	Message   string
	R         int
	G         int
	B         int
	IconPath  string
}

// emojiToImage maps emoji characters to their image paths
var emojiToImage = map[string]string{
	"✅": "./images/check.png",
	"❌": "./images/close.png",
	"📚": "./images/book.png",
	"🎯": "./images/target.png",
	"🚀": "./images/rocket.png",
}

// PDFGenerator handles PDF document creation and styling
type PDFGenerator struct {
	pdf          *gofpdf.Fpdf
	title        string
	entries      []ChatEntry
	margin       float64
	pageWidth    float64
	pageHeight   float64
	logoPath     string
	headerHeight float64
	footerHeight float64
}

// Add a function to check if images exist
func (g *PDFGenerator) checkImages() {
	for emoji, path := range emojiToImage {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("Warning: Image file not found: %s for emoji %s\n", path, emoji)
		}
	}
}

// NewPDFGenerator creates a new PDF document with default settings
func NewPDFGenerator(title string) *PDFGenerator {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(true, 20)

	// Set page dimensions (A4 size in mm)
	pageWidth := 210.0
	pageHeight := 297.0
	margin := 20.0

	generator := &PDFGenerator{
		pdf:          pdf,
		title:        title,
		logoPath:     "logo.png",
		margin:       margin,
		pageWidth:    pageWidth,
		pageHeight:   pageHeight,
		headerHeight: 40.0,
		footerHeight: 20.0,
	}

	// Check if images exist
	generator.checkImages()

	return generator
}

// AddChatEntry adds a chat entry to the document
func (g *PDFGenerator) AddChatEntry(entry ChatEntry) {
	// Store the entry
	g.entries = append(g.entries, entry)
}

func (g *PDFGenerator) addHeader() {
	g.pdf.SetFont("Arial", "B", 24)
	g.pdf.SetTextColor(0, 0, 0)

	// Add logo if exists
	if _, err := os.Stat(g.logoPath); err == nil {
		g.pdf.Image(g.logoPath, g.margin, g.margin, 30, 30, false, "", 0, "")
	}

	// Add title
	g.pdf.SetY(g.margin + 10)
	g.pdf.SetX(g.margin + 35)
	g.pdf.Cell(0, 20, g.title)
}

func (g *PDFGenerator) addFooter() {
	g.pdf.SetY(g.pageHeight - g.footerHeight)
	g.pdf.SetFont("Arial", "I", 8)
	g.pdf.SetTextColor(128, 128, 128)
	g.pdf.Cell(0, 10, fmt.Sprintf("Generated on %s", time.Now().Format("2006-01-02 15:04:05")))
}

func (g *PDFGenerator) addText(text string, x, y float64, fontSize float64) {
	g.pdf.SetFont("Arial", "", fontSize)
	g.pdf.Text(x, y, text)
}

// GeneratePDF creates the PDF document
func (g *PDFGenerator) GeneratePDF(filename string) error {
	g.pdf.AddPage()
	g.addHeader()

	contentTop := g.margin + g.headerHeight
	contentBottom := g.pageHeight - g.footerHeight
	g.pdf.SetY(contentTop)

	for _, entry := range g.entries {
		// Add timestamp and user
		timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")
		g.pdf.SetFont("Arial", "", 10)
		g.pdf.SetTextColor(100, 100, 100)
		g.pdf.Cell(150, 10, timestamp)
		g.pdf.Cell(100, 10, entry.User)

		// Add message
		g.pdf.SetY(g.pdf.GetY() + 10)
		g.pdf.SetFont("Arial", "", 12)
		g.pdf.SetTextColor(0, 0, 0)

		messageWidth := g.pageWidth - (2 * g.margin)
		g.pdf.MultiCell(messageWidth, 10, entry.Message, "", "", false)
		g.pdf.SetY(g.pdf.GetY() + 10)

		// Check if we need a new page
		if g.pdf.GetY() > contentBottom {
			g.pdf.AddPage()
			g.addHeader()
			g.pdf.SetY(contentTop)
		}
	}

	g.addFooter()
	return g.pdf.OutputFileAndClose(filename)
}

func main() {
	// Create sample chat entries
	entries := []ChatEntry{
		{
			Timestamp: time.Now().Add(-2 * time.Hour),
			User:      "System",
			Message:   "Criteria Pass: ✅ Material: Both the reference and candidate components use the same materials for the connector: PBT for the housing and Copper Alloy with Gold plating for the metal parts. This indicates full compatibility in terms of material composition.",
			R:         0,
			G:         128,
			B:         0,
		},
		{
			Timestamp: time.Now().Add(-1 * time.Hour),
			User:      "System",
			Message:   "Not a match: ❌ Number of pins: The reference component has 4 pins, while the candidate component has 5 pins. For connectors, the number of pins must match exactly to ensure compatibility. Therefore, the candidate component is not compatible with the reference component.",
			R:         255,
			G:         0,
			B:         0,
		},
	}

	// Create PDF generator
	generator := NewPDFGenerator("Compatibility Report")

	// Add chat entries
	for _, entry := range entries {
		generator.AddChatEntry(entry)
	}

	// Generate the PDF
	err := generator.GeneratePDF("compatibility_report.pdf")
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
		return
	}

	fmt.Println("PDF generated successfully: compatibility_report.pdf")
}
