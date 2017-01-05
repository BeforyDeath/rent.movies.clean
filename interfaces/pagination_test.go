package interfaces

import (
	"encoding/json"
	"testing"
)

func TestNewPagination(t *testing.T) {
	var jsonParams = []map[string]interface{}{
		{"limit": json.Number("10")},
		{"page": json.Number("10")}, // limit default 100
		{"limit": json.Number("10"), "page": json.Number("5")},
	}

	pages, _ := NewPagination(jsonParams[0])
	if pages.Limit != 10 {
		t.Error("Expected 10, got", pages.Limit)
	}

	// limit default 100
	pages, _ = NewPagination(jsonParams[1])
	if pages.Offset != 900 {
		t.Error("Expected 900, got", pages.Offset)
	}

	pages, _ = NewPagination(jsonParams[2])
	if pages.Offset != 40 {
		t.Error("Expected 40, got", pages.Offset)
	}

}
