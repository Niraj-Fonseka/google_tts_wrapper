package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google_tts_wrapper/generatespeech"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type TextToSpeechResponse struct {
	AudioContent string `json:"audioContent"`
}

type DataToDecode struct {
	Data string `json:"data"`
}

func GETDailyReport(c *gin.Context) {
	news := generatespeech.GenerateNewsStatement()

	_, err := getDecodedBase64(string(news))

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	//generating audio file
	generateAudioFile()
}

func RunDailyReport(c *gin.Context) {
	var decodeFromPOST DataToDecode

	err := c.BindJSON(&decodeFromPOST)

	fmt.Printf("User Input : %s \n", decodeFromPOST.Data)
	if err != nil {
		fmt.Println(err)
	}

	_, err = getDecodedBase64(decodeFromPOST.Data)

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	//generating audio file
	generateAudioFile()

	// //Playing the audio file
	// cmd := exec.Command("play", "dest_audio_file.mp3")

	// if err = cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func generateAudioFile() {
	cmdBase := exec.Command("sh", os.Getenv("SCRIPT"))
	//cmd := exec.Command("cat", "dailyReport.txt")

	if err := cmdBase.Run(); err != nil {
		log.Fatal(err)
	}
}

type TextToSpeechConfig struct {
	AudioConfig struct {
		AudioEncoding string `json:"audioEncoding"`
		Pitch         string `json:"pitch"`
		SpeakingRate  string `json:"speakingRate"`
	} `json:"audioConfig"`
	Input struct {
		Text string `json:"text"`
	} `json:"input"`
	Voice struct {
		LanguageCode string `json:"languageCode"`
		Name         string `json:"name"`
	} `json:"voice"`
}

func getDecodedBase64(toDecode string) (int, error) {
	url := "https://texttospeech.googleapis.com/v1beta1/text:synthesize?key=" + os.Getenv("KEY")
	fmt.Println("URL:>", url)

	var ttsConfig TextToSpeechConfig
	ttsConfig.AudioConfig.AudioEncoding = "LINEAR16"
	ttsConfig.AudioConfig.Pitch = "0.00"
	ttsConfig.AudioConfig.SpeakingRate = "1.00"

	ttsConfig.Input.Text = toDecode

	ttsConfig.Voice.LanguageCode = "en-US"
	ttsConfig.Voice.Name = "en-US-Wavenet-D"

	data, _ := json.Marshal(ttsConfig)
	textToSpeechConfigJson := []byte(string(data))

	fmt.Printf("Text to speech Config : %s \n", string(textToSpeechConfigJson))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(textToSpeechConfigJson))

	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var responseTTS TextToSpeechResponse

	json.Unmarshal(body, &responseTTS)

	//Decoded Data being written to the file
	f, err := os.Create("dailyReport.txt")
	if err != nil {
		return -1, err
	}

	defer f.Close()

	n3, err := f.WriteString(responseTTS.AudioContent)
	if err != nil {
		return -1, err
	}

	fmt.Println("Byte Count : ")
	fmt.Println(n3)
	return n3, err
}
