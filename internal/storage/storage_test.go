package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/emmd474/devlog/internal/model"
)

func TestGenerateID(t *testing.T) {
	id1, err := generateID()
	if err != nil {
		t.Fatalf("generateID() failed: %v", err)
	}

	id2, err := generateID()
	if err != nil {
		t.Fatalf("generateID() failed: %v", err)
	}

	if id1 == id2 {
		t.Errorf("generateID() produced duplicate IDs: %s", id1)
	}

	if len(id1) != 32 { // 16 bytes = 32 hex characters
		t.Errorf("generateID() produced ID of wrong length: got %d, want 32", len(id1))
	}
}

func TestSaveEntry_AssignsID(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// Save an entry
	err := SaveEntry("test message")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	// Load entries
	entries, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	if entries[0].ID == "" {
		t.Error("SaveEntry() did not assign an ID to the entry")
	}

	if entries[0].Message != "test message" {
		t.Errorf("Expected message 'test message', got '%s'", entries[0].Message)
	}
}

func TestLoadEntries_AssignsIDsToExistingEntries(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// Manually create a log file with entries without IDs
	dir := filepath.Join(tmpDir, ".devlog")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	testData := `[
  {
    "id": "",
    "message": "entry without id",
    "date": "2026-01-15T10:00:00Z"
  }
]`
	path := filepath.Join(dir, "logs.json")
	if err := os.WriteFile(path, []byte(testData), 0644); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}

	// Load entries
	entries, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	if entries[0].ID == "" {
		t.Error("LoadEntries() did not assign an ID to the entry")
	}

	// Verify the file was updated
	entries2, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed on second call: %v", err)
	}

	if entries2[0].ID != entries[0].ID {
		t.Error("ID was not persisted to file")
	}
}

func TestUpdateEntry(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// Save two entries
	err := SaveEntry("first entry")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	err = SaveEntry("second entry")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	// Load entries
	entries, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	// Update the first entry
	updated := entries[0]
	updated.Message = "updated first entry"
	err = UpdateEntry(updated)
	if err != nil {
		t.Fatalf("UpdateEntry() failed: %v", err)
	}

	// Load entries again
	entries, err = LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if entries[0].Message != "updated first entry" {
		t.Errorf("Expected message 'updated first entry', got '%s'", entries[0].Message)
	}

	if entries[1].Message != "second entry" {
		t.Errorf("Second entry was modified unexpectedly: got '%s'", entries[1].Message)
	}
}

func TestDeleteEntry_Logic(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	// Save three entries
	err := SaveEntry("entry 1")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	err = SaveEntry("entry 2")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	err = SaveEntry("entry 3")
	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	// Load entries
	entries, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	// Simulate delete logic (filter out entry 2)
	toDelete := entries[1]
	var remaining []model.Entry
	for _, e := range entries {
		if e.ID != toDelete.ID {
			remaining = append(remaining, e)
		}
	}

	// Save the filtered entries
	err = SaveEntries(remaining)
	if err != nil {
		t.Fatalf("SaveEntries() failed: %v", err)
	}

	// Load entries again
	entries, err = LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries after delete, got %d", len(entries))
	}

	if entries[0].Message != "entry 1" {
		t.Errorf("Expected first entry to be 'entry 1', got '%s'", entries[0].Message)
	}

	if entries[1].Message != "entry 3" {
		t.Errorf("Expected second entry to be 'entry 3', got '%s'", entries[1].Message)
	}

	// Verify the deleted entry is not present
	for _, e := range entries {
		if e.ID == toDelete.ID {
			t.Error("Deleted entry is still present")
		}
	}
}

func TestSaveEntry_PreservesDate(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	before := time.Now()
	err := SaveEntry("test message")
	after := time.Now()

	if err != nil {
		t.Fatalf("SaveEntry() failed: %v", err)
	}

	entries, err := LoadEntries()
	if err != nil {
		t.Fatalf("LoadEntries() failed: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	entryDate := entries[0].Date
	if entryDate.Before(before) || entryDate.After(after) {
		t.Errorf("Entry date %v is not between %v and %v", entryDate, before, after)
	}
}
