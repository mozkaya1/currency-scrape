package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Data struct {
	Gold    string `json:"Gold"`
	Dolar   string `json:"Dolar"`
	Euro    string `json:"Euro"`
	Sterlin string `json:"Sterlin"`
	Bist100 string `json:"Bist100"`
	Bitcoin string `json:"Bitcoin"`
	Silver  string `json:"Silver"`
	Brent   string `json:"Brent"`
}

func getData(url string, parseID string) (Data, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Data{}, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 OPR/94.0.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return Data{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Data{}, fmt.Errorf("failed to parse HTML: %v", err)
	}
	goldPrice := doc.Find(parseID).Eq(0).Text()
	dollarPrice := doc.Find(parseID).Eq(1).Text()
	euroPrice := doc.Find(parseID).Eq(2).Text()
	sterlinPrice := doc.Find(parseID).Eq(3).Text()
	Bist100 := doc.Find(parseID).Eq(4).Text()
	Bitcoin := doc.Find(parseID).Eq(5).Text()
	Silver := doc.Find(parseID).Eq(6).Text()
	Petrol := doc.Find(parseID).Eq(7).Text()

	return Data{
		Gold:    goldPrice,
		Dolar:   dollarPrice,
		Euro:    euroPrice,
		Sterlin: sterlinPrice,
		Bist100: Bist100,
		Bitcoin: Bitcoin,
		Silver:  Silver,
		Brent:   Petrol,
	}, nil
}

func main() {
	// Fetch the data
	result, err := getData("https://www.doviz.com", "span.value")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	resp, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Error json encoding: %v", err)
	}
	fmt.Println(string(resp))
}
