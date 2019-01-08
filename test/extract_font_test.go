package gofont

import (
	"extractfont"
	"testing"
)

func TestExtractRemoteURLs(t *testing.T) {
	payload := ".class { local: url(https://fonts.googleapis.com/font.ttf) format('ttf') }"
	expected := "https://fonts.googleapis.com/font.ttf"
	actual := extractfont.ExtractRemoteURLs(payload)
	if len(actual) != 1 {
		t.Error("Expected length of remoteURLs to be 1, got ", len(actual))
		return
	}
	if actual[0] != expected {
		t.Error("Expected remote url to be https://fonts.googleapis.com/font.ttf, got ", actual[0])
		return
	}
}
