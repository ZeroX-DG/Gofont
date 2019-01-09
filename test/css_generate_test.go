package test

import (
	"cssgenerate"
	"testing"
)

func TestGenerateCSS(t *testing.T) {
	// sample css for font downloaded from google font
	payload1, err := LoadTestData("sample_single_font.css")
	if err != nil {
		t.Error(err)
	}

	// sample css for generated css based on css file location
	payload2, err := LoadTestData("sample_single_font_generated_css.css")
	if err != nil {
		t.Error(err)
	}
	remoteURLs := []string{
		"https://fonts.gstatic.com/s/lato/v14/S6uyw4BMUTPHjxAwXjeu.woff2", "https://fonts.gstatic.com/s/lato/v14/S6uyw4BMUTPHjx4wXg.woff2",
	}
	localPaths := []string{
		"/home/fonts/S6uyw4BMUTPHjxAwXjeu.woff2",
		"/home/fonts/S6uyw4BMUTPHjx4wXg.woff2",
	}
	destinationCSS := "/home/css/style.css"
	actual, err := cssgenerate.GenerateCSS(payload1, remoteURLs, localPaths, destinationCSS)
	if err != nil {
		t.Error(err)
	}
	if actual != payload2 {
		t.Error("Actual generated CSS doesn't match")
		return
	}
}

func TestGenerateRelativePath(t *testing.T) {
	src := "/home/path/to/css"
	dest := "/home/path/to/fonts/font.ttf"

	expected := "../fonts/font.ttf"
	actual := cssgenerate.GetRelativePath(src, dest)

	if actual != expected {
		t.Error("Expected relative path to be ", "expected", ", got ", actual)
		return
	}

	src = "/home/path/to/css"
	dest = "/home/path/to/css/font.ttf"

	expected = "font.ttf"
	actual = cssgenerate.GetRelativePath(src, dest)

	if actual != expected {
		t.Error("Expected relative path to be ", "expected", ", got ", actual)
		return
	}

	src = "/home/path/to/css"
	dest = "/home/path/to/css/fonts/font.ttf"

	expected = "fonts/font.ttf"
	actual = cssgenerate.GetRelativePath(src, dest)

	if actual != expected {
		t.Error("Expected relative path to be ", "expected", ", got ", actual)
		return
	}
}
