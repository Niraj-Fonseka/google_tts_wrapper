package controllers

import (
	"fmt"
	"io/ioutil"
	//"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
	} `json:"articles"`
}

func GetBreakingNews(c *gin.Context) {
	source := c.Param("source")
	apiKey := os.Getenv("NEWSAPI")
	newsURLBuilder := fmt.Sprintf("https://newsapi.org/v2/top-headlines?sources=%s&apiKey=%s", source, apiKey)
	resp, _ := http.Get(newsURLBuilder)

	body, _ := ioutil.ReadAll(resp.Body)
	//var news News

	// err := json.Unmarshal(body,&news)

	// fmt.Println(err)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, string(body))
}
