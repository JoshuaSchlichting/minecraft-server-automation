package main

import (
	"io"
	"net/http"
	"strings"
	"time"

	log "github.com/JoshuaSchlichting/minecraft-server-automation/logger"
	"golang.org/x/exp/rand"
)

var snappleFacts string

func init() {

	urlPayload, err := getStringFromURL("https://raw.githubusercontent.com/anfederico/fact-bot/master/Facts.txt")
	if err != nil {
		log.Error("Error getting Snapple facts:", err)
	}
	snappleFacts += urlPayload
	// if last char isn't a newline, add one
	if snappleFacts[len(snappleFacts)-1] != '\n' {
		snappleFacts += "\n"
	}
	urlPayload, err = getStringFromURL("https://gist.githubusercontent.com/emctague/c47bea79b013419274d5c353aec31edc/raw/84626d6f03a21dba79e4fa6952c686e960b23e71/all-snapple-facts.txt")
	if err != nil {
		log.Error("Error getting Snapple facts:", err)
	}
	snappleFacts += urlPayload
}

func getStringFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getRandomSnappleFact() string {
	facts := strings.Split(snappleFacts, "\n")
	rand.Seed(uint64(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(facts))
	return facts[randomIndex]
}
