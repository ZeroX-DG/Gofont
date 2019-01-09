package main

import (
	"cssgenerate"
	"extractfont"
	"fmt"
	"gofont/loggers/downloadedfontlogger"
	"gofont/loggers/generatecsslogger"
	"io/ioutil"
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
	css, err := FetchHTML(option.fontURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// extract font urls and names from google font api
	fontURLs, fontNames := extractfont.Extract(css)
	var fontLocalPaths []string

	// download all fonts to destination
	for _, fontURL := range fontURLs {
		fileName := path.Base(fontURL)
		fullpath := path.Join(option.destination, fileName)
		err := DownloadFile(fontURL, fullpath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fontLocalPaths = append(fontLocalPaths, fullpath)
	}

	if option.cssFile != "" {
		newcss, err := cssgenerate.GenerateCSS(css, fontURLs, fontLocalPaths, option.cssFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(option.cssFile, []byte(newcss), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// register loggers
	logger.Register(downloadedfontlogger.Log)
	logger.Register(generatecsslogger.Log)

	// log data
	data := logger.LogData{
		FontNames: fontNames,
		CSSFile:   option.cssFile,
	}
	logger.Log(data)
}
