package extractfont

import (
	"regexp"
)

// Extract extract font urls and names
func Extract(css string) ([]string, []string) {
	remoteURLs := ExtractRemoteURLs(css)
	fontNames := ExtractNames(css)
	return remoteURLs, fontNames
}

// ExtractRemoteURLs Analyze the css and returns the remote urls
func ExtractRemoteURLs(css string) []string {
	result := []string{}
	re := regexp.MustCompile(`url\((.*?)\) format`)
	matches := re.FindAllStringSubmatch(css, -1)
	for _, match := range matches {
		remoteURL := match[1]
		result = append(result, remoteURL)
	}
	return result
}

// ExtractNames Analyze the css and returns the name of each font
func ExtractNames(css string) []string {
	nameMaps := map[string]bool{}
	re := regexp.MustCompile(`font-family: '(.*?)'`)
	matches := re.FindAllStringSubmatch(css, -1)
	for _, match := range matches {
		remoteURL := match[1]
		_, isDuplicate := nameMaps[remoteURL]
		if !isDuplicate {
			nameMaps[remoteURL] = true
		}
	}

	fontnames := make([]string, len(nameMaps))
	i := 0
	for k := range nameMaps {
		fontnames[i] = k
		i++
	}
	return fontnames
}
