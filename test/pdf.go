package main

// package main

// import (
// 	"fmt"
// 	pdfsdk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
// 	"log"
// )

// func main() {

// 	// Create new PDF generator
// 	pdfg, err := pdfsdk.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set global options
// 	pdfg.Dpi.Set(300)
// 	pdfg.Grayscale.Set(true)

// 	// Create a new input page from an URL
// 	page := pdfsdk.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

// 	// Set options for this page
// 	page.FooterRight.Set("[page]")
// 	page.FooterFontSize.Set(10)
// 	page.Zoom.Set(0.95)

// 	// Add to document
// 	pdfg.AddPage(page)

// 	// Create PDF document in internal buffer
// 	err = pdfg.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write buffer contents to file on disk
// 	err = pdfg.WriteFile("./simplesample.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Done")
// 	// Output: Done
// }
