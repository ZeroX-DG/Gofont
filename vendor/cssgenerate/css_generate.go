package cssgenerate

import (
	"os"
	"path"
	"strings"
)

// GenerateCSS generate css with the path pointing to the local path
func GenerateCSS(css string, fontRemoteURLs []string, fontLocalPaths []string, destinationCSSFile string) (string, error) {
	result := css
	absoluteCSSFile, err := abs(destinationCSSFile)
	absoluteCSSFileDir := path.Dir(absoluteCSSFile)
	if err != nil {
		return "", err
	}
	for i, url := range fontRemoteURLs {
		absoluteLocalPath, err := abs(fontLocalPaths[i])
		if err != nil {
			return "", err
		}
		relativeLocalPath := GetRelativePath(absoluteCSSFileDir, absoluteLocalPath)
		result = strings.Replace(result, url, relativeLocalPath, -1)
	}
	return result, nil
}

// GetRelativePath given a src path, generate the relative path to the destination path
func GetRelativePath(src string, dest string) string {
	sep := "/"
	srcParts := strings.Split(src, sep)
	destParts := strings.Split(dest, sep)
	fileName := destParts[len(destParts)-1]
	relativePrefix := ""

	for !strings.Contains(dest, strings.Join(srcParts, sep)) {
		srcParts = srcParts[:len(srcParts)-1]
		relativePrefix += ".." + sep
	}

	partWithoutFilename := destParts[:len(destParts)-1]
	remainingDestPath := partWithoutFilename[len(srcParts):]
	if len(remainingDestPath) > 0 {
		return relativePrefix + strings.Join(remainingDestPath, sep) + sep + fileName
	}
	return relativePrefix + fileName
}

// Abs transform relative path to absolute path
func abs(name string) (string, error) {
	if path.IsAbs(name) {
		return name, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, name), err
}
