package main

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// ChatEntry represents a single chat message
type ChatEntry struct {
	Timestamp time.Time
	User      string
	Message   string
}

// PDFGenerator handles PDF document creation and styling
type PDFGenerator struct {
	pdf *gofpdf.Fpdf
}

// NewPDFGenerator creates a new PDF document with default settings
func NewPDFGenerator() *PDFGenerator {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(true, 15)
	pdf.AddPage()

	// Set font that supports Unicode and emojis
	pdf.SetFont("Arial", "", 12)

	return &PDFGenerator{pdf: pdf}
}

// AddHeader adds a header with logo and title
func (g *PDFGenerator) AddHeader(title string) {
	g.pdf.SetY(10)

	// Add logo (placeholder - replace with actual logo path)
	g.pdf.Image("logo.png", 10, 10, 30, 0, false, "", 0, "")

	// Add title if provided
	if title != "" {
		g.pdf.SetX(50)
		g.pdf.SetFont("Arial", "B", 16)
		g.pdf.Cell(0, 10, title)
		g.pdf.Ln(20)
	}
}

// AddFooter adds page numbers and hyperlink
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

// ... existing code ...

// AddChatEntry adds a single chat entry to the document
func (g *PDFGenerator) AddChatEntry(entry ChatEntry) {
	g.pdf.SetFont("Arial", "", 12)

	// Format timestamp
	timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")

	// Add entry with proper formatting
	g.pdf.Cell(40, 10, timestamp)
	g.pdf.Cell(30, 10, entry.User)
	g.pdf.MultiCell(0, 10, entry.Message, "", "", false)
	g.pdf.Ln(5)
}

// Save saves the PDF to the specified path
func (g *PDFGenerator) Save(filename string) error {
	return g.pdf.OutputFileAndClose(filename)
}

func main() {
	// Create sample chat entries
	entries := []ChatEntry{
		{
			Timestamp: time.Now().Add(-2 * time.Hour),
			User:      "Alice",
			Message:   "Project status update ✅ All tasks completed on time!",
		},
		{
			Timestamp: time.Now().Add(-1 * time.Hour),
			User:      "Bob",
			Message:   "Need to review the latest changes ❌",
		},
		{
			Timestamp: time.Now().Add(-45 * time.Minute),
			User:      "Charlie",
			Message:   "Documentation updated with new features 📚",
		},
		{
			Timestamp: time.Now().Add(-30 * time.Minute),
			User:      "Alice",
			Message:   "Tests passing with 100% coverage 🎯",
		},
		{
			Timestamp: time.Now().Add(-15 * time.Minute),
			User:      "Bob",
			Message:   "Ready for deployment! 🚀",
		},
	}

	// Create PDF generator
	generator := NewPDFGenerator()

	// Add header
	generator.AddHeader("Chat Log Report")

	// Add chat entries
	for _, entry := range entries {
		generator.AddChatEntry(entry)
	}

	// Save the PDF
	err := generator.Save("chat_log.pdf")
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
		return
	}

	fmt.Println("PDF generated successfully: chat_log.pdf")
}
