package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RetrieveString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func RetrieveJson(url string, target interface{}) (interface{}, error) {
	// Convert response body to Todo struct
	str, err := RetrieveString(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(str), &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}
