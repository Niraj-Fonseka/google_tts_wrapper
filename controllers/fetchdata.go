package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchNewsData(source string, apiKey string) []byte {

	newsURLBuilder := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", apiKey)
	resp, _ := http.Get(newsURLBuilder)

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
