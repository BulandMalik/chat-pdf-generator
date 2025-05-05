package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/signintech/gopdf"
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
	"✅": "./images/check.svg",
	"❌": "./images/close.svg",
	"📚": "./images/book.svg",
	"🎯": "./images/target.svg",
	"🚀": "./images/rocket.svg",
}

// PDFGenerator handles PDF document creation and styling
type PDFGenerator struct {
	pdf          *gopdf.GoPdf
	title        string
	entries      []ChatEntry
	logoPath     string
	pageHeight   float64
	pageWidth    float64
	margin       float64
	pageCount    int
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
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Set page dimensions (A4 size in points)
	pageWidth := 595.28  // A4 width in points
	pageHeight := 841.89 // A4 height in points

	// Add fonts
	err := pdf.AddTTFFont("DejaVuSans", "./fonts/DejaVuSans.ttf")
	if err != nil {
		fmt.Printf("Error adding DejaVuSans font: %v\n", err)
	}

	// Try multiple emoji font paths
	emojiPaths := []string{
		"./fonts/NotoColorEmoji.ttf",
		"/System/Library/Fonts/Apple Color Emoji.ttc",       // macOS
		"/usr/share/fonts/truetype/noto/NotoColorEmoji.ttf", // Linux
		"C:\\Windows\\Fonts\\seguiemj.ttf",                  // Windows
	}

	var emojiFontLoaded bool
	for _, path := range emojiPaths {
		err = pdf.AddTTFFont("NotoColorEmoji", path)
		if err == nil {
			emojiFontLoaded = true
			fmt.Printf("Successfully loaded emoji font from: %s\n", path)
			break
		}
	}

	if !emojiFontLoaded {
		fmt.Printf("Warning: Could not load any emoji font. Icons may not display correctly.\n")
	}

	generator := &PDFGenerator{
		pdf:          &pdf,
		title:        title,
		logoPath:     "logo.png",
		pageHeight:   pageHeight,
		pageWidth:    pageWidth,
		margin:       20,
		pageCount:    0,
		headerHeight: 50,
		footerHeight: 30,
		entries:      make([]ChatEntry, 0),
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

// addHeader adds a header with logo and title to the current page
func (g *PDFGenerator) addHeader() {
	// Add logo
	if _, err := os.Stat(g.logoPath); err == nil {
		// Create a rectangle for the image
		rect := &gopdf.Rect{W: 30, H: 30}
		g.pdf.Image(g.logoPath, g.margin, g.margin, rect)
	}

	// Add title
	err := g.pdf.SetFont("DejaVuSans", "", 24)
	if err != nil {
		fmt.Printf("Error setting header font: %v\n", err)
	}
	g.pdf.SetTextColor(0, 0, 0)
	g.pdf.SetX(g.margin + 40)
	g.pdf.SetY(g.margin)
	g.pdf.Cell(nil, g.title)

	// Add horizontal line
	g.pdf.SetLineWidth(0.5)
	g.pdf.Line(g.margin, g.margin+g.headerHeight-10, g.pageWidth-g.margin, g.margin+g.headerHeight-10)

	// Reset position for content
	g.pdf.SetY(g.margin + g.headerHeight)
}

// addFooter adds a footer to the current page
func (g *PDFGenerator) addFooter() {
	// Add horizontal line first
	g.pdf.SetLineWidth(0.5)
	g.pdf.Line(g.margin, g.pageHeight-g.margin-g.footerHeight+5, g.pageWidth-g.margin, g.pageHeight-g.margin-g.footerHeight+5)

	// Add footer text below the line
	err := g.pdf.SetFont("DejaVuSans", "", 10)
	if err != nil {
		fmt.Printf("Error setting footer font: %v\n", err)
	}
	g.pdf.SetTextColor(100, 100, 100) //black
	g.pdf.SetY(g.pageHeight - g.margin - g.footerHeight + 10)
	g.pdf.SetX(g.margin)
	g.pdf.Cell(nil, "https://chiphub.com")

	// Add page number
	g.pageCount++
	g.pdf.SetX(g.pageWidth - g.margin - 30)
	g.pdf.Cell(nil, fmt.Sprintf("Page %d", g.pageCount))
}

// GeneratePDF creates the PDF document
func (g *PDFGenerator) GeneratePDF(filename string) error {
	// Add first page
	g.pdf.AddPage()
	g.addHeader()

	// Content area boundaries
	contentBottom := g.pageHeight - g.margin - g.footerHeight

	// Process entries
	for _, entry := range g.entries {
		// Check if we need a new page
		if g.pdf.GetY() > contentBottom-50 {
			g.addFooter()
			g.pdf.AddPage()
			g.addHeader()
		}

		// Set font for timestamp and user
		err := g.pdf.SetFont("DejaVuSans", "", 10)
		if err != nil {
			fmt.Printf("Error setting DejaVuSans font: %v\n", err)
		}

		// Format timestamp
		timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")

		// Add timestamp and user in a fixed-width format
		g.pdf.SetX(g.margin)
		g.pdf.SetTextColor(100, 100, 100) // Gray color for metadata
		g.pdf.Cell(&gopdf.Rect{W: 150, H: 20}, timestamp)
		g.pdf.Cell(&gopdf.Rect{W: 100, H: 20}, entry.User)

		// Set font and color for message
		err = g.pdf.SetFont("DejaVuSans", "", 12)
		if err != nil {
			fmt.Printf("Error setting DejaVuSans font: %v\n", err)
		}
		g.pdf.SetTextColor(uint8(entry.R), uint8(entry.G), uint8(entry.B))

		// Add message with proper indentation and wrapping
		messageX := g.margin + 280
		g.pdf.SetX(messageX)
		messageWidth := g.pageWidth - messageX - g.margin

		// Process message text and replace emojis with images
		message := entry.Message
		for emoji, imagePath := range emojiToImage {
			if strings.Contains(message, emoji) {
				// Split message by emoji
				parts := strings.Split(message, emoji)

				// Add first part of text
				if parts[0] != "" {
					g.pdf.MultiCell(&gopdf.Rect{W: messageWidth, H: contentBottom - g.pdf.GetY()}, parts[0])
				}

				// Add emoji image
				iconRect := &gopdf.Rect{W: 10, H: 10}
				g.pdf.Image(imagePath, g.pdf.GetX(), g.pdf.GetY(), iconRect)
				g.pdf.SetX(g.pdf.GetX() + 15)

				// Add remaining text
				if len(parts) > 1 && parts[1] != "" {
					g.pdf.MultiCell(&gopdf.Rect{W: messageWidth - 25, H: contentBottom - g.pdf.GetY()}, parts[1])
				}

				// We've handled this emoji, move to next iteration
				continue
			}
		}

		// If no emojis were found, just add the message as is
		if !strings.ContainsAny(message, "✅❌📚🎯🚀") {
			g.pdf.MultiCell(&gopdf.Rect{W: messageWidth, H: contentBottom - g.pdf.GetY()}, message)
		}

		// Add spacing between entries
		g.pdf.SetY(g.pdf.GetY() + 15)
	}

	// Add footer to the last page
	g.addFooter()

	return g.pdf.WritePdf(filename)
}

func main() {
	// Create sample chat entries
	entries := []ChatEntry{
		{
			Timestamp: time.Now().Add(-2 * time.Hour),
			User:      "Alice",
			Message:   "Project status update ✅ - All tasks completed on time!",
			R:         0,
			G:         128,
			B:         0,
			IconPath:  "./images/check.png",
		},
		{
			Timestamp: time.Now().Add(-1 * time.Hour),
			User:      "Bob",
			Message:   "Need to review the latest changes",
			R:         255,
			G:         0,
			B:         0,
			IconPath:  "./images/close.png",
		},
		{
			Timestamp: time.Now().Add(-45 * time.Minute),
			User:      "Charlie",
			Message:   "Documentation updated ❌: with new features",
			R:         0,
			G:         0,
			B:         255,
			IconPath:  "./images/book.png",
		},
		{
			Timestamp: time.Now().Add(-30 * time.Minute),
			User:      "Alice",
			Message:   "Tests passing with 100% coverage",
			R:         0,
			G:         128,
			B:         0,
			IconPath:  "./images/target.png",
		},
		{
			Timestamp: time.Now().Add(-15 * time.Minute),
			User:      "Bob",
			Message:   "Ready for deployment!",
			R:         0,
			G:         128,
			B:         0,
			IconPath:  "./images/rocket.png",
		},
	}

	// Create PDF generator
	generator := NewPDFGenerator("Chat Log Report")

	// Add chat entries
	for _, entry := range entries {
		generator.AddChatEntry(entry)
	}

	// Generate the PDF
	err := generator.GeneratePDF("chat_log.pdf")
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
		return
	}

	fmt.Println("PDF generated successfully: chat_log.pdf")
}
