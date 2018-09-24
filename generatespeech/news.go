package generatespeech

import (
	"time"
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

func GenerateNewsStatement() string {
	// news := FetchNewsData("cnn", os.Getenv("NEWSAPI"))

	// var newsStructured News

	// json.Unmarshal(news, &newsStructured)

	// var newsOutput string
	// newsOutput += "Here are the top headlines. "
	// for _, record := range newsStructured.Articles {
	// 	newsOutput += record.Title + ". "
	// }

	newsOutput := "News Statement is here"
	return newsOutput
}
