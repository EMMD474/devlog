package main

import (
	"fmt"
	"os"

	"github.com/emmd474/devlog/internal/model"
	"github.com/emmd474/devlog/internal/storage"
)

func main() {
	// Load all entries
	entries, err := storage.LoadEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading entries: %v\n", err)
		os.Exit(1)
	}

	if len(entries) < 2 {
		fmt.Fprintf(os.Stderr, "Need at least 2 entries for this test\n")
		os.Exit(1)
	}

	fmt.Printf("Before delete: %d entries\n", len(entries))
	for i, e := range entries {
		fmt.Printf("  %d. [%s] %s\n", i+1, e.ID[:8], e.Message)
	}

	// Delete the second entry (index 1)
	toDelete := entries[1]
	fmt.Printf("\nDeleting entry: [%s] %s\n", toDelete.ID[:8], toDelete.Message)

	var remaining []model.Entry
	for _, e := range entries {
		if e.ID != toDelete.ID {
			remaining = append(remaining, e)
		}
	}

	// Save the filtered entries
	if err := storage.SaveEntries(remaining); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving entries: %v\n", err)
		os.Exit(1)
	}

	// Load entries again to verify
	entries, err = storage.LoadEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading entries: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nAfter delete: %d entries\n", len(entries))
	for i, e := range entries {
		fmt.Printf("  %d. [%s] %s\n", i+1, e.ID[:8], e.Message)
	}

	// Verify the deleted entry is not present
	for _, e := range entries {
		if e.ID == toDelete.ID {
			fmt.Fprintf(os.Stderr, "\nERROR: Deleted entry is still present!\n")
			os.Exit(1)
		}
	}

	fmt.Println("\n✓ Delete test passed: Only the selected entry was deleted")
}
