package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const OriginURL = "http://localhost:8080"
const LanguageURL = "https://www.lingq.com/api/v2/languages/"
const CardURL = "https://www.lingq.com/api/v2/%v/cards/?page_size=1&page=%v"
const Token = "LINGQ_API_KEY"

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{OriginURL}
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.GET("/card", CardHandler)
		api.GET("/languages", LanguageHandler)
	}
	router.Run(":3000")
}

func LanguageHandler(c *gin.Context) {
	resp, err := http.Get(LanguageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Connection with LingQ server failed"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(body)
	languages := fromJsonToLanguages(content)
	c.JSON(http.StatusOK, languages)
}

func fromJsonToLanguages(content string) []Language {
	languages := make([]Language, 0)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var language Language
	for decoder.More() {
		err := decoder.Decode(&language)
		if err != nil {
			panic(err)
		}
		languages = append(languages, language)
	}
	return languages
}

func CardHandler(c *gin.Context) {
	language := c.Query("lang")

	language = strings.TrimSpace(language)
	if len(language) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Language is required"})
		return
	}

	cardNumber := getRandomCardNumber(language)

	// if card number is -1 it means it was not possible to stablish a connection with LingQ server
	if cardNumber == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Connection with LingQ server failed"})
		return
	}

	// if card number is 0 it means no card was found for the given language
	if cardNumber == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Card not found"})
		return
	}

	cardContent := getCard(cardNumber, language)
	card := fromJsonToCard(cardContent)
	c.JSON(http.StatusOK, card)
}

func getRandomCardNumber(language string) int {
	// setting page as 1 because what we care about here is only the total of cards
	// which is always retrieved by the api no matter the page
	content := string(getCard(1, language))

	// if there's no count node, some error may have happened when connecting with LingQ server
	if !gjson.Get(content, "count").Exists() {
		return -1
	}

	totalCards := gjson.Get(content, "count").Int()
	if totalCards == 0 {
		return 0
	}

	// it's necessary to seed Random otherwise it will generate the same sequence of numbers
	// if the server is restarted
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(totalCards))
}

func getCard(page int, language string) []byte {
	url := fmt.Sprintf(CardURL, language, page)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return body
}

func fromJsonToCard(content []byte) Card {
	value := gjson.Get(string(content), "results")
	decoder := json.NewDecoder(strings.NewReader(value.String()))
	_, err := decoder.Token()
	checkError(err)

	var card Card
	err = decoder.Decode(&card)
	checkError(err)
	return card
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

type Language struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

type Hint struct {
	Text string `json:"text"`
}

type Card struct {
	Term  string `json:"term"`
	Hints []Hint `json:"hints"`
}
