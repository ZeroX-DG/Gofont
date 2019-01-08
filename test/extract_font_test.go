package gofont

import (
	"extractfont"
	"testing"
)

func TestExtractRemoteURLs(t *testing.T) {
	payload, err := LoadTestData("sample_single_font.css")
	if err != nil {
		t.Error(err)
	}
	expected := []string{
		"https://fonts.gstatic.com/s/lato/v14/S6uyw4BMUTPHjxAwXjeu.woff2", "https://fonts.gstatic.com/s/lato/v14/S6uyw4BMUTPHjx4wXg.woff2",
	}
	actual := extractfont.ExtractRemoteURLs(payload)
	if len(actual) != 2 {
		t.Error("Expected length of remoteURLs to be 2, got ", len(actual))
		return
	}
	if actual[0] != expected[0] {
		t.Error("Expected remote url to be ", expected[0], ", got ", actual[0])
		return
	}
	if actual[1] != expected[1] {
		t.Error("Expected remote url to be ", expected[1], ", got ", actual[1])
		return
	}
}

func TestExtractNames(t *testing.T) {
	payload, err := LoadTestData("sample_single_font.css")
	if err != nil {
		t.Error(err)
	}
	expected := "Lato"
	actual := extractfont.ExtractNames(payload)
	if len(actual) != 1 {
		t.Error("Expected length of names to be 1, got ", len(actual))
		return
	}
	if actual[0] != expected {
		t.Error("Expected first font name to be ", expected, ", got ", actual[0])
		return
	}
}
