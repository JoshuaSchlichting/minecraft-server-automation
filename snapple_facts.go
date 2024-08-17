package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

var snappleFacts string

func init() {
	url := "https://raw.githubusercontent.com/anfederico/fact-bot/master/Facts.txt"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download facts: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	snappleFacts = string(body)
}

func getRandomSnappleFact() string {
	facts := strings.Split(snappleFacts, "\n")
	rand.Seed(uint64(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(facts))
	return facts[randomIndex]
}
