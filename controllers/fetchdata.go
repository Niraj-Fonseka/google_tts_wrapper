package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchNewsData(source string, apiKey string) []byte {

	newsURLBuilder := fmt.Sprintf("https://newsapi.org/v2/top-headlines?sources=%s&apiKey=%s", source, apiKey)
	resp, _ := http.Get(newsURLBuilder)

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
