package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetBreakingNews(c *gin.Context) {
	source := c.Param("source")
	apiKey := os.Getenv("NEWSAPI")

	news := FetchNewsData(source, apiKey)
	// fmt.Println(err)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, string(news))
}

func FetchNewsData(source string, apiKey string) []byte {

	newsURLBuilder := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", apiKey)
	resp, _ := http.Get(newsURLBuilder)

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
