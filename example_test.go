package epaper_test

import (
	"fmt"
	"log"

	"github.com/fstanis/epaper"
)

const port = "/dev/ttyUSB0"

func Example() {
	ep, err := epaper.New(port)
	if err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer ep.Close()

	if err := ep.Handshake(); err != nil {
		log.Fatalf("Handshake failed: %v", err)
	}

	if err := ep.Clear(); err != nil {
		log.Fatalf("Clear failed: %v", err)
	}

	if err := ep.DisplayText(100, 100, "Hello world!"); err != nil {
		log.Fatalf("DisplayText failed: %v", err)
	}
	fmt.Println("Text displayed")

	if err := ep.DrawRect(50, 50, 250, 150); err != nil {
		log.Fatalf("DrawRect failed: %v", err)
	}
	fmt.Println("Rect drawn")

	if err := ep.Update(); err != nil {
		log.Fatalf("Update failed: %v", err)
	}

	// Output:
	// Text displayed
	// Rect drawn
}
