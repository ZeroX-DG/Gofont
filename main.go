package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	option, err := getOption()
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := http.Get(option.fontURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", html)

}
