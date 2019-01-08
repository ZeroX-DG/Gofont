package main

import (
	"extractfont"
	"fmt"
	"gofont/loggers/downloadedfontlogger"
	"logger"
	"path"
)

func main() {
	option, err := getOption()
	if err != nil {
		fmt.Println(err)
		return
	}

	// fetch css from google font api
	html, err := FetchHTML(option.fontURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// extract font urls and names from google font api
	fontURLs, fontNames := extractfont.Extract(html)

	// download all fonts to destination
	for _, fontURL := range fontURLs {
		fileName := path.Base(fontURL)
		fullpath := path.Join(option.destination, fileName)
		err := DownloadFile(fontURL, fullpath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// register loggers
	logger.Register(downloadedfontlogger.Log)

	// log data
	data := logger.LogData{
		FontNames: fontNames,
	}
	logger.Log(data)
}
