package gofont

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

// LoadTestData load test data for test
func LoadTestData(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(path.Join("./data", filename))
	if err != nil {
		errorMessage := "Error while loading test data " + filename
		fmt.Println(errorMessage)
		return "", errors.New(errorMessage)
	}
	return string(bytes), nil
}
