package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Req - make a http request to API
func Req() {
	response, err := http.Get("http://quotes.rest/qod/categories.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// similar of using final in JS
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
